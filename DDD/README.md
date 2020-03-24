# ドメイン設計

レイヤードアーキテクチャ

```
.
├── README.md
├── account
│   ├── Dockerfile
│   ├── application
│   │   └── usecase
│   │       └── xxx_usecase.go
│   ├── domain
│   │   ├── model
│   │   │   └── xxx.go
│   │   ├── repository
│   │   │   └── xxx_repository.go
│   │   └── service
│   │       └── xxx_service.go
│   ├── infrastracture
│   │   └── persistence
│   │       └── datastore
│   │           └── xxx_repository.go
│   ├── interfaces
│   ├── main.go
│   └── registry
├── docker-compose.yml
├── go.mod
├── graphql
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── graph.go
│   ├── main.go
│   ├── models.go
│   ├── mutation_resolver.go
│   ├── query_resolver.go
│   └── schema.graphql
├── k8s
├── mysql
│   ├── Dockerfile
│   ├── conf.d
│   │   └── my.cnf
│   └── init
│       ├── account.sql
│       └── insert.sql
└── proto
    └── account
        ├── account.pb.go
        └── account.proto
```
