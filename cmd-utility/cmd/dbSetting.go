/*
Copyright © 2023 ELLOFAE <sergei.bykovskiy2003@gmail.com>
*/
package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

const dbStoragePath string = `../db-storage/`
const fileStoragePath string = `../file-storage/`

var db *sql.DB

type dbWrapper struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 string
}

var DATA []dbWrapper

// dbSettingCmd represents the dbSetting command
var dbSettingCmd = &cobra.Command{
	Use:     "dbSetting",
	Aliases: []string{"dbInfo", "dbState"},
	Short:   "Command supplier the user with functions to manage availabe data in the project",
	Long:    `Command supplier the user with functions to manage availabe data in the project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database managing system called:")
		availableFunctionality()

		dbCommand, _ := cmd.Flags().GetString("ops")
		dbName, _ := cmd.Flags().GetString("dbName")
		dbFilePath := dbStoragePath + dbName
		var err error

		db, err = sql.Open("sqlite3", dbFilePath)
		if err != nil {
			log.Fatalf("Didn't manage to open the .db file '%s'\n", dbFilePath)
		}

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
				nameSplit := strings.Split(dbName, ".")
				fileName := fileStoragePath + nameSplit[0] + ".txt"

				f, err := os.Open(fileName)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				amountOfColumns := 0
				reader := bufio.NewReader(f)
				for {
					line, err := reader.ReadString('\n')
					if err == io.EOF {
						break
					} else if err != nil {
						log.Fatal(err)
					}

					lineSplit := strings.Split(line, " | ")
					amountOfColumns = len(lineSplit)
					break
				}

				if amountOfColumns == 4 {
					DATA, err = printDatabaseTableData(amountOfColumns)
					for _, item := range DATA {
						fmt.Printf("%q\n", item)
					}
				} else {
					log.Fatal("Utility can print out the db only with 4 columns")
				}

			} else {
				fmt.Println("Pass name of the database to the option --dbName {name.db}")
				break
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
	fmt.Println("\t--ops show --dbName {name.db} - print out the db data")
	fmt.Println()
}

func printDatabaseTableData(colms int) ([]dbWrapper, error) {
	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var temp1 string
	var temp2 string
	var temp3 string
	var temp4 string

	if colms == 4 {
		for rows.Next() {
			err = rows.Scan(&temp1, &temp2, &temp3, &temp4)
			temp := dbWrapper{temp1, temp2, temp3, temp4}
			DATA = append(DATA, temp)
		}
	}

	return DATA, nil
}
