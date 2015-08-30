package main

import (
//  "encoding/json"
  "fmt"
//  "io/ioutil"
//  "log"
//  "net/http"
//  "net/url"
  "os"
//  "os/user"
//  "path/filepath"
//  "time"

//  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/calendar/v3"
  "github.com/codegangsta/cli"
)


func main() {
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
  	

	calendarService, err := calendar.New(oauthHttpClient)

        event := &calendar.Event{
      	
        Summary: c.Args().First(),
      	Location: "",
      	Description: c.Args().First(),
      	Start: &calendar.EventDateTime{
            DateTime: "2016-05-28T09:00:00-07:00",
        	  TimeZone: "America/Detroit",
      	},

      	End: &calendar.EventDateTime{
            DateTime: "2016-05-28T17:00:00-07:00",
        		TimeZone: "America/Detroit",
      		},
      		
    	  }

    	calendarId := "primary"
    	
	event, err = srv.Events.Insert(calendarId, event).Do()

	if err != nil {
  		
		log.Fatalf("Unable to create event. %v\n", err)

	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)

      },
    },

  }

  app.Run(os.Args)
}
