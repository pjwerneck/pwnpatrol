package pwnpatrolmain

import (
	"encoding/hex"
	"bufio"
	"os"
	"strings"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB(dbFile string) {
	// Connect to the sqlite DB file
	logger.Info("Connecting to database file: ", dbFile)
	var err error

	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		logger.Fatal(err)
	}

}

func InitDB(dumpFile string) {

	logger.Infof("Initializing database")

	sqlStmt := `
    DROP TABLE IF EXISTS pwned;
    CREATE TABLE pwned (prefix text, hash blob, occurrences integer);
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
	logger.Infof("Loading data from dump file at %s", dumpFile)

	fd, err := os.Open(dumpFile)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	stmt, err := tx.Prepare("INSERT INTO pwned(prefix, hash, occurrences) VALUES (?, ?, ?)")
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

		occurrences := tokens[1]

		hash, err := hex.DecodeString(tokens[0])


		_, err = stmt.Exec(tokens[0][:5], hash, occurrences)
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

	// Index and vacuum the DB
	logger.Info("Indexing...")

	_, err = DB.Exec("CREATE INDEX pwned_prefix_idx on pwned (prefix);")
	if err != nil {
		logger.Errorf("%q: %s\n", err)
		return
	}


	logger.Info("Done.")

}
