# Movie Aggregated Rating Microservice

---

this backend contians three microservices :

- Rating service : This contains aggreagated results of user input on diffrent films and tv shows. It uses Apache kafka to communicate with the rating ingestor
- Metadata service : this contains the metadeta of the movies
- Movie servcie : this is the front facing grpc api that acts as the API gateway for all user request .

---

App contains both in memory database and service directory .

## Service structure

```
./movieapp
├── LICENSE
├── README.md
├── api
│   └── movie.proto
├── cmd
│   ├── ratingingestor
│   │   ├── main.go
│   │   └── ratingsdata.json
│   └── sizecompare
│       ├── main.go
│       └── main_test.go
├── gen
│   ├── mock
│   │   ├── metadata
│   │   │   └── repository
│   │   │       └── repository.go
│   │   ├── movie
│   │   │   └── repository
│   │   │       └── repository.go
│   │   └── rating
│   │       └── repository
│   │           └── repository.go
│   ├── movie.pb.go
│   └── movie_grpc.pb.go
├── go.mod
├── go.sum
├── internal
│   └── grpcutil
│       └── grpcutil.go
├── logs
│   └── zap-logger
│       └── log.go
├── metadata
│   ├── Dockerfile
│   ├── cmd
│   │   ├── config.go
│   │   └── main.go
│   ├── configs
│   │   └── base.yaml
│   ├── internal
│   │   ├── controller
│   │   │   └── metadata
│   │   │       ├── controller.go
│   │   │       └── controller_test.go
│   │   ├── handler
│   │   │   ├── grpc
│   │   │   │   └── grpc.go
│   │   │   └── metahttp
│   │   │       └── http.go
│   │   └── repository
│   │       ├── error.go
│   │       ├── memory
│   │       │   └── memory.go
│   │       └── mysql
│   │           └── mysql.go
│   ├── kubernetes-deployment.yml
│   ├── main
│   └── pkg
│       ├── model
│       │   ├── mapper.go
│       │   └── metadata.go
│       └── testutil
│           └── testutil.go
├── movie
│   ├── Dockerfile
│   ├── cmd
│   │   ├── config.go
│   │   └── main.go
│   ├── configs
│   │   └── base.yaml
│   ├── internal
│   │   ├── controller
│   │   │   └── movie
│   │   │       ├── controller.go
│   │   │       └── controller_test.go
│   │   ├── gateway
│   │   │   ├── errors.go
│   │   │   ├── metadata
│   │   │   │   ├── grpc
│   │   │   │   │   └── grpc.go
│   │   │   │   └── http
│   │   │   │       └── http.go
│   │   │   └── rating
│   │   │       ├── grpc
│   │   │       │   └── grpc.go
│   │   │       └── http
│   │   │           └── http.go
│   │   └── handler
│   │       ├── grpc
│   │       │   └── grpc.go
│   │       └── http
│   │           └── http.go
│   ├── kubernetes-deployment.yml
│   ├── main
│   └── pkg
│       ├── model
│       │   └── model.go
│       └── testutil
│           └── testutil.go
├── pkg
│   ├── discovery
│   │   ├── consul
│   │   │   └── consul.go
│   │   ├── discovery.go
│   │   └── memory
│   │       └── memory.go
│   └── tracing
│       └── tracing.go
├── rating
│   ├── Dockerfile
│   ├── cmd
│   │   ├── config.go
│   │   └── main.go
│   ├── configs
│   │   └── base.yaml
│   ├── internal
│   │   ├── controller
│   │   │   └── rating
│   │   │       ├── controller.go
│   │   │       └── controller_test.go
│   │   ├── handler
│   │   │   ├── grpc
│   │   │   │   └── grpc.go
│   │   │   └── http
│   │   │       └── http.go
│   │   ├── ingester
│   │   │   └── kafka
│   │   │       └── ingester.go
│   │   └── repository
│   │       ├── errors.go
│   │       ├── memory
│   │       │   └── memory.go
│   │       └── mysql
│   │           └── mysql.go
│   ├── kubernetes-deployment.yml
│   ├── main
│   └── pkg
│       ├── model
│       │   └── rating.go
│       └── testutil
│           └── testutil.go
├── schema
│   └── schema.sql
└── test
    └── integration
        └── main.go
```

Run app
