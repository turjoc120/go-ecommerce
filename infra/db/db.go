package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/turjoc120/ecom/config"
)

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnString(cnf)
	dbConn, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConn, nil
}

func GetConnString(cnf *config.DbConfig) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cnf.HOST, cnf.PORT, cnf.USER, cnf.PASSWORD, cnf.DBNAME)
	if !cnf.SSLMode {
		connString += " sslmode=disable"
	}
	return connString
}
