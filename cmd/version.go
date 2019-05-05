package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Version is set at link time
	Version = "UNKNOWN"
	// Branch is set at compile time
	Branch = "UNKNOWN"
	// Commit is set at compile time
	Commit = "UNKNOWN"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rpi-go-led-sign",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("rpi-go-led-sign: version %s, branch %s, commit %s", Version, Branch, Commit)
	},
}
