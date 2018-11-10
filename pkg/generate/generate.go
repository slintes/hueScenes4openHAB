package generate

import (
	"hueScenes4openHAB/pkg/hue"
	"hueScenes4openHAB/pkg/types"

	log "github.com/sirupsen/logrus"
)

type Generator interface {
	Generate() error
}

type generator struct {
	config *types.Config
	hue    hue.Hue
}

func NewGenerator(config *types.Config) Generator {
	g := generator{
		config: config,
		hue:    hue.NewHue(config),
	}
	return &g
}

func (g *generator) Generate() error {
	log.Infof("let's generate all the things!")
	log.Debugf("config: %+v", *g.config)

	log.Infof("requesting groups and scenes from Hue bridge")

	hueGroups, err := g.hue.GetGroups()
	if err != nil {
		return err
	}
	log.Debugf("groups: %+v", hueGroups)
	log.Infof("got %v groups", len(hueGroups))

	hueScenes, err := g.hue.GetScenes()
	if err != nil {
		return err
	}
	log.Debugf("scenes: %+v", hueScenes)
	log.Infof("got %v szenes", len(hueScenes))

	log.Errorf("sorry, this is work in progress, and that's it for now, please check back later :)")

	return nil
}
