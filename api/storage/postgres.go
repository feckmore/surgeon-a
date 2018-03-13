package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"git.arthrex.io/dschultz/surgeon-a/api/surgeon"
)

var (
	host     = "localhost"
	port     = "5432"
	username = "postgres"
	password = ""
	dbname   = "postgres"
)

var db *sql.DB
var dataSourceName string

// surgeonDBRepo satisfies the Repositor interface
type surgeonDBRepo struct {
	db *sql.DB
}

func init() {
	// populate db specific variables from environment
	if h := os.Getenv("DATABASE_HOST"); h != "" {
		host = h
	}
	if p := os.Getenv("DATABASE_PORT"); p != "" {
		port = p
	}
	if u := os.Getenv("DATABASE_USERNAME"); u != "" {
		username = u
	}
	if pw := os.Getenv("DATABASE_PASSWORD"); pw != "" {
		password = pw
	}
	dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	dataSourceName = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbname)
}

// NewSurgeonDBRepository returns object that implements the Repositor interface
func NewSurgeonDBRepository() surgeon.Repositor {
	log.Println("NewSurgeonRepository")
	repo := new(surgeonDBRepo)

	var err error
	// Loop until database is ready and pings back successfully
	for { //TODO: consider max time to wait for database
		log.Println("Connecting to database...", dataSourceName)
		// During startup inside of a container, must wait for database container to also be ready
		time.Sleep(5 * time.Second)

		db, err = sql.Open("postgres", dataSourceName)
		if err == nil {
			log.Println(err)
		}

		if err = db.Ping(); err == nil {
			break // successful, get out of the for loop
		}
		log.Println(err)
	}

	// in development environment, create the database, tables & seed
	if err := createDatabase(); err != nil {
		log.Println(err) // TODO: FATAL?
	}

	fmt.Println("Successfully connected to database")
	repo.db = db

	return repo
}
func (ur *surgeonDBRepo) Store(surgeon *surgeon.Surgeon) error {
	//TODO: IMPLEMENT
	return nil
}
func (ur *surgeonDBRepo) FindByID(id surgeon.ID) (*surgeon.Surgeon, error) {
	log.Println("surgeonDBRepo.FindByID:", id)
	surgeon := surgeon.Surgeon{}

	rows, err := ur.db.Query(`SELECT id, first_name, last_name, email FROM surgeons WHERE id = $1`, id) //TODO: populate all of the fields lazy bum
	if err != nil {
		log.Fatal("Error in Query:", err) // Log fatal just throws 502 Bad Gateway with text message "An invalid response was received from the upstream server"
	}
	defer rows.Close()
	for rows.Next() {
		// err := rows.Scan(&id, &name)
		err := rows.Scan(&surgeon.ID, &surgeon.Firstname, &surgeon.Lastname, &surgeon.Email) //TODO: populate all of the fields
		if err != nil {
			log.Fatal("Error in row scan:", err)
		}
		log.Println(surgeon)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Error in rows:", err)
	}

	return &surgeon, nil
}
func (ur *surgeonDBRepo) FindAll() []*surgeon.Surgeon {
	//TODO: implement
	return nil
}
func (ur *surgeonDBRepo) Close() {
	defer ur.db.Close()
	log.Println("Closing database")
}

// ***************
// VERY TEMPORARY... MOVE TO SEED / MIGRATION PROCESS?
func createDatabase() error {
	var err error

	development, err := strconv.ParseBool(os.Getenv("DEVELOPMENT"))

	if development {
		// TODO: determine if surgeon db is being used after being created... might just be adding tables & content to default db
		result, err := db.Exec(`CREATE DATABASE surgeon`)
		if err != nil {
			return err
		}

		log.Println("Create Database 'surgeon':", result)

		createTable := `
		DROP TABLE IF EXISTS "public"."surgeons";
		CREATE TABLE "public"."surgeons" (
			"id" int4 NOT NULL,
			"first_name" varchar COLLATE "default",
			"last_name" varchar COLLATE "default",
			"address1" varchar COLLATE "default",
			"latitude" float8,
			"longitude" float8,
			"created_at" timestamp(6) NOT NULL,
			"updated_at" timestamp(6) NOT NULL,
			"suffix" varchar COLLATE "default",
			"prefix" varchar COLLATE "default",
			"city" varchar COLLATE "default",
			"state" varchar COLLATE "default",
			"zip" varchar COLLATE "default",
			"country" varchar COLLATE "default",
			"phone" varchar COLLATE "default",
			"fax" varchar COLLATE "default",
			"email" varchar COLLATE "default",
			"website" varchar COLLATE "default",
			"photo" varchar COLLATE "default",
			"organization" varchar COLLATE "default",
			"address2" varchar COLLATE "default"
		)
		WITH (OIDS=FALSE);
		ALTER TABLE "public"."surgeons" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
	`
		result, err = db.Exec(createTable)
		if err != nil {
			log.Println(err)
		}

		addSurgeons := `
		BEGIN;
		INSERT INTO "public"."surgeons" VALUES ('1', 'Dilbert', 'Adam', '123 Main St.', '26.273729', '-81.785632', '2016-10-13 13:40:39.904475', '2016-10-13 13:40:39.904475', 'MD', 'Dr.', 'Naples', 'FL', '34110', 'United States', '239-555-5555', '239-555-5555', 'dilbert@adams.com', 'http://www.dilbert.com', null, 'Dilbert', '');
		COMMIT;
		`
		result, err = db.Exec(addSurgeons)
		if err != nil {
			log.Println(err)
		}

	}

	return err
}
