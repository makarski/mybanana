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

Project structure
------------------

```
├── config
│   └── config.go                       // config load and deserialization logic
├── db
│   ├── banana                          // db manager
│   │   └── banana.go
│   └── db.go                           // db connection and ORM mapping
├── handler
│   ├── banana
│   │   ├── banana.go
│   │   └── get_banana_handler.go
│   └── handler.go                      // business logic (common for all handlers)
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
└── main.go                             // app entry point
```
