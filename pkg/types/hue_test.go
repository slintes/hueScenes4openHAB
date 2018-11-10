package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseGroups(t *testing.T) {

	groupsJson := `
	{
    "1": {
        "name": "Group 1",
        "lights": [
            "1",
            "2"
        ],
        "type": "LightGroup",
        "action": {
            "on": true,
            "bri": 254,
            "hue": 10000,
            "sat": 254,
            "effect": "none",
            "xy": [
                0.5,
                0.5
            ],
            "ct": 250,
            "alert": "select",
            "colormode": "ct"
        }
    },
    "2": {
        "name": "Group 2",
        "lights": [
            "3",
            "4",
            "5"
        ],
        "type": "LightGroup",
        "action": {
            "on": true,
            "bri": 153,
            "hue": 4345,
            "sat": 254,
            "effect": "none",
            "xy": [
                0.5,
                0.5
            ],
            "ct": 250,
            "alert": "select",
            "colormode": "ct"
        }
    }
	}
	`

	var hueGroups HueGroups
	if err := json.Unmarshal([]byte(groupsJson), &hueGroups); err != nil {
		t.Errorf("unmarshal groups failed: %v", err)
	}
	fmt.Printf("groups: %v", hueGroups)

}

func TestParseScenes(t *testing.T) {

	scenesJson := `
	{
	"4e1c6b20e-on-0": {
		"name": "Kathy on 1449133269486",
		"lights": ["2", "3"],
		"owner": "ffffffffe0341b1b376a2389376a2389",
		"recycle": true,
		"locked": false,
		"appdata": {},
		"picture": "",
		"lastupdated": "2015-12-03T08:57:13",
		"version": 1
	},
	"3T2SvsxvwteNNys": {
		"name": "Cozy dinner",
		"lights": ["1", "2"],
		"owner": "ffffffffe0341b1b376a2389376a2389",
		"recycle": true,
		"locked": false,
		"appdata": {
			"version": 1,
			"data": "myAppData"
		},
		"picture": "",
		"lastupdated": "2015-12-03T10:09:22",
		"version": 2
	}
	}
	`

	var hueScenes HueScenes
	if err := json.Unmarshal([]byte(scenesJson), &hueScenes); err != nil {
		t.Errorf("unmarshal scenes failed: %v", err)
	}
	fmt.Printf("scenes: %v", hueScenes)

}
