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
.
├── pkg
│   ├── config
│   │   └── config.go
│   ├── db
│   │   ├── banana
│   │   │   ├── mocks
│   │   │   └── banana.go
│   │   └── db.go
│   ├── handler
│   │   ├── banana
│   │   │   ├── banana.go
│   │   │   ├── get_banana_handler.go
│   │   │   └── get_banana_handler_test.go
│   │   ├── mocks
│   │   │   └── url_param_reader.go
│   │   └── handler.go
│   └── log
│       └── log.go
├── resources
│   └── migrations
│       └── 00001_banana_fixtures.sql
├── Dockerfile
├── Dockerfile-MariaDB
├── Makefile
├── README.md
├── db_schema_setup.sql
├── docker-compose.yml
├── glide.lock
├── glide.yaml
└── main.go
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
