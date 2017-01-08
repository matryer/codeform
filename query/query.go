// Package query provides high-level code querying capabilities.
package query

import "github.com/matryer/codeform/model"

// Query represents a code query.
type Query struct {
	// packages contains a list of packages to include.
	packages []string
	// names contains a slice of named things to include
	// regardless of kind (struct, interface, etc.)
	names []string
	// interfaces is a list of interfaces to include.
	interfaces []string
	// structs is a list of structs to include.
	structs []string
	// funcs is a list of global functions to include.
	funcs []string
}

// New creates a new Query.
func New() *Query {
	return &Query{}
}

// Run executes the query on the code, returning a new model.Code.
func (q *Query) Run(code model.Code) (*model.Code, error) {
	if len(q.packages)+len(q.names)+len(q.interfaces)+len(q.structs)+len(q.funcs) == 0 {
		// no querying - just return the code
		return &code, nil
	}
	var pkgs []model.Package
	for _, pkg := range code.Packages {
		include, err := q.runPackage(&pkg)
		if err != nil {
			return nil, err
		}
		if include {
			pkgs = append(pkgs, pkg)
		}
	}
	code.Packages = pkgs
	return &code, nil
}

// Name specifies names of things to query for.
// Names are always checked first, if the names match, the item
// will be included.
func (q *Query) Name(names ...string) *Query {
	q.names = append(q.names, names...)
	return q
}

// Interface specifies names of interfaces to include.
func (q *Query) Interface(names ...string) *Query {
	q.interfaces = append(q.interfaces, names...)
	return q
}

// Struct specifies names of structs to include.
func (q *Query) Struct(names ...string) *Query {
	q.structs = append(q.structs, names...)
	return q
}

// Package specifies names of packages to include.
func (q *Query) Package(names ...string) *Query {
	q.packages = append(q.packages, names...)
	return q
}

// Func specifies names of global functions to include.
func (q *Query) Func(names ...string) *Query {
	q.funcs = append(q.funcs, names...)
	return q
}

// runPackage runs the query on the specified package.
func (q *Query) runPackage(pkg *model.Package) (bool, error) {

	// packages
	if len(q.packages) > 0 {
		var packageMentioned bool
		for _, packageName := range q.packages {
			if pkg.Name == packageName {
				packageMentioned = true
				break
			}
		}
		if !packageMentioned {
			return false, nil
		}
	}

	var include bool

	// interfaces
	var ifaces []model.Interface
	for _, iface := range pkg.Interfaces {
		if q.shouldInterface(&iface) {
			ifaces = append(ifaces, iface)
			include = true
		}
	}
	pkg.Interfaces = ifaces

	// structs
	var structs []model.Struct
	for _, structure := range pkg.Structs {
		if q.shouldStruct(&structure) {
			structs = append(structs, structure)
			include = true
		}
	}
	pkg.Structs = structs

	// funcs
	var funcs []model.Func
	for _, fn := range pkg.Funcs {
		if q.shouldFunc(&fn) {
			funcs = append(funcs, fn)
			include = true
		}
	}
	pkg.Funcs = funcs

	return include, nil
}

// shouldInterface gets whether the interface should be included
// or not.
func (q *Query) shouldInterface(v *model.Interface) bool {
	if q.shouldInclude(v.Name) {
		return true
	}
	for _, iface := range q.interfaces {
		if v.Name == iface {
			return true
		}
	}
	return false
}

// shouldStruct gets whether the struct should be included
// or not.
func (q *Query) shouldStruct(v *model.Struct) bool {
	if q.shouldInclude(v.Name) {
		return true
	}
	for _, name := range q.structs {
		if v.Name == name {
			return true
		}
	}
	return false
}

// shouldFunc gets whether the struct should be included
// or not.
func (q *Query) shouldFunc(v *model.Func) bool {
	if q.shouldInclude(v.Name) {
		return true
	}
	for _, name := range q.funcs {
		if v.Name == name {
			return true
		}
	}
	return false
}

// shouldInclude gets whether the name is explicitly listed in the
// names slice or not.
func (q *Query) shouldInclude(name string) bool {
	for _, actualName := range q.names {
		if actualName == name {
			return true
		}
	}
	return false
}
