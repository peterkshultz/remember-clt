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

Update, 09/10/2015: utilizing another [GitHub user's repository](https://github.com/pocke/google-calendar-cli) for reference will prove helpful. Of course, there are still some kinks to work around---golang isn't as well documented of a language as some others, especially on websites like StackOverflow.
