package database

import (
	"whatsleft/backend/config"
	"whatsleft/backend/pkg/util"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func NewMySQLDatabaseConnection(config *config.Config) (*sqlx.DB, error) {
	mysqlUrl := util.NewConnectionUrlBuilder("mysql", config.Database)
	db, err := sqlx.Connect("mysql", mysqlUrl)
	if err != nil {
		log.Errorf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("error, can't ping the database, %s", err.Error())
		return nil, err
	}

	log.Info("Connected to mysql database successfully")
	return db, nil
}
