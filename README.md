# remember-clt

Work in progress: a command line tool for adding Google Calendar events. Written in Go. 

The tool utilizes the great package [cli](https://github.com/codegangsta/cli) written by [codegangsta](https://github.com/codegangsta).

The project serves as a personal introduction to Go and Google APIs. Some links that I'm using for reference include:

+ https://golang.org/doc/code.html
+ https://developers.google.com/google-apps/calendar/quickstart/go
+ https://developers.google.com/google-apps/calendar/create-events
+ https://developers.google.com/google-apps/calendar/concepts
+ https://github.com/codegangsta/cli/blob/master/README.md
+ https://github.com/google/google-api-go-client
+ https://godoc.org/google.golang.org/api/calendar/v3
+ https://github.com/peterkshultz/google-calendar-cli

The biggest issue right now is trying to get the "add" method functioning. The API call to create an event is not recognized as valid by the go compiler. If anyone knows of a solution to this, let me know through some medium of communication! I'll be taking a look at the [godoc API documentation](https://godoc.org/google.golang.org/api/calendar/v3) as a next step.

Update, 09/02/2015: Paradoxically, Google has very little Calendar API support for Go. The best tutorials are with Java, Python, and Ruby. I'm going to look into other cli-like packages. There may simply not be enough documentation.

Update, 09/07/2015: I found [another project on GitHub](https://github.com/pocke/google-calendar-cli) that does what I'm trying to do. I'm going to continue building out this project to learn about golang!
