package cmd

import (
	"hueScenes4openHAB/pkg/generate"
	"hueScenes4openHAB/pkg/types"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hueIP    string
	hueUser  string
	filename string
	debug    bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&hueIP, "hueIP", "", "IP of the Hue bridge (mandatory)")
	rootCmd.PersistentFlags().StringVar(&hueUser, "hueUser", "", "User for the Hue Bridge (mandatory)")
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "huescenes", "Name without extension for the generated files")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")

	rootCmd.MarkPersistentFlagRequired("hueIP")
	rootCmd.MarkPersistentFlagRequired("hueUser")

	//log.SetFormatter(&log.JSONFormatter{})
	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hueScenes4openHAB",
	Short: "hueScenes4openHAB creates openHAB config files Hue scenes",
	Long:  `An opinionated openHAB config files generator for Hue scenes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if debug {
			log.SetLevel(log.DebugLevel)
		}
		config := types.Config{
			HueIP:    hueIP,
			HueUser:  hueUser,
			Filename: filename,
		}
		generator := generate.NewGenerator(&config)
		return generator.Generate()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("something went terrible wrong: %v", err)
	}
}
