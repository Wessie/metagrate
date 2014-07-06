package db

import (
	"os"
	"path/filepath"

	"github.com/Wessie/appdirs"
	"github.com/Wessie/congo"
	"github.com/boltdb/bolt"
)

var (
	// FileMode to create database file with
	databaseFileMode = os.FileMode(0750)
	// DB is an initialized database instance
	DB *bolt.DB
	// DatabaseOptions is the options given to bolt.Open
	DatabaseOptions *bolt.Options = nil
	DatabaseConfig  DBConfig
)

func init() {
	congo.AddSub("database", &DatabaseConfig)
}

type DBConfig struct {
	Directory string `json:"dir"`
	Filename  string `json:"name"`
}

func (dbc *DBConfig) Default() {
	dbc.Directory = appdirs.UserDataDir("seasonal", "seasonals", "", false)
	dbc.Filename = "data.db"
}

func Init() error {
	if err := os.MkdirAll(DatabaseConfig.Directory, databaseFileMode); err != nil {
		return err
	}

	path := filepath.Join(DatabaseConfig.Directory, DatabaseConfig.Filename)
	db, err := bolt.Open(path, databaseFileMode, DatabaseOptions)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
