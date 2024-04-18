package store

import (
	"database/sql"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type SqlServer struct{}

var dbConnectionString string

func NewSqlServer() IStore {
	LoadEnv()
	return &SqlServer{}
}

func (sqlServer *SqlServer) openDatabaseConnection() (*sql.DB, error) {
	return sql.Open("mssql", dbConnectionString)
}

func (sqlServer *SqlServer) closeDatabaseConnection(databaseConnection *sql.DB) error {
	return databaseConnection.Close()
}

func LoadEnv() {
	dbConnectionString = os.Getenv("DB_STRING")
	if dbConnectionString == "" {
		panic("No hay cadena de conexion a base de datos")
	}
}
