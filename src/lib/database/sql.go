package database

import (
	sqlx "database/sql"
	"fmt"

	config "github.com/Beelzebub0/go-crud-boilerplate/src/conf"
	_ "github.com/go-sql-driver/mysql"
)

type SQL interface {
	Connect() (*sqlx.DB, error)
}

type sql struct {
	conf config.SQLConfig
}

func InitSQL(conf config.SQLConfig) SQL {
	var d *sql
	d = &sql{
		conf: conf,
	}
	d.ping()
	return d
}

func (d *sql) ping() {
	dbDriver := d.conf.Driver
	dbHost := d.conf.Host
	dbPort := d.conf.Port
	dbUser := d.conf.User
	dbPass := d.conf.Password
	dbName := d.conf.Database
	dbProtocol := d.conf.Protocol
	connString := dbUser + ":" + dbPass + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := sqlx.Open(dbDriver, connString)
	if err != nil {
		fmt.Print(err.Error() + "\n")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error() + "\n")
		panic(err)
	}

	db.Close()
}

func (d *sql) Connect() (*sqlx.DB, error) {
	dbDriver := d.conf.Driver
	dbHost := d.conf.Host
	dbPort := d.conf.Port
	dbUser := d.conf.User
	dbPass := d.conf.Password
	dbName := d.conf.Database
	dbProtocol := d.conf.Protocol
	connString := dbUser + ":" + dbPass + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := sqlx.Open(dbDriver, connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
