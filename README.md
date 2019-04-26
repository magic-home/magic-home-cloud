# magic-home-cloud
magic home cloud stack in docker swarm

## Overview
### Services
#### device-state
maintain device's console and device state
##### API
| API  | DESC  |
|------|-------|
| /device-state/<device_id>  | get device state of a device |
| /console-state/<device_id>  | get console state of a device  |
| /state/clean  |  reset device and console state |
> ** current API is for test, and subjected to change 

#### device-conf
store device configuration
##### API
| API  | DESC  |
|------|-------|
| /devices  | get the list of devices |
| /device/<device_id>  | get configuration of specified device |
> ** current API is for test, and subjected to change 

## Getting Started
#### Prerequisites
> [docker](https://www.docker.com/get-started)   
> [docker-swarm](https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/)   
#### Build & deploy
```bash
./build.sh
./deploy.sh
```
#### Configuration
##### Devices
```
magic-home-sourish
magic-home-orijeet
```
##### Current Entities 
###### magic-home-sourish
```json
{
  "id": "magic-home-sourish",
  "boards": [
    {
      "Id": "room1",
      "sensors": [
        {
          "Id": "light-sensor-001",
          "Type": "light-sensor"
        }
      ],
      "switches": [
        {
          "Id": "switch-001",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-002",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-003",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-004",
          "Type": "toggle-switch"
        }
      ]
    },
    {
      "Id": "room2",
      "sensors": null,
      "switches": [
        {
          "Id": "switch-001",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-002",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-003",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-004",
          "Type": "toggle-switch"
        }
      ]
    }
  ]
}
```
###### magic-home-orijeet
```
{
  "id": "magic-home-orijeet",
  "boards": [
    {
      "Id": "room1",
      "sensors": null,
      "switches": [
        {
          "Id": "switch-001",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-002",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-003",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-004",
          "Type": "toggle-switch"
        }
      ]
    },
    {
      "Id": "room2",
      "sensors": null,
      "switches": [
        {
          "Id": "switch-001",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-002",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-003",
          "Type": "toggle-switch"
        },
        {
          "Id": "switch-004",
          "Type": "toggle-switch"
        }
      ]
    }
  ]
}
```
##### Topics
```
// SRD's
srd-magic-home-sourish
srd-magic-home-orijeet

// SEC's
sec-magic-home-sourish
sec-magic-home-orijeet
```
#### Changing device state using SRD
* target topic is `srd-magic-home-sourish`   
* version should be incremental
* SRD data to change `magic-home-sourish` state   
```json
{
  "id" : "magic-home-sourish",
  "version" : 1,
  "states" : [
     {
        "board" : "room2",
        "entity" : "switch-001",
        "value" : "off"
     },
     {
        "board" : "room2",
        "entity" : "switch-002",
        "value" : "off"
     },
     {
        "board" : "room2",
        "entity" : "switch-003",
        "value" : "on"
     },
     {
        "board" : "room2",
        "entity" : "switch-004",
        "value" : "off"
     },
     {
        "board" : "room1",
        "entity" : "switch-001",
        "value" : "off"
     },
     {
        "board" : "room1",
        "entity" : "switch-002",
        "value" : "off"
     },
     {
        "board" : "room1",
        "entity" : "switch-003",
        "value" : "on"
     },
     {
        "board" : "room1",
        "entity" : "switch-004",
        "value" : "on"
     },
     {
        "board" : "room1",
        "entity": "light-sensor-001",
        "value": "30-100"
     }
  ]
}
```
#### Query device state
```
curl <docker-machine-ip>:3111/device/magic-home-sourish
```
