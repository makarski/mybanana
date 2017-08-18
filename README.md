My Banana
=========

My Banana is an example of implementation of an HTTP based API in Golang.

Setup
------

```
// Run the app (build if started for the first time)
make run

// View logs
make logs

// Stop application docker containers
make stop

// Build Containers without running
make build
```

Dependencies
------------

#### Depends ON
- mysql (maria db)
- github.com/...

#### Known Clients
- github.com/...


Project structure
------------------

```
├── config
│   └── config.go                       // config load and deserialization logic
├── db
│   ├── banana                          // db manager
│   │   ├── mocks
│   │   │   └── BananaFinder.go
│   │   └── banana.go
│   └── db.go                           // db connection and ORM mapping
├── handler
│   ├── banana
│   │   ├── banana.go
│   │   ├── get_banana_handler.go
│   │   └── get_banana_handler_test.go
│   ├── mocks
│   │   └── URLParamReader.go
│   └── handler.go                      // business logic
├── log
│   └── log.go
├── migrations                          // db migrations
│   └── 00001_banana_fixtures.sql
├── Dockerfile                          // Docker setup
├── Dockerfile-MariaDB
├── Makefile                            // Project setup
├── README.md
├── db_schema_setup.sql                 // db schema setup
├── docker-compose.yml
├── glide.lock
├── glide.yaml
└── main.go                             // app entry point
```

Consideration Points
--------------------
- Readme structure (setup, run, migrate, access docs)
- Dependency Injection (Interfaces)
- Middleware Usage
- Package Isolation
- RESTful Endpoints, Status Codes
- Error Handling (format)
- CLI helpers, migrations
