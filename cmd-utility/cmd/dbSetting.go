/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dbSettingCmd represents the dbSetting command
var dbSettingCmd = &cobra.Command{
	Use:   "dbSetting",
	Short: "Command supplier the user with functions to manage availabe data in the project",
	Long:  `Command supplier the user with functions to manage availabe data in the project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbSetting called")

	},
}

func init() {
	rootCmd.AddCommand(dbSettingCmd)

}
