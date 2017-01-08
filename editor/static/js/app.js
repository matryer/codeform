var app = riot.observable()

app.counter = 1;

app.uniqueValue = function() {
	return app.counter++
}