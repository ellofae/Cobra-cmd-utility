/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const dbStoragePath string = `../db-storage/`

var db *sql.DB

// dbSettingCmd represents the dbSetting command
var dbSettingCmd = &cobra.Command{
	Use:   "dbSetting",
	Short: "Command supplier the user with functions to manage availabe data in the project",
	Long:  `Command supplier the user with functions to manage availabe data in the project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database managing system called:")
		availableFunctionality()

		dbCommand, _ := cmd.Flags().GetString("ops")
		dbName, _ := cmd.Flags().GetString("dbName")

		switch dbCommand {
		case "list-db":
			fmt.Printf("Available databases: (%s)\n", dbStoragePath)
			entries, err := os.ReadDir(dbStoragePath)
			if err != nil {
				log.Fatal(err)
			}

			for _, item := range entries {
				fmt.Printf("\t.db file: %v\n", item.Name())
			}
		case "show":
			if dbName != "none" {
				//dbFilePath := dbStoragePath
				//db, err := sql.Open("sqlite3")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbSettingCmd)
	dbSettingCmd.Flags().String("ops", "list-db", "Options for db managment")
	dbSettingCmd.Flags().String("dbName", "none", "Name of the specific database")
}

func availableFunctionality() {
	fmt.Println("Welcome to the db manager. Available options:")
	fmt.Println("\t--ops list-db (default) - prints out the list of all available databases")
	fmt.Println("\t--ops show --dbName {name of the db file} - print out the db data")
	fmt.Println()
}
