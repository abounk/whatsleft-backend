package util

import (
	"fmt"
	"whatsleft/backend/config"

	log "github.com/sirupsen/logrus"
)

func NewConnectionUrlBuilder(connection string, db config.Database) string {
	var url string
	switch connection {
	case "mysql":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			db.Username,
			db.Password,
			db.Hostname,
			db.Port,
			db.Database,
		)
	case "dns":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?tls=true&parseTime=true",
			db.Username,
			db.Password,
			db.Hostname,
			db.Database,
		)
	default:
		log.Errorf("error, unknow the stuff: %s", connection)
	}

	return url
}
