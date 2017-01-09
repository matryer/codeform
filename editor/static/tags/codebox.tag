<codebox>
	<style>

		.wrapper {
			display: flex;
			flex: 1 100%;
			flex-direction: column;
		}

		.buttons {
			flex-shrink: 0;
		}

		.wrapper textarea {
			flex-grow: 1;
		}
		.wrapper .busy {
			flex-grow: 1;
		}

		.codeboxtextarea {
			min-height: 1em !important;
			max-height: unset  !important;
			font-family: monospace;
			resize: none !important;
			border: none !important;
			top: -15px !important;
			background-color: #FFD !important;
		}

		.codeboxtextarea[readonly="readonly"] {
			background-color: #eee !important;
		}

	</style>
	<div class='wrapper'>
		<div class='busy ui basic segment' show={ busy && showbusy }>
			<div class="ui active loader"></div>
		</div>
		<textarea placeholder={placeholder} show={ (!busy && showbusy) || !showbusy } ref='actualbox' id={"box"+this.uniqueValue} class='codeboxtextarea' readonly={ readonly } onkeyup={ onkeyup }></textarea>
		<!--<div class="ui tiny icon buttons">
			<button id={"copybtn"+this.uniqueValue} class="ui copy button" data-clipboard-target={"#box"+this.uniqueValue}><i class="copy icon"></i></button>
		</div>-->
	</div>
	<script>

		this.on('mount', function(){
			this.actualBox = $(this.refs.actualbox)
			this.uniqueValue = app.uniqueValue()
			if (opts.readonly == 'true') {
				this.readonly = true
			}
			if (opts.showbusy == 'true') {
				this.showbusy = true
			}
			this.placeholder = opts.placeholder
			this.update()
			this.clipboard = new Clipboard('#copybtn'+this.uniqueValue)
			this.clipboard.on('success', function(){
				app.trigger('glance', {message: 'Copied into clipboard', icon: 'checkmark'})
			})
			this.clipboard.on('error', function(){
				app.trigger('glance', {message: 'Failed to copy - please do it manually', icon: 'remove'})
			})
		}.bind(this))

		onkeyup() {
			clearTimeout(this.onkeyupTimeout)
			this.onkeyupTimeout = setTimeout(this.onkeyup_hit, 500);
		}
		onkeyup_hit() {
			this.trigger('change')
		}

		val() {
			return this.actualBox.val.apply(this.actualBox, arguments)
		}

	</script>
</codebox>