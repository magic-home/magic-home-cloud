package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	httprouter "github.com/julienschmidt/httprouter"
)

type EntityValue struct {
	EntityId string   `json: entity`
	Value    struct{} `json: state`
}

// Represent Entity State
type EntityState struct {
	Board string `json: "board"`
	EntityValue
	//	EntityId string   `json: "entityId"`
	//	Value    struct{} `json: "value"`
}

// SCE: State Chnage Event
type SCE struct {
	Id      string `json: "id"`      // Device Id
	Version int    `json: "version"` // SEC Verison
	Type    string `json: "type"`

	Event *EntityState `json: event`
}

// SRD: State Replication Data
type SRD struct {
	Id      string `json: "id"`      // Device Id
	Version int    `json: "version"` // SRD Verison

	Entities []*EntityState `json: states` // entity states
}

// DeviceState: states of the device
type DeviceState struct {
	Version int `json: "version"`

	Boards map[string][]*EntityValue `json: baards`
}

var consoleState map[string]*DeviceState
var consolemux sync.Mutex

var deviceState map[string]*DeviceState
var devicemux sync.Mutex

// Get conf of a device
func GetDeviceState(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device")

	device, ok := deviceState[deviceId]
	if !ok {
		log.Printf("invalid device Id: %s", deviceId)
		http.Error(w, fmt.Sprintf("invalid device Id: %s", deviceId), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(device)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Get conf of a device
func GetConsoleState(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device")

	device, ok := consoleState[deviceId]
	if !ok {
		log.Printf("invalid device Id: %s", deviceId)
		http.Error(w, fmt.Sprintf("invalid device Id: %s", deviceId), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(device)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Clear the state of the console
func ClearState(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deviceState = make(map[string]*DeviceState)
	consoleState = make(map[string]*DeviceState)

	response := map[string]string{"state": "ok"}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	// opts.SetUsername(uri.User.Username())
	// password, _ := uri.User.Password()
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

// mqttConnect connect to a mqtt broker
func mqttConnect(topic string) mqtt.Client {
	brokerUrl := os.Getenv("MQTT_BROKER_URL") + "/" + topic
	log.Printf("Connecting %s", brokerUrl)
	uri, err := url.Parse(brokerUrl)
	if err != nil {
		log.Fatal(err)
	}

	opts := createClientOptions(topic+"-sub", uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

// ComsumeSRD consumes and process SRD from given token
func consumeSRD(topic string) {
	client := mqttConnect(topic)

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {

		payloadSRD := msg.Payload()
		srd := &SRD{}

		err := json.Unmarshal(payloadSRD, srd)
		if err != nil {
			log.Printf("Faield to parse SRD payload for topoc %s, err %v", topic, err)
		}

		deviceId := srd.Id

		devicemux.Lock()
		device, found := deviceState[deviceId]
		if !found {
			device = &DeviceState{Version: srd.Version}
			deviceState[deviceId] = device
		} else {
			if device.Version >= srd.Version {
				devicemux.Unlock()
				return
			}
		}

		// Flush the state values
		device.Boards = make(map[string][]*EntityValue)

		for _, entity := range srd.Entities {

			board, ok := device.Boards[entity.Board]
			if !ok {
				board = make([]*EntityValue, 0)
			}
			device.Boards[entity.Board] = append(board, &entity.EntityValue)
		}

		devicemux.Unlock()
	})
}

// ComsumeSEC consumes and process SEC from given token
func consumeSEC(topic string) {
	client := mqttConnect(topic)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {

	})
}

// Start listeners for all devices
func StartListeners() {
	// TODO: fetch list of devices from device-conf
	// TODO: maintain max no of Listenenr
	// TODO: maintain least recently used using sec request

	go consumeSRD("srd-magic-home-sourish")
	//	mqttConnect("sec-magic-home-sourish", "sec")

	go consumeSRD("srd-magic-home-orijeet")
	//	mqttConnect("sec-magic-home-orijeet", "sec")
}

func main() {

	StartListeners()

	// Start rest service
	router := httprouter.New()

	router.GET("/device-state/:device", GetDeviceState)
	router.GET("/console-state/:device", GetConsoleState)
	router.GET("/state/clean", ClearState)

	log.Fatal(http.ListenAndServe(":3312", router))
}
