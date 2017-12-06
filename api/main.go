package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"git.arthrex.io/dschultz/surgeon-a/api/storage"
	"git.arthrex.io/dschultz/surgeon-a/api/surgeon"

	_ "github.com/lib/pq" // required for init of postgres database
)

func main() {
	// Select which repo to use
	repo := storage.NewSurgeonDBRepository() // uncomment to use database as repo
	//repo := storage.NewSurgeonInMemoryRepository() // uncomment to use temporary in memory object as repo
	defer repo.Close()

	// Create the service, passing in the repo
	svc := surgeon.NewSurgeonService(repo)

	// Create the handler for the surgeon (note: transport currently decided in endpoint)
	handler := surgeon.NewHandler(svc)
	log.Println("port:", os.Getenv("GOPORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GOPORT")), handler))
}
