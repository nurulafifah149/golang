package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Port     uint   `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`

	MaxIdleConnection int `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int `mapstructure:"maxOpenConnection"`
	MaxIdleTime       int `mapstructure:"maxIdleTime"`
}

func NewPostgresConn() (db *sql.DB) {
	db, err := sql.Open("postgres", postgresDSN())
	if err != nil {
		panic(err)
	}

	postgresPoolConf(db)

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return
}

func NewPostgresGormConn() (db *gorm.DB) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(postgresDSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	postgresPoolConf(dbSQL)

	if err := dbSQL.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connect to Postgres")
	return db
}

func postgresDSN() string {
	return fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable`, Load.DataSource.Postgres.Master.Host, Load.DataSource.Postgres.Master.Port, Load.DataSource.Postgres.Master.Username, Load.DataSource.Postgres.Master.Password, Load.DataSource.Postgres.Master.DBName)
}

func postgresPoolConf(dbSQL *sql.DB) {
	// set extended config
	dbSQL.SetMaxIdleConns(Load.DataSource.Postgres.Master.MaxIdleConnection)
	dbSQL.SetMaxOpenConns(Load.DataSource.Postgres.Master.MaxOpenConnection)
	dbSQL.SetConnMaxIdleTime(time.Duration(Load.DataSource.Postgres.Master.MaxIdleTime))
}
