package storage

import (
	"database/sql"
	"fmt"
	"log"

	"git.arthrex.io/dschultz/surgeon-a/api/surgeon"
)

const (
	host     = "localhost"
	port     = 5432
	username = "dschultz"
	password = ""
	dbname   = "surgeon"
)

var db *sql.DB

// surgeonDBRepo satisfies the Repositor interface
type surgeonDBRepo struct {
	db *sql.DB
}

// NewSurgeonDBRepository returns object that implements the Repositor interface
func NewSurgeonDBRepository() surgeon.Repositor {
	log.Println("NewSurgeonRepository")
	repo := new(surgeonDBRepo)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
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

	rows, err := ur.db.Query(`SELECT id, first_name, last_name, email FROM surgeons WHERE id = $1`, id)
	if err != nil {
		log.Fatal("Error in Query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		// err := rows.Scan(&id, &name)
		err := rows.Scan(&surgeon.ID, &surgeon.Firstname, &surgeon.Lastname, &surgeon.Email)
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
