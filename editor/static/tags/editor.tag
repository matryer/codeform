<editor>
	<style>
		.template {
			display: flex;
			flex: 1 100%;
			flex-direction: column;
		}
		.box {
			display: flex;
			flex: 1 100%;
		}
		.preview {
			display: flex;
			flex: 1 100%;
			flex-direction: column;
			margin-left: 20px;
		}
		.message {
			flex-shrink: 0;
		}
		.source {
			margin-bottom: 10px;
		}
		.tabular.menu {
			flex-shrink: 0;
		}
		.ui.top.attached.tabular.menu {
			margin: 0px;
		}
		.ui.bottom.attached.tab.active.segment {
			display: flex;
			flex: 1 100%;
			margin: 0px;
		}
		.tab-body {
			display: flex;
			flex: 1 100%;
			flex-direction: column;
			align-items: stretch;
			align-content: stretch;
		}

		iframe {
			border: 0;
			display: flex;
			flex: 1;
		}

	</style>
	<div class='template'>
		<codebox ref='templatebox' placeholder='Template' class='box'></codebox>
	</div>
	<div class='preview'>
		<div name='messagebox' class='ui message' show={ message }>
			{ message }
		</div>
		<div class="ui top attached tabular menu">
			<a class="item" data-tab="source">Source</a>
			<a class="active item" data-tab="preview">Preview</a>
			<a class="item" data-tab="docs">Docs</a>
		</div>
		<div class="ui bottom attached tab segment" data-tab="source">
			<div class='tab-body'>
				<p>
					Source code will be used to render the preview - must be valid Go code
				</p>
				<codebox ref='sourcebox' placeholder="Go source code" class='box'></codebox>
			</div>
		</div>
		<div class="ui bottom attached tab active segment" data-tab="preview">
			<div class='tab-body'>
				<codebox ref='previewbox' placeholder="The preview will render here once you have entered some source, and a template." class='box' readonly='true' busy={ busy } showbusy='true'></codebox>
			</div>
		</div>
		<div class="ui bottom attached tab segment" data-tab="docs">
			<div class='tab-body'>
				<div class="ui three item secondary link menu">
					<a class='active item' href='https://godoc.org/github.com/matryer/codeform/model' target='docs'>Model API</a>
					<a class='item' href='https://godoc.org/github.com/matryer/codeform/render' target='docs'>Template utilities</a>
					<a class='item' href='https://matryer.github.io/codeform/examples.html' target='docs'>Examples</a>
				</div>
				<iframe id='docs' src='https://godoc.org/github.com/matryer/codeform/model'></iframe>
			</div>
		</div>

	</div>
	<script>

		this.on('mount', function(){
			$('.accordion').accordion()
			this.messageBox = $('#messagebox', this.root)
			this.refs.templatebox.on('change', this.render)
			this.refs.sourcebox.on('change', this.render)
			$('.menu .item', this.root).tab()
			this.items = $('.link.menu .item', this.root).click(function(e){
				this.items.removeClass('active')
				$(e.target).addClass('active')
			}.bind(this))
			this.defaultsCounter = 0;
			$.ajax({
				type: "get",
				url: "/default-source",
				success: function(code){
					this.refs.sourcebox.val(code)
					this.defaultsCounter++
					this.checkDefaultsLoaded()
				}.bind(this)
			})
			$.ajax({
				type: "get",
				url: "/default-template",
				success: function(code){
					this.refs.templatebox.val(code)
					this.defaultsCounter++
					this.checkDefaultsLoaded()
				}.bind(this)
			})

		}.bind(this))

		checkDefaultsLoaded() {
			if (this.defaultsCounter < 2) { return }
			this.render()
		}

		app.on('glance', function(data) {
			console.info(data.message)
		}.bind(this))

		render() {
			this.busy = true
			this.update()
			var data = {
				source: this.refs.sourcebox.val(),
				template: this.refs.templatebox.val()
			}
			$.ajax({
				type: 'post',
				url: "/preview",
				data: JSON.stringify(data),
				success: function(data){
					this.refs.previewbox.val(data.output)
					this.message = null
					this.update()
				}.bind(this),
				error: function(res){
					var msg = "Something went wrong"
					if (res.statusText) {
						msg = "Something weng wrong: " + res.statusText
					}
					if (res.responseJSON && res.responseJSON.error) {
						msg = res.responseJSON.error
					}
					this.message = msg
					this.update()
				}.bind(this),
				complete: function(){
					this.busy = false
					this.update()
				}.bind(this)
			})
		}

	</script>
</editor>