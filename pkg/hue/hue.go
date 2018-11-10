package hue

import (
	"encoding/json"
	"fmt"
	"hueScenes4openHAB/pkg/types"
	"net/http"
)

type Hue interface {
	GetGroups() (types.HueGroups, error)
	GetScenes() (types.HueScenes, error)
}

type hue struct {
	config *types.Config
}

func NewHue(config *types.Config) Hue {
	hue := hue{
		config: config,
	}
	return &hue
}

func (h hue) GetGroups() (types.HueGroups, error) {
	var groups types.HueGroups
	err := h.getAndParseResource("groups", &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (h hue) GetScenes() (types.HueScenes, error) {
	var scenes types.HueScenes
	err := h.getAndParseResource("scenes", &scenes)
	if err != nil {
		return nil, err
	}

	return scenes, nil
}

func (h hue) getAndParseResource(resource string, target interface{}) error {
	res, err := http.Get(h.getHueUrl(resource))
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func (h hue) getHueUrl(resource string) string {
	return fmt.Sprintf("http://%s/api/%s/%s", h.config.HueIP, h.config.HueUser, resource)
}
