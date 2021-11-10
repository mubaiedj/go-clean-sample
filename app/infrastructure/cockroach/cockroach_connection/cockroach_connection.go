package cockroach_connection

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

const LogMode = true

var connection *gorm.DB

type CockroachConnection interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
}

type CockroachDBConnection struct {
	opts *CockroachOptions
	url  string
}

func NewCockroachConnection(opts ...*CockroachOptions) *CockroachDBConnection {
	databaseOptions := MergeOptions(opts...)
	url := databaseOptions.GetUrlConnection()
	if url == "" {
		log.Fatal(errors.New("Error creating connection, empty url").Error())
	}
	return &CockroachDBConnection{
		opts: databaseOptions,
		url:  url,
	}
}

func (r *CockroachDBConnection) GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil || !isAlive() {
		log.Info("Trying to connect to DB")
		connection, err = gorm.Open("postgres", r.url)
		if err != nil {
			log.WithError(err).Error("error trying to connect to DB")
			return nil, err
		} else {
			log.Info("Connected to DB")
		}
	}
	connection.LogMode(LogMode)
	return connection, nil
}

func (r *CockroachDBConnection) CloseConnection() {
	if err := connection.Close(); err != nil {
		log.WithError(err).Error("error trying to close connection")
	} else {
		log.Info("Connection Closed")
	}
}

func isAlive() bool {
	if err := connection.DB().Ping(); err != nil {
		log.WithError(err).Error("error trying to Ping to Db")
		return false
	}
	return true
}
