package pwnpatrolmain

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func RangeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// The prefix must be 5 valid hexa chars
	prefix := strings.ToUpper(ps.ByName("prefix"))

	matched, err := regexp.MatchString("^[A-F0-9]{5}$", prefix)
	if err != nil {
		logger.Fatal(err)
	}

	if !matched {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "The hash prefix was not in a valid format")
		return
	}

	// Ping the DB
	err = DB.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	// Query the prefix
	logger.Debug("Querying prefix:", prefix)
	rows, err := DB.Query("SELECT hex(hash), occurrences FROM pwned WHERE prefix = ?", prefix)
	if err != nil {
		logger.Fatal("invalid query: ", err)
	}
	defer rows.Close()

	// Return all suffixes. If we got here, there should never be an
	// empty response.
	for rows.Next() {
		var hash string
		var occurrences int

		err = rows.Scan(&hash, &occurrences)
		if err != nil {
			logger.Fatal(err)
		}
		fmt.Fprintf(w, "%s:%d\n", hash[5:], occurrences)

	}

	err = rows.Err()
	if err != nil {
		logger.Fatal(err)
	}

}
