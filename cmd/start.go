/*
Copyright © 2024 Nicholas Rubright <nicholasrubright>
*/
package cmd

import (
	"github.com/nicholasrubright/oxyche/internal"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the proxy cache server",
	Run: func(cmd *cobra.Command, args []string) {
		portFlag, _ := cmd.Flags().GetInt("port")
		originFlag, _ := cmd.Flags().GetString("origin")

		internal.UpdateConfig(originFlag, portFlag)
		internal.BuildImage()
		internal.CreateAndStartContainer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Int("port", 3000, "Port for the server to be ran on")
	startCmd.Flags().String("origin", "", "Origin for the requests to be sent to")
}
