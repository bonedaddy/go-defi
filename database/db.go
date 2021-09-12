// Copyright (c) 2021 Alexandre Trottier (bonedaddy). All rights reserved.
// Use of this source code is governed by the Apache 2.0 License that can be found in
// the LICENSE file.

package database

import (
	"fmt"
	"strings"

	"github.com/bonedaddy/go-defi/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// Opts is used to configure database connections
type Opts struct {
	Type           string
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	DBPath         string
	SSLModeDisable bool
}

// OptsFromConfig returns database options from a config file
func OptsFromConfig(cfg config.Database) *Opts {
	return &Opts{
		Type:           cfg.Type,
		Host:           cfg.Host,
		Port:           cfg.Port,
		User:           cfg.User,
		Password:       cfg.Pass,
		DBName:         cfg.DBName,
		DBPath:         cfg.DBPath,
		SSLModeDisable: cfg.SSLModeDisable,
	}
}

// DSN returns the config string used with gorm
func (db *Opts) DSN() string {
	var sslModeDisable string
	if db.SSLModeDisable {
		sslModeDisable = "sslmode=disable"
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		db.Host, db.User, db.Password, db.DBName, db.Port, sslModeDisable,
	)
}

// Open returns the gorm dialector for the corresponding db type
func (db *Opts) Open() (gorm.Dialector, error) {
	switch strings.ToLower(db.Type) {
	case "postgres":
		return postgres.Open(db.DSN()), nil
	case "sqlite":
		return sqlite.Open(db.DBPath + db.DBName + ".db"), nil
	default:
		return nil, errors.New("unsupported db type")
	}

}

// New returns a new database client
func New(opts *Opts) (*Database, error) {
	dialector, err := opts.Open()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialector, nil)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

// Migrator returns the underlying database migrator, useful for managing the database
func (d *Database) Migrator() gorm.Migrator {
	return d.db.Migrator()
}

// AutoMigrate is used to automatically migrate datbase tables
func (d *Database) AutoMigrate() error {
	var tables []interface{}
	// make sure to uncomment and change the objects below accordingly
	// the current objects are an example of what it would look like
	// tables = append(tables, &Address{}, &Token{}, &Backlog{})
	for _, table := range tables {
		if err := d.db.AutoMigrate(table); err != nil {
			return errors.Wrap(err, "auto migrate")
		}
	}
	return nil
}

// DB returns the underlying gorm database
func (d *Database) DB() *gorm.DB {
	return d.db
}

// Close shuts down the database
func (d *Database) Close() error {
	sdb, err := d.db.DB()
	if err != nil {
		return errors.Wrap(err, "db")
	}
	return sdb.Close()
}
