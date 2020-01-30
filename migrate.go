package main

import (
	"crypto/sha1"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"hash"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"
)

type Migration struct {
	ID      int       `db:"id"`
	UUID    uuid.UUID `db:"uuid"`
	Created time.Time `db:"time"`
	Name    string    `db:"name"`
	Hash    hash.Hash `db:"hash"`
}

var (
	rName = regexp.MustCompile(`\d_(create|update|delete)_[a-zA-Z]+_(up|down).sql`)
)

func runMigrations(db *sql.DB, migrationScripts []string) error {
	// upsert migration if no conflict on hash value of sql query of migration
	insertStmt, err := db.Prepare("INSERT INTO migration (name, hash) VALUES ($1, $2) ON CONFLICT DO NOTHING;")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	// execute each new migration
	for _, migrationScript := range migrationScripts {

		// get query and calculate its md5 checksum
		migrationQuery, err := ioutil.ReadFile(migrationScript)
		if err != nil {
			log.Fatal(err)
		}
		h := sha1.New().Sum(migrationQuery)

		// insert new migration
		migrationName := rName.FindString(migrationScript)

		// execute each migration
		_, err = db.Exec(string(migrationQuery))

		if err != nil {
			log.Fatal(err)
		}
		// record each migration
		result, err := insertStmt.Exec(migrationName, strconv.QuoteToASCII(string(h)))
		if err != nil {
			log.Fatal(err)
		}
		affected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		if affected  > 0 {
			log.Printf("😍 executing %s \n", migrationName)
		} else {
			log.Printf("😢 passing on %s \n", migrationName)
		}

	}

	return nil
}
