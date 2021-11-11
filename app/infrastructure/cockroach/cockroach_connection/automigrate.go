package cockroach_connection

import (
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

type Migrate struct {
	connection CockroachConnection
}

func NewMigrate(connection CockroachConnection) *Migrate {
	return &Migrate{connection: connection}
}

func (m *Migrate) AutoMigrateAll(tables ...interface{}) {
	db, err := m.connection.GetConnection()
	if err != nil {
		log.WithError(err).Fatal("Error getting connection to database")
	}
	db = db.AutoMigrate(tables...)
	if db.Error != nil {
		log.WithError(db.Error).Fatal("Error migrating entities")
	}
}
