package main

import (
	"flag"
	"log"
	"os"

	"github.com/mongj/gds-onecv-swe-assignment/internal/config"
	"github.com/mongj/gds-onecv-swe-assignment/internal/database"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "seedDB":
		seedDB()
	case "migrateDB":
		migrateDB(migrate.Up)
	case "rollbackDB":
		migrateDB(migrate.Down)
	default:
		_ = errors.Errorf("Unknown command: %s", flag.Arg(0))
	}
}

// seedDB seeds the database with the seed.sql file
func seedDB() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error loading config"))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error connecting to database"))
	}

	// Get seed sql file
	seedFilePath := config.RuntimeWorkingDirectory + "/seeds/seed.sql"
	q, err := os.ReadFile(seedFilePath)
	if err != nil {
		log.Fatalf("Failed to read seed file '%s': %v\n", seedFilePath, err)
	}
	if err := db.Exec(string(q)).Error; err != nil {
		log.Fatalf("Failed to seed database: %v\n", err)
	}

	// Seed the database
	log.Print("Successfully seeded database")
}

// migrateDB migrates the database in the specified direction
// for now all available migrations will be applied/rolled back
func migrateDB(direction migrate.MigrationDirection) {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error loading config"))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error connecting to database"))
	}

	migrations := &migrate.FileMigrationSource{Dir: "migrations"}

	numSteps := 0 // Apply all migrations for now
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v\n", err)
	}

	steps, err := migrate.ExecMax(sqlDB, "postgres", migrations, direction, numSteps)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	if direction == migrate.Up {
		log.Printf("Applied %d migration(s)!\n", steps)
	} else {
		log.Printf("Rolled back %d migration(s)!\n", steps)
	}
}
