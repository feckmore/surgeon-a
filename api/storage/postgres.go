package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/feckmore/surgeon-a/api/surgeon"
)

var db *sql.DB

// surgeonRepo satisfies the surgeon.Repositor interface by implementing Store, FindByID, FindAll & Close
type surgeonRepo struct {
	db *sql.DB
}

// NewSurgeonDBRepository returns object that implements the Repositor interface
func NewSurgeonDBRepository(scheme, host, port, connectToDBName, username, password, createDBName string, devEnv bool) (surgeon.Repositor, error) {
	log.Println("NewSurgeonDBRepository")

	repo := new(surgeonRepo)

	dataSourceName := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", scheme, username, password, host, connectToDBName)
	fmt.Println(dataSourceName)
	err := connectToServer(dataSourceName)
	if err != nil {
		return repo, err //TODO?
	}
	repo.db = db
	fmt.Println("db: ", db)
	if devEnv == true {
		err := createDatabase(createDBName)
		if err != nil {
			log.Println(err) // TODO: FATAL?
			return repo, err
		}
		err = seedDatabase(createDBName)
		if err != nil {
			log.Println(err) // ?
			return repo, err
		}
	}

	log.Println("Successfully connected to database")

	return repo, nil
}

func (repo *surgeonRepo) Store(surgeon *surgeon.Surgeon) error {
	//TODO: IMPLEMENT
	return nil
}

func (repo *surgeonRepo) FindByID(id surgeon.ID) (*surgeon.Surgeon, error) {
	log.Println("surgeonRepo.FindByID:", id)
	surgeon := surgeon.Surgeon{}

	rows, err := repo.db.Query(`
		SELECT id, firstname, lastname, email
		FROM surgeons
		WHERE id = $1
		`, id) //TODO: populate all of the fields lazy bum
	if err != nil {
		log.Fatal("Error in Query:", err) // Log fatal just throws 502 Bad Gateway with text message "An invalid response was received from the upstream server"
	}
	defer rows.Close()
	rows.Next()
	err = rows.Err()
	if err != nil {
		log.Fatal("Error in rows:", err)
	}

	err = rows.Scan(&surgeon.ID, &surgeon.Firstname, &surgeon.Lastname, &surgeon.Email) //TODO: populate all of the fields
	if err != nil {
		log.Fatal("Error in row scan:", err)
	}
	// surgeon.Addresses = repo.addressIDs(surgeon.ID)

	log.Println(surgeon)

	return &surgeon, nil
}

func (repo *surgeonRepo) FindAll() []*surgeon.Surgeon {
	//TODO: implement
	return nil
}

func (repo *surgeonRepo) Close() {
	defer repo.db.Close()
	log.Println("Closing database")
}

func connectToServer(dataSourceName string) error {
	var err error

	// Loop until database is ready and pings back successfully
	for { //TODO: consider max time to wait for database
		log.Println("Connecting to database server:", dataSourceName)
		// During startup inside of a container, must wait for database container to also be ready
		time.Sleep(5 * time.Second)

		db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			log.Println("error connecting to database server:", err)
		}
		fmt.Println("connected to database server: ", db)

		if err = db.Ping(); err == nil {
			log.Println("successfully pinged database")
			break // successful, get out of the for loop
		}
		log.Println(err)
	}

	return err
}

func createDatabase(name string) error {
	log.Println("Creating database ", name)

	createStatement := fmt.Sprintf("CREATE DATABASE %v", name)
	fmt.Println(createStatement)
	result, err := db.Exec(createStatement)
	log.Println(result)

	return err
}

func seedDatabase(name string) error {
	log.Println("creating database tables")

	createTables := `
	DROP TABLE IF EXISTS "public"."%s";
	CREATE TABLE "public"."%s" (
		"id" int4 NOT NULL,
		"firstname" varchar COLLATE "default",
		"lastname" varchar COLLATE "default",
		"created" timestamp(6) NOT NULL,
		"updated" timestamp(6) NOT NULL,
		"suffix" varchar COLLATE "default",
		"prefix" varchar COLLATE "default",
		"phone" varchar COLLATE "default",
		"fax" varchar COLLATE "default",
		"email" varchar COLLATE "default",
		"website" varchar COLLATE "default",
		"photourl" varchar COLLATE "default",
		"organization" varchar COLLATE "default"
	)
	WITH (OIDS=FALSE);
	ALTER TABLE "public"."%s" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

`
	// DROP TABLE IF EXISTS "public"."surgeonaddresses";
	// CREATE TABLE "public"."surgeonaddresses" (
	// 	"surgeonid" int4 NOT NULL,
	// 	"addressid" uuid NOT NULL
	// )
	// WITH (OIDS=FALSE);
	// ALTER TABLE "public"."surgeonaddresses" ADD PRIMARY KEY ("surgeonid", "addressid") NOT DEFERRABLE INITIALLY IMMEDIATE;

	createTables = fmt.Sprintf(createTables, name, name, name)
	result, err := db.Exec(createTables)
	log.Println(result, err)
	if err != nil {
		log.Println(err)
	}

	log.Println("Inserting surgeon records")
	insertSurgeons := `
	BEGIN;
	INSERT INTO "public"."%s" VALUES ('1', 'Dilbert', 'Adam', '2016-10-13 13:40:39.904475', '2016-10-13 13:40:39.904475', 'MD', 'Dr.', '239-555-5555', '239-555-5555', 'dilbert@adams.com', 'http://www.dilbert.com', null, 'Dilbert');
	COMMIT;
	`

	insertSurgeons = fmt.Sprintf(insertSurgeons, name)
	result, err = db.Exec(insertSurgeons)
	log.Println(result, err)
	if err != nil {
		log.Println(err)
	}

	// TODO addresses

	// insertSurgeonAddresses := `
	// BEGIN;
	// INSERT INTO "public"."surgeonaddresses" VALUES ('1', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11');
	// INSERT INTO "public"."surgeonaddresses" VALUES ('1', '5634686a-3137-4cb4-9286-2ce5274b7f8a');
	// COMMIT;
	// `

	// result, err = db.Exec(insertSurgeonAddresses)
	// if err != nil {
	// 	log.Println(err)
	// }

	return err
}
