package main

import (
	"log"
	"net/http"

	"git.arthrex.io/dschultz/surgeon-a/storage"
	"git.arthrex.io/dschultz/surgeon-a/surgeon"

	_ "github.com/lib/pq" // required for init of postgres database
)

func main() {
	// Select which repo to use
	// repo := storage.NewSurgeonDBRepository() // uncomment to use database as repo
	repo := storage.NewSurgeonInMemoryRepository() // uncomment to use temporary in memory object as repo
	defer repo.Close()
	// Create the service, passing in the repo
	svc := surgeon.NewSurgeonService(repo)

	// Create the handler for the surgeon (note: transport currently decided in endpoint)
	http.Handle("/surgeon/", surgeon.MakeGetSurgeonHandler(svc))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
