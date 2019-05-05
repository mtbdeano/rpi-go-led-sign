package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(guiCmd)
}

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Show a simulated LED sign",
	Long:  `Simple WIDTH x HEIGHT pixel sign simulator`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("%s", "Show something")
	},
}
