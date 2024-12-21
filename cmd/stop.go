/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nicholasrubright/oxyche/internal"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the proxy cache server",
	Run: func(cmd *cobra.Command, args []string) {
		internal.StopContainer()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
