package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type connectionString struct {
	master string
	slave  string
	domain string
}

// why is this needed? because sonarqube will say there's a duplication if you write it multiple times
var stringType = "type"
var stringConnection = "connection"

// Connections return the wrapped actions of postgres database
type Connections interface {
	Open() *Database
}

// Database where the opened database connection is being used
type Database struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func Initialize(master string, slave string, domain string) Connections {
	return &connectionString{
		master: master,
		slave:  slave,
		domain: domain,
	}
}

// Open is creating the database postgres connections, both master and slave
func (cs *connectionString) Open() *Database {
	// log fields for logrus, no need to write this multiple times
	logFields := logrus.Fields{
		"platform": "mysql",
		"domain":   cs.domain,
	}
	logMasterFields := logrus.Fields{
		stringType:       "master",
		stringConnection: cs.master,
	}
	logSlaveFields := logrus.Fields{
		stringType:       "slave",
		stringConnection: cs.slave,
	}

	logrus.WithFields(logFields).Info("Connecting to mysql DB")

	logrus.WithFields(logFields).Info("Opening Connection to Master")
	dbMaster, err := sqlx.Open("mysql", cs.master)
	if err != nil {
		logrus.WithFields(logMasterFields).Fatal(err)
		panic(err)
	}
	err = dbMaster.Ping()
	if err != nil {
		logrus.WithFields(logMasterFields).Fatal(err)
		panic(err)
	}

	logrus.WithFields(logFields).Info("Opening Connection to Slave")
	dbSlave, err := sqlx.Open("postgres", cs.master)
	if err != nil {
		logrus.WithFields(logSlaveFields).Fatal(err)
		panic(err)
	}
	err = dbSlave.Ping()
	if err != nil {
		logrus.WithFields(logSlaveFields).Fatal(err)
		panic(err)
	}

	return &Database{
		Master: dbMaster,
		Slave:  dbSlave,
	}
}
