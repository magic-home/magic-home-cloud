package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Device struct {
	Id     string  `json:"id"`
	Boards []Board `json:"boards"`
}

type Board struct {
	Id      string   `json: "id"`
	Sensors []Entity `json:"sensors"`
	Switchs []Entity `json:"switches"`
}

type Entity struct {
	Id   string `json: "id"`
	Type string `json: "type"`
}

// Static list of devices for example
var devices = map[string]Device{

	"magic-home-sourish": Device{
		Id: "magic-home-sourish",
		Boards: []Board{
			Board{
				Id: "room1",
				Sensors: []Entity{
					Entity{
						Id:   "light-sensor-001",
						Type: "light-sensor",
					},
				},
				Switchs: []Entity{
					Entity{
						Id:   "switch-001",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-002",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-003",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-004",
						Type: "toggle-switch",
					},
				},
			},
			Board{
				Id: "room2",
				Switchs: []Entity{
					Entity{
						Id:   "switch-001",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-002",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-003",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-004",
						Type: "toggle-switch",
					},
				},
			},
		},
	},

	"magic-home-orijeet": Device{
		Id: "magic-home-orijeet",
		Boards: []Board{
			Board{
				Id: "room1",
				Switchs: []Entity{
					Entity{
						Id:   "switch-001",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-002",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-003",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-004",
						Type: "toggle-switch",
					},
				},
			},
			Board{
				Id: "room2",
				Switchs: []Entity{
					Entity{
						Id:   "switch-001",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-002",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-003",
						Type: "toggle-switch",
					},
					Entity{
						Id:   "switch-004",
						Type: "toggle-switch",
					},
				},
			},
		},
	},
}

// Get list of devices
func DeviceList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deviceList := make([]string, 0)
	for device, _ := range devices {
		deviceList = append(deviceList, device)
	}
	deviceListData := map[string][]string{"devices": deviceList}
	data, _ := json.Marshal(deviceListData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Get conf of a device
func DeviceConf(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device")
	device, found := devices[deviceId]
	if !found {
		log.Printf("invalid device Id: %s", deviceId)
		http.Error(w, fmt.Sprintf("invalid device Id: %s", deviceId), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(device)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	router := httprouter.New()
	router.GET("/devices", DeviceList)
	router.GET("/device/:device", DeviceConf)

	log.Fatal(http.ListenAndServe(":3311", router))
}
