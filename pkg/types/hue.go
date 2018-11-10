package types

type HueGroups map[int]HueGroupParams

type HueGroupParams struct {
	Name   string   `json:"name"`
	Lights []string `json:"lights"`
	Type   string   `json:"type"`
	Action struct {
		On        bool      `json:"on"`
		Bri       int       `json:"bri"`
		Hue       int       `json:"hue"`
		Sat       int       `json:"sat"`
		Effect    string    `json:"effect"`
		Xy        []float64 `json:"xy"`
		Ct        int       `json:"ct"`
		Alert     string    `json:"alert"`
		Colormode string    `json:"colormode"`
	} `json:"action"`
}

type HueScenes map[string]HueSceneParams

type HueSceneParams struct {
	Name    string   `json:"name"`
	Lights  []string `json:"lights"`
	Owner   string   `json:"owner"`
	Recycle bool     `json:"recycle"`
	Locked  bool     `json:"locked"`
	Appdata struct {
		Version int    `json:"version"`
		Data    string `json:"data"`
	} `json:"appdata"`
	Picture     string `json:"picture"`
	Lastupdated string `json:"lastupdated"`
	Version     int    `json:"version"`
}
