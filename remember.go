package main

import (

    "fmt"
    "os"
    "github.com/peterkshultz/remember/client_secret.json"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/calendar/v3"
    "github.com/codegangsta/cli"

)


func main() {

    config := &oauth2.Config {

      ClientID:     "406595811286-kntrl98g7aujgsesob2oi6sfelsqlh4n.apps.googleusercontent.com",
      ClientSecret: "ypjZIJwF1rnfkn-MqIzQPa3i",
      Endpoint:     google.Endpoint,
      Scopes:       []string{calendar.CalendarScope},
    }

    ctx := context.Background()

    c := newOAuthClient(ctx, config)

    svc, err := calendar.New(c)

    if err != nil {

      panic(err)

    }

    app := cli.NewApp()
    app.Name = "remember"
    app.Usage = "utilize the awesome UI that is the command line to create Google Calendar events. Maybe if you use it you'll put your life in order :)"
    app.Action = func(c *cli.Context) {
    
      println(c.Args().First())

    }

  app.Commands = []cli.Command{
  {
    
    Name: "add",
    Usage: "add an event to the calendar",
    Action: func(c *cli.Context) {
  	

	     calendarService, err := calendar.New("https://www.googleapis.com/auth/calendar")

       event := &calendar.Event{
      	
       Summary: c.Args().First(),
       Location: "",
       Description: c.Args().First(),
       Start: &calendar.EventDateTime {

            DateTime: "2016-05-28T09:00:00-07:00",
        	  TimeZone: "America/Detroit",

       },

       End: &calendar.EventDateTime{

            DateTime: "2016-05-28T17:00:00-07:00",
        		TimeZone: "America/Detroit",

      		},
      		
    	 }

    	 calendarId := "primary"
    	
	event, err = srv.events.insert(calendarId, event).Do()

	if err != nil {
  		
		log.Fatalf("Unable to create event. %v\n", err)

	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)

      },
    },

  }

  app.Run(os.Args)
}
