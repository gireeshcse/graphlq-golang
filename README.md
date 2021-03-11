### Project Setup

* Clone the project
* copy and update config as per your system configuration

```
cp config_sample.yml  config.yml
```

### Initial Project Setup

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

#### Dummy Query

```
query {
	dummyLinks{
    title
    address,
    user{
      name
    }
  }
}
```

Output

```
{
  "data": {
    "dummyLinks": [
      {
        "title": "our dummy link",
        "address": "https://address.org",
        "user": {
          "name": "admin"
        }
      }
    ]
  }
}
```

### Database Setup

* Run mysql server and create database hackernews;

* Create following folder structure

```
internal/pkg/db/migrations/mysql
```

* Install go mysql driver and golang-migrate packages

```
go get -u github.com/go-sql-driver/mysql
go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
cd internal/pkg/db/migrations/
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table
```

* Update migrations 

File : internal/pkg/db/migrations/mysql/000001_create_users_table.up.sql

```
CREATE TABLE IF NOT EXISTS Users(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
)
```

File: internal/pkg/db/migrations/mysql/000002_create_links_table.up.sql
```
CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)
```

* Run the migration

```
migrate -database "mysql://root:dbpass@(localhost:3316)/hackernews"  -path internal/pkg/db/migrations/mysql up
```

* Create and update **internal/pkg/db/mysql/mysql.go** which connects the database server



### References

[graphql-go](https://www.howtographql.com/graphql-go/)