package main

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"testing"
)


func TestGoMigrate(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	migrationScripts, err := filepath.Glob("test/migrations/*.sql")
	if err != nil {
		log.Fatal("Error migration file lookup")
	}

	// load envar
	err = godotenv.Load("test/.env")
	if err != nil {
		log.Fatal("Error loading .env file, missing db credentials")
	}

	// load DB
	db := GetDB(os.Getenv("DB_URL"))
	sort.Strings(migrationScripts)

	initMigration, _ := ioutil.ReadFile("test/migrations/0_create_module_up.sql")
	stmt, err := db.Prepare(string(initMigration))
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec()
	initMigration, _ = ioutil.ReadFile("test/migrations/1_create_migration_up.sql")
	stmt, err = db.Prepare(string(initMigration))

	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	// run migartions
	err = runMigrations(db, migrationScripts)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetDB(t *testing.T) {
	err := godotenv.Load("test/.env")
	if err != nil {
		log.Fatal("Error loading .env file, missing db credentials")
	}

	// load DB
	_ = GetDB(os.Getenv("DB_URL"))
}
