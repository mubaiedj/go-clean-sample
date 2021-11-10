package cockroach

import (
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/cockroach/cockroach_connection"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/cockroach/db_model"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

func AutoMigrateEntities(connection cockroach_connection.CockroachConnection) {
	log.Info("AutoMigrateEntities...")
	migrate := cockroach_connection.NewMigrate(connection)
	migrate.AutoMigrateAll(
		db_model.MessageStatus{},
	)
	log.Info("AutoMigrateEntities... OK")
}

func CreateCockroachDbConnection() *cockroach_connection.CockroachDBConnection {
	cockroachHost := config.GetString("datasource.host")
	cockroachPort := config.GetInt("datasource.port")
	cockroachDatabase := config.GetString("datasource.database")
	cockroachUser := config.GetString("datasource.user")
	cockroachPassword := config.GetString("datasource.password")
	connection := cockroach_connection.NewCockroachConnection(cockroach_connection.Config().
		Host(cockroachHost).
		Port(cockroachPort).
		DatabaseName(cockroachDatabase).
		User(cockroachUser).
		Password(cockroachPassword),
	)
	return connection
}
