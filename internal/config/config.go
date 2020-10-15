package config

import (
	"errors"
	"fmt"
	"log"
	"template/internal/store"
)

// Config is a struct for our Configuration Files
type Config struct {

	//BaseURL fupisha's fully qualified domain name.
	BaseURL string `envconfig:"FRIKA_BASE_URL"`
	//Title name of the application e.g. Frika.
	Title string `envconfig:"FRIKA_TITLE"`
	//TextLogging write api requests to file.
	TextLogging bool `envconfig:"FRIKA_TEXT_LOGGING"`
	//LogLevel category of the log.
	LogLevel string `envconfig:"FRIKA_LOG_LEVEL"`

	// This struct is for handling Json web tokens secrets
	JWT struct {
		//Secret secret jwt signing key.
		Secret string `envconfig:"FRIKA_JWT_SECRET"`
		//ExpireDelta duration after which the token is rendered invalid.
		ExpireDelta string `envconfig:"FRIKA_JWT_EXPIRE_DELTA"`
	}

	// Working with SMTP servers for email
	SMTP struct {
		//Port smtp port
		Port int `envconfig:"FRIKA_SMTP_PORT"`
		//Host smtp host e.g. smtp.gmail.com
		Host string `envconfig:"FRIKA_SMTP_HOST"`
		//Username smtp username.
		Username string `envconfig:"FRIKA_SMTP_USERNAME"`
		//Password smtp password.
		Password string `envconfig:"FRIKA_SMTP_PASSWORD"`
		//FromName email sender's name.
		FromName string `envconfig:"FRIKA_SMTP_FROM_NAME"`
		//FromAddress email sender's email address
		FromAddress string `envconfig:"FRIKA_SMTP_FROM_ADDRESS"`
	}

	// Store for database Drivers
	Store struct {
		//Type the type of database. e.g. mongo
		Type string `envconfig:"FRIKA_STORE_TYPE"`

		//PostgreSQL Parameters For Db  connection
		PostgreSQL struct {
			//Address postgresql host and port. e.g. localhost:5432
			Address string `envconfig:"FRIKA_STORE_POSTGRESQL_ADDRESS"`
			//Username postgresql user.
			Username string `envconfig:"FRIKA_STORE_POSTGRESQL_USERNAME"`
			//Password postgresql password associated with the user.
			Password string `envconfig:"FRIKA_STORE_POSTGRESQL_PASSWORD"`
			//Database postgresql database name.
			Database string `envconfig:"FRIKA_STORE_POSTGRESQL_DATABASE"`
			//SSLMode if enabled postgresql will encrypt the communication to and from.
			SSLMode string `envconfig:"FRIKA_STORE_POSTGRESQL_SSLMODE"`
			//SSLRootCert requires if SSLMode is enabled.
			SSLRootCert string `envconfig:"FRIKA_STORE_POSTGRESQL_SSLROOTCERT"`
		}

		// Mongo Parameters For Db  connection
		Mongo struct {
			//Address mongo host and port. e.g localhost:27017
			Address string `envconfig:"FRIKA_STORE_MONGO_ADDRESS"`
			//Username mongo user.
			Username string `envconfig:"FRIKA_STORE_MONGO_USERNAME"`
			//Password mongo user's password.
			Password string `envconfig:"FRIKA_STORE_MONGO_PASSWORD"`
			//Database mongo database name.
			Database string `envconfig:"FRIKA_STORE_MONGO_DATABASE"`
		}

		// MySQL Parameters For Db  connection
		MySQL struct {
			//Address mysql host and port. e.g. localhost:3306
			Address string `envconfig:"FRIKA_STORE_MYSQL_ADDRESS"`
			//Username mysql user.
			Username string `envconfig:"FRIKA_STORE_MYSQL_USERNAME"`
			//Password mysql user's passsword.
			Password string `envconfig:"FRIKA_STORE_MYSQL_PASSWORD"`
			//Database mysql database name.
			Database string `envconfig:"FRIKA_STORE_MYSQL_DATABASE"`
		}
	}
}

func (cfg *Config) GetStore() (store.Store, error) {

	switch cfg.Store.Type {
	case "mongo":
		var (
			address  = cfg.Store.Mongo.Address
			username = cfg.Store.Mongo.Username
			password = cfg.Store.Mongo.Password
			database = cfg.Store.Mongo.Database
		)

		if (address == "") || (username == "") || (database == "") || (password == "") {
			err := errors.New("missing mongodb configuration variable")
			log.Fatalf("config: %s", err.Error())
		}
		/*
			address := fmt.Sprintf("%s:%s", host, port)
			return mongodb.Connect(
				address,
				username,
				password,
				database,
			)
		*/
	case "postgres":
		var (
			address     = cfg.Store.PostgreSQL.Address
			username    = cfg.Store.PostgreSQL.Username
			password    = cfg.Store.PostgreSQL.Password
			database    = cfg.Store.PostgreSQL.Database
			sslmode     = cfg.Store.PostgreSQL.SSLMode
			sslrootcert = cfg.Store.PostgreSQL.SSLRootCert
		)

		if (address == "") || (username == "") || (database == "") || (password == "") || (sslmode == "") || (sslrootcert == "") {
			err := errors.New("missing postgres configuration variable")
			log.Fatalf("config: %s", err.Error())
		}
	}
	return nil, fmt.Errorf("config: unknown store type: %s", cfg.Store.Type)
}
