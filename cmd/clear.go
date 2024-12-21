/*
Copyright © 2024 Nicholas Rubright <nicholasrubright>
*/
package cmd

import (
	"github.com/nicholasrubright/oxyche/internal"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the cache for the currently running server",

	Run: func(cmd *cobra.Command, args []string) {
		internal.ClearCache()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
