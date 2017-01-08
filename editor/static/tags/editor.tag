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
	</style>
	<div class='template'>
		<codebox ref='templatebox' placeholder="Template" class='box'></codebox>
	</div>
	<div class='preview'>
		<div class="ui source accordion">
			<div class="title">
			<i class="dropdown icon"></i>
				Source
			</div>
			<div class="content">
				<p>
					This source code will be used to render the preview:
				</p>
				<codebox ref='sourcebox' placeholder="Go source code" class='box'></codebox>
			</div>
		</div>
		<div name='messagebox' class='ui message' show={ message }>
			{ message }
		</div>
		<codebox ref='previewbox' placeholder="The preview will render here once you have entered some source, and a template." class='box' readonly='true' busy={ busy } showbusy='true'></codebox>
	</div>
	<script>

		this.on('mount', function(){
			$('.accordion').accordion()
			this.messageBox = $('#messagebox', this.root)
			this.refs.templatebox.on('change', this.render)
			this.refs.sourcebox.on('change', this.render)

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
			this.message = data.message
			this.update()
			this.messageBox.hide().fadeIn()
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