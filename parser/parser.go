// Package parser turns Go code into model.Code.
package parser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path"
	"path/filepath"
	"strings"

	"github.com/matryer/codeform/model"
	"github.com/matryer/codeform/source"
	"github.com/pkg/errors"
)

// Parser parses source code and generates a model.Code.
type Parser struct {
	src *source.Source
	// TargetPackage is the name or path of the package
	// that becomes the point of view for type names.
	// If empty or the same package, types will not be qualified,
	// otherwise they will include the package name.
	TargetPackage string
	qualifier     types.Qualifier
	fset          *token.FileSet
	code          *model.Code
}

// New makes a new Generator for the given package (folder) or
// Go code file.
// If src is specified as a []byte or io.Reader, it will be used as
// the source.
func New(src *source.Source) *Parser {
	return &Parser{
		src: src,
	}
}

// Parse parses the source file and returns the model.Code.
func (p *Parser) Parse() (*model.Code, error) {
	p.qualifier = func(other *types.Package) string {
		//log.Println("----- (name) other:", other.Name(), "p.TargetPackage:", p.TargetPackage, ".")
		if other.Name() == p.TargetPackage {
			return ""
		}
		return other.Name()
	}
	if strings.Contains(p.TargetPackage, "/") {
		p.qualifier = func(other *types.Package) string {
			//log.Println("----- (path)", other.Path(), p.TargetPackage)
			if other.Path() == p.TargetPackage {
				return ""
			}
			return other.Name()
		}
	}
	importer := newSmartImporter(importer.Default())
	code := &model.Code{}
	fset := token.NewFileSet()
	var pkgs map[string]*ast.Package
	if p.src.IsDir {
		var err error
		pkgs, err = parser.ParseDir(fset, p.src.Path, nil, parser.SpuriousErrors)
		if err != nil {
			return nil, err
		}
		for _, pkg := range pkgs {
			if len(p.TargetPackage) == 0 {
				p.TargetPackage = pkg.Name
			}
			if err := p.parsePackage(code, pkg, fset, importer); err != nil {
				return nil, err
			}
		}
	} else {
		file, err := parser.ParseFile(fset, p.src.Path, p.src, parser.SpuriousErrors)
		if err != nil {
			return nil, err
		}
		files := []*ast.File{file}
		conf := types.Config{Importer: importer}
		tpkg, err := conf.Check(p.src.Path, fset, files, nil)
		if err != nil {
			return nil, err
		}
		if len(p.TargetPackage) == 0 {
			p.TargetPackage = tpkg.Name()
		}
		packageModel := model.Package{
			Name: tpkg.Name(),
		}
		if err := p.parseGlobals(&packageModel, tpkg); err != nil {
			return nil, err
		}
		code.Packages = append(code.Packages, packageModel)
	}
	code.TargetPackageName = filepath.Base(p.TargetPackage)
	return code, nil
}

func (p *Parser) parsePackage(code *model.Code, pkg *ast.Package, fset *token.FileSet, importer types.Importer) error {
	i := 0
	files := make([]*ast.File, len(pkg.Files))
	for _, file := range pkg.Files {
		files[i] = file
		i++
	}
	conf := types.Config{Importer: importer}
	tpkg, err := conf.Check(p.src.Path, fset, files, nil)
	if err != nil {
		return err
	}
	timports := tpkg.Imports()
	imports := make([]model.Import, len(timports))
	for i, tpkg := range timports {
		imports[i] = model.Import{Name: tpkg.Path()}
	}
	packageModel := model.Package{
		Name:    tpkg.Name(),
		Imports: imports,
	}
	if err := p.parseGlobals(&packageModel, tpkg); err != nil {
		return err
	}
	code.Packages = append(code.Packages, packageModel)
	return nil
}

func (p *Parser) parseGlobals(packageModel *model.Package, tpkg *types.Package) error {
	scope := tpkg.Scope()
	for _, name := range scope.Names() {

		thing := scope.Lookup(name)
		typ := thing.Type().Underlying()

		parsedType, err := p.parseType(packageModel, tpkg, thing)
		if err != nil {
			return errors.Wrap(err, "parseType")
		}

		switch thing.(type) {
		case *types.Const:
			packageModel.Consts = append(packageModel.Consts, *parsedType)
			continue
		case *types.Var:
			packageModel.Vars = append(packageModel.Vars, *parsedType)
			continue
		}
		switch val := typ.(type) {
		case *types.Signature:
			fn, err := p.parseSignature(val, tpkg)
			if err != nil {
				return errors.Wrap(err, "parseSignature")
			}
			fn.Name = thing.Name()
			packageModel.Funcs = append(packageModel.Funcs, *fn)
		case *types.Interface:
			iface, err := p.parseInterface(val.Complete(), tpkg)
			if err != nil {
				return errors.Wrap(err, "parseInterface")
			}
			iface.Name = thing.Name()
			packageModel.Interfaces = append(packageModel.Interfaces, *iface)
		case *types.Struct:
			structure, err := p.parseStruct(thing.Type(), val, tpkg)
			if err != nil {
				return errors.Wrap(err, "parseStruct")
			}
			structure.Name = thing.Name()
			packageModel.Structs = append(packageModel.Structs, *structure)
		default:
			return fmt.Errorf("%T not supported", thing)
		}
	}
	return nil
}

func (p *Parser) parseStruct(typ types.Type, s *types.Struct, tpkg *types.Package) (*model.Struct, error) {
	structModel := &model.Struct{}
	var err error
	if structModel.Methods, err = p.parseMethods(typ, tpkg); err != nil {
		return nil, errors.Wrap(err, "parseMethods")
	}
	numFields := s.NumFields()
	for i := 0; i < numFields; i++ {
		field := s.Field(i)
		fieldVar, err := p.parseVar(field, tpkg)
		if err != nil {
			return nil, errors.Wrap(err, "parseVar")
		}
		fieldVar.Tag = s.Tag(i)
		structModel.Fields = append(structModel.Fields, *fieldVar)
	}
	return structModel, nil
}

func (p *Parser) parseMethods(typ types.Type, tpkg *types.Package) ([]model.Func, error) {
	methods := make(map[string]model.Func)
	for _, t := range []types.Type{typ, types.NewPointer(typ)} {
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			methodModel, err := p.parseMethod(mset.At(i), tpkg)
			if err != nil {
				return nil, errors.Wrap(err, "parseMethod")
			}
			if _, present := methods[methodModel.Name]; present {
				continue // skip duplicates
			}
			methods[methodModel.Name] = *methodModel
		}
	}
	// turn them into an array
	methodSlice := make([]model.Func, len(methods))
	var i int
	for _, method := range methods {
		methodSlice[i] = method
		i++
	}
	return methodSlice, nil
}

func (p *Parser) parseType(pkg *model.Package, tpkg *types.Package, obj types.Object) (*model.Var, error) {
	typ, err := p.parseVarType(obj, tpkg)
	if err != nil {
		return nil, errors.Wrap(err, "parseVarType")
	}
	return &model.Var{
		Name: obj.Name(),
		Type: typ,
	}, nil
}

func (p *Parser) parseInterface(iface *types.Interface, tpkg *types.Package) (*model.Interface, error) {
	interfaceModel := &model.Interface{}
	mlen := iface.NumMethods()
	for i := 0; i < mlen; i++ {
		method := iface.Method(i).Type().(*types.Signature)
		fn, err := p.parseSignature(method, tpkg)
		if err != nil {
			return nil, err
		}
		fn.Name = iface.Method(i).Name()
		interfaceModel.Methods = append(interfaceModel.Methods, *fn)
	}
	return interfaceModel, nil
}

func (p *Parser) parseMethod(sel *types.Selection, tpkg *types.Package) (*model.Func, error) {
	funcObj, ok := sel.Obj().(*types.Func)
	if !ok {
		return nil, fmt.Errorf("expected *types.Func but got %T", sel.Obj())
	}
	sig, ok := funcObj.Type().(*types.Signature)
	if !ok {
		return nil, fmt.Errorf("expected *types.Signature but got %T", funcObj.Type())
	}
	sigModel, err := p.parseSignature(sig, tpkg)
	if err != nil {
		return nil, errors.Wrap(err, "parseSignature")
	}
	sigModel.Name = funcObj.Name()
	return sigModel, nil
}

func (p *Parser) parseSignature(fn *types.Signature, tpkg *types.Package) (*model.Func, error) {
	args, err := p.parseVars(fn.Params(), tpkg)
	if err != nil {
		return nil, errors.Wrap(err, "Params")
	}
	retArgs, err := p.parseVars(fn.Results(), tpkg)
	if err != nil {
		return nil, errors.Wrap(err, "Results")
	}
	variadic := fn.Variadic()
	if variadic {
		// update the last arg to tell it it's variadic also
		args[len(args)-1].Variadic = true
	}
	return &model.Func{
		Variadic:   variadic,
		Args:       args,
		ReturnArgs: retArgs,
	}, nil
}

func (p *Parser) parseVars(vars *types.Tuple, tpkg *types.Package) ([]model.Var, error) {
	var varModels []model.Var
	paramsLen := vars.Len()
	for i := 0; i < paramsLen; i++ {
		argModel, err := p.parseVar(vars.At(i), tpkg)
		if err != nil {
			return nil, errors.Wrap(err, "param")
		}
		varModels = append(varModels, *argModel)
	}
	return varModels, nil
}

func (p *Parser) parseVar(param *types.Var, tpkg *types.Package) (*model.Var, error) {
	typ, err := p.parseVarType(param, tpkg)
	if err != nil {
		return nil, errors.Wrap(err, "parseVarType")
	}
	n := param.Name()
	arg := &model.Var{
		Anonymous: len(n) == 0,
		Name:      n,
		Type:      typ,
	}
	return arg, nil
}

func (p *Parser) parseVarType(obj types.Object, tpkg *types.Package) (model.Type, error) {
	typeStr := types.TypeString(obj.Type(), p.qualifier)
	if strings.Contains(typeStr, "/") {
		// turn *path/to/package.Type into *package.Type
		pointer := strings.HasPrefix(typeStr, "*")
		typeStr = path.Base(typeStr)
		if pointer {
			typeStr = "*" + typeStr
		}
	}
	return model.Type{
		Name: typeStr,
	}, nil
}
