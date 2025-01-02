package database

import "fmt"

var (
	UserDbConnStr = fmt.Sprintf("%s://%s:%s@%s:%d/%s", database, user, password, host, port, userDb)
)

// connStr := "postgres://myuser:mypassword@localhost:5432/mydb"

const (
	port     = 5432
	database = "postgres"
	user     = "postgres"
	password = "Gautam@123"
	host     = "localhost"
	userDb   = "postgres"
)

const (
	SignUpTable = "users"
)

var DbManager = &DBManager{}

func InitApplicationLayer() {
	DbManager = InitPgDBConnection()
}
