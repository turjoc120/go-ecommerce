package db

import (
	"ecoommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnString(cnf *config.DbConfig) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cnf.HOST, cnf.PORT, cnf.USER, cnf.PASSWORD, cnf.DBNAME)
	if !cnf.SSLMode {
		connString += " sslmode=disable"
	}
	return connString

}

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := getConnString(cnf)
	dbConn, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return dbConn, nil
	}
	return dbConn, nil
}
