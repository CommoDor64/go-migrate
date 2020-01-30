package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var (
	dir     = flag.String("dir", "./pkg/database/migrations/", "directory containing migration files")
	envfile = flag.String("envfile", ".env", "env file")
	dbURL   = flag.String("dburl", "DATABASE_URL", "db url connection key name in env file")
)

func main() {
	flag.Parse()
	log.Printf("Running migration from folder %s using env file %s \n", *dir, *envfile)

	migrationScripts, err := filepath.Glob(*dir + "*.sql")
	if err != nil {
		log.Fatal("Error migration file lookup")
	}

	// load envar
	err = godotenv.Load(*envfile)
	if err != nil {
		log.Fatal("Error loading .env file, missing db credentials")
	}

	// load DB
	db := GetDB(os.Getenv(*dbURL))
	sort.Strings(migrationScripts)

	// run migartions
	err = runMigrations(db, migrationScripts)
	if err != nil {
		log.Fatal(err)
	}
}
