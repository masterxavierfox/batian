package handlers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/ishuah/batian/models"
	"github.com/ishuah/batian/engines"
	"github.com/gorilla/websocket"
)

type Message struct {
    AppID   string
    Type 	string
    Duration	int
}

var (
	upgrader = websocket.Upgrader{}
)

func HandleWebSocket(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print(err)
			return
		}
		
		defer c.Close()

		apps, err := getAllApps(db)
		if err != nil {
			log.Print(err)
		}
		sendData(c, apps, "allapps")

		for {
			_, msgBytes, err := c.ReadMessage()
			if err != nil {
				log.Print(err)
				return
			}

			var msg Message
			json.Unmarshal(msgBytes, &msg)

			switch msg.Type {
				case "appAnalysis":
					var events models.Events
					events, err = db.GetAppEvents(msg.AppID, msg.Duration)
					if err != nil {
						if err.Error() == "not found" {
							response, _ := json.Marshal(struct { string `json:"error"`}{"No events in the given time window."})
							c.WriteMessage(websocket.TextMessage, response)
							return
						}
						response, _ := json.Marshal(struct { string `json:"error"`}{"Query: General Failure"})
						c.WriteMessage(websocket.TextMessage, response)
						return
					}
					report, err := engines.AppAnalysis(events)
					data, err := json.Marshal(struct { 
						engines.Report `json:"data"`
						string `json:"type"`
						}{report, "appAnalysis"})
					err = c.WriteMessage(websocket.TextMessage, data)
					if err != nil {
						log.Print(err)
						return
					}
				default:
					panic("Unrecognized option")
			}
			
			fmt.Printf("%s\n", msgBytes)
		}
		})
}

func sendData(c *websocket.Conn, apps models.Apps, messageType string) {
	data, err := json.Marshal(struct { 
					models.Apps `json:"data"`
					string `json:"type"`
					}{apps, messageType})

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Print(err)
		return
	}
}

func getAllApps(db *models.DbManager) (models.Apps, error){
	apps, err := db.AllApps()
	return apps, err
}