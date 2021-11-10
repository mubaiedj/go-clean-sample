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
		log.WithSource("MIGRATIONS").WithError(err).Fatal(err.Error())
	}
	db = db.AutoMigrate(tables...)
	if db.Error != nil {
		log.WithSource("MIGRATIONS").WithError(db.Error).Fatal(db.Error.Error())
	}
}
