package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/feckmore/surgeon-a/api/storage"
	"github.com/feckmore/surgeon-a/api/surgeon"

	_ "github.com/lib/pq" // required for init of postgres database
)

func main() {
	devEnv, _ := strconv.ParseBool(os.Getenv("DEVELOPMENT")) // try pull all env vars in main
	// // Select which repo to use
	// //repo := storage.NewSurgeonInMemoryRepository() // uncomment to use temporary in memory object as repo
	repo, err := storage.NewSurgeonDBRepository(os.Getenv("DATABASE_SCHEME"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("CREATE_DATABASE_NAME"), devEnv)
	defer repo.Close()
	if err != nil {
		//TODO
	}

	// Create the service, passing in the repo
	svc := surgeon.NewSurgeonService(repo)

	// Create the handler for the surgeon (note: transport currently decided in endpoint)
	handler := surgeon.NewHandler(svc)
	log.Println("port:", os.Getenv("GOPORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GOPORT")), handler))
}
