package pwnpatrolmain

import (
	"bufio"
	"os"
	"strings"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectDB() {
	// Connect to the sqlite DB file
	dbFile := viper.GetString("dbFile")

	logger.Info("Connecting to database file: ", dbFile)
	var err error

	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		logger.Fatal(err)
	}

}

func InitDB() {
	logger.Info("Initializing database file")
	return

	sqlStmt := `
    DROP TABLE IF EXISTS pwned;
    CREATE TABLE pwned (prefix text, suffix text, occurrences integer);
    CREATE INDEX pwned_prefix_idx on pwned (prefix);

    `
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		logger.Errorf("%q: %s\n", err, sqlStmt)
		return
	}

	// populate pwned table
	tx, err := DB.Begin()
	if err != nil {
		logger.Fatal(err)
	}

	dumpfile, err := os.Open(viper.GetString("dumpFile"))
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer dumpfile.Close()

	scanner := bufio.NewScanner(dumpfile)

	stmt, err := tx.Prepare("INSERT INTO pwned(prefix, suffix, occurrences) VALUES (?, ?, ?)")
	if err != nil {
		logger.Fatal(err)
	}
	defer stmt.Close()

	prefix := ""

	logger.Info("Loading rows. This should take a few minutes.")

	for scanner.Scan() {
		line := scanner.Text()

		//fmt.Println(line)

		tokens := strings.Split(line, ":")

		_, err = stmt.Exec(tokens[0][:5], tokens[0][5:], tokens[1])
		if err != nil {
			logger.Fatal(err)
		}

		if tokens[0][:2] != prefix {
			prefix = tokens[0][:2]
			logger.Debug(prefix)
		}

	}

	// Commit the transaction
	tx.Commit()

	logger.Info("Done.")

}
