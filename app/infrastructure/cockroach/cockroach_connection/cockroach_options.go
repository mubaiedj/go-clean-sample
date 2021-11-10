package cockroach_connection

import (
	"fmt"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"os"
)

type CockroachOptions struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *CockroachOptions {
	return new(CockroachOptions)
}

func (c *CockroachOptions) DatabaseName(databaseName string) *CockroachOptions {
	c.databaseName = &databaseName
	return c
}

func (c *CockroachOptions) Host(host string) *CockroachOptions {
	c.host = &host
	return c
}

func (c *CockroachOptions) Port(port int) *CockroachOptions {
	c.port = &port
	return c
}

func (c *CockroachOptions) User(user string) *CockroachOptions {
	c.user = &user
	return c
}

func (c *CockroachOptions) Password(password string) *CockroachOptions {
	c.password = &password
	return c
}

func MergeOptions(opts ...*CockroachOptions) *CockroachOptions {
	option := new(CockroachOptions)

	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

var (
	cockroachDefaultPort = 26257
)

func (d *CockroachOptions) GetUrlConnection() string {
	UrlCockroachFormat := "postgresql://%v:%v@%v:%v/%v"

	if d.port == nil {
		d.port = &cockroachDefaultPort
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" || environment == "" {
		UrlCockroachFormat = "postgresql://%v:%v@%v:%v/%v?sslmode=disable"
	}

	log.Info("Connection: %s", fmt.Sprintf(UrlCockroachFormat, *d.user, "************", *d.host, *d.port, *d.databaseName))
	return fmt.Sprintf(UrlCockroachFormat, *d.user, *d.password, *d.host, *d.port, *d.databaseName)
}
