package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Database struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		DBName    string `json:"dbname"`
		Charset   string `json:"charset"`
		ParseTime bool   `json:"parseTime"`
		Loc       string `json:"loc"`
	} `json:"database"`
}

var db *sql.DB

func Connect() {
	//Open and read the config file

	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		panic(err)
	}

	// Construct the DSN from the config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.Charset,
		config.Database.ParseTime,
		config.Database.Loc,
	)

	d, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *sql.DB {
	return db
}
