/*
Copyright © 2024 Pavlo Tarasiuk <pasha.tarasyuk@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion = "Version"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of MAVBot",
	Long: `The version command displays the current version of MAVBot,
	helping to identify the software version when interacting with users or other systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
