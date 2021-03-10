### Setup

```
go mod init github.com/gireeshcse/graphlq-golang
go get -v github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

### Important files

* **graph/schema.graphqls** — This is the file where we will add GraphQL schemas.

* **graph/schema.resolvers.go** — This is where your application code lives. *generated.go* will call into this to get the data the user has requested.

* **server.go** — This is a minimal entry point that sets up an *http.Handler* to the generated GraphQL server. start the server with **go run server.go** and open our browser and we should see the graphql playground

* **gqlgen.yml** — The gqlgen config file, knobs for controlling the generated code.

* **graph/generated/generated.go** — The GraphQL execution runtime, the bulk of the generated code.

* **graph/model/models_gen.go** — Generated models required to build the graph. Often we will override these with our own models. Still very useful for input types.

### Notes

* After updation of schemas we need to run the below command 

```
go run github.com/99designs/gqlgen generate
```

* If you remove schemas then we have to related functions from schema.resolvers.go as well.Otherwise validation errors will occur



### References

[graphql-go](https://www.howtographql.com/graphql-go/)