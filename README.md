# Surgeons Service

### Install Golang's `dep` tool
```
$ brew install dep
```

### Ensure dependencies are up to date
>> NOTE: Use the `-vendor-only` flag to keep Dockerfile dependencies
that are not imported in the Go code.

From the `api/` directory:
```
$ dep ensure -vendor-only
```

### Run in container using convox
From the root directory:
```
$ convox start
```

### Call to service running in container
```
curl "https://surgeons.arthrex.xyz/surgeon/1" \
     -H 'Authorization: Bearer MY_JWT_GOES_HERE'
```

### Run locally
```
$ env GOPORT=8080 go run api/main.go
```

### Call to locally running surgeons service
```
$ curl "http://localhost:8080/surgeon/2"
```

### Add a new dependency to the project
From the `api/` directory:
```
$ dep ensure -add github.com/foo/bar
```

### Switch between in-memory persistence & database
Modify the main() func in api/main.go.  Create & pass in either in-memory repository to the service:
```
repo := storage.NewSurgeonInMemoryRepository()
```

or pass in the database repository to the service
```
repo := storage.NewSurgeonDBRepository()
```
