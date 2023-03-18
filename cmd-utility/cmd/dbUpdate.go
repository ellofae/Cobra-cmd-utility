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

var FILL_DATA []dbWrapper

// dbUpdateCmd represents the dbUpdate command
var dbUpdateCmd = &cobra.Command{
	Use:   "dbUpdate",
	Short: "Insert data from the file into the db table",
	Long:  `Insert data from the file into the db table`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbUpdate called")
		fileName, _ := cmd.Flags().GetString("df")
		fileData := fileStoragePath + fileName + ".txt"
		dbFilePath := dbStoragePath + fileName + ".db"

		FILL_DATA, err := ReadingDataFromFile(fileData)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range FILL_DATA {
			fmt.Printf("%q\n", item)
		}

		db, err = sql.Open("sqlite3", dbFilePath)
		if err != nil {
			log.Fatalf("Didn't manage to open the .db file '%s'\n", dbFilePath)
		}

		InsertDataIntoDBFromFile()
	},
}

func init() {
	rootCmd.AddCommand(dbUpdateCmd)
	dbUpdateCmd.Flags().String("df", "none", "File with data to insert")
}

func ReadingDataFromFile(filename string) ([]dbWrapper, error) {

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Didn't manage to read data from file '%s'\n", filename)
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error while reading a line in the file '%s'\n", filename)
			return nil, err
		}

		lineSplited := strings.Split(line, " | ")
		for _, value := range lineSplited {
			value = strings.TrimRight(value, " ")
		}

		lastOne := lineSplited[3]
		lastOne = lastOne[:len(lastOne)-1]

		newDbWrapper := dbWrapper{
			Field1: lineSplited[0],
			Field2: lineSplited[1],
			Field3: lineSplited[2],
			Field4: lastOne,
		}
		FILL_DATA = append(FILL_DATA, newDbWrapper)

	}
	return FILL_DATA, nil
}

func InsertDataIntoDBFromFile() error {
	for _, dbData := range FILL_DATA {
		_, err := db.Exec("INSERT INTO data (f1, f2, f3, f4) VALUES (?,?,?,?)", dbData.Field1, dbData.Field2, dbData.Field3, dbData.Field4)
		if err != nil {
			return fmt.Errorf("Insertion error: %v", err)
		}
	}

	return nil
}
