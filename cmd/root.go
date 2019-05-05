package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rpi-go-led-sign",
	Short: "A simple Neopixel animated sign simulator and driver",
	Long: `A simple neopixel driver in go to create an animated sign, also simulates
said sign using rpi graphics if available`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute does the thing
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
