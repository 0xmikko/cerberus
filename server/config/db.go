/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */
package config

import (
	"errors"
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	DBname   string
	Password string
}

func getDBConfigFromJSON(fname string) (*DBConfig, error) {
	devConfig := &DBConfig{}
	err := LoadConfigFromJSON(fname, &devConfig)
	return devConfig, err
}

// Get Developers config from JSON files stored in config folder
func GetDBConfig(configT int) (*DBConfig, error) {
	switch configT {
	case DEV:
		return getDBConfigFromJSON("./config/dev_db.json")
	case TEST:
		return getDBConfigFromJSON("./config/test_db.json")
	case PROD:
		return getProdConfig()

	}
	return nil, errors.New("Wrong databse configuration type providedd")
}

func getProdConfig() (*DBConfig, error) {

	devConfig := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		DBname:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD")}

	return devConfig, nil
}

func (c DBConfig) String() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", c.User, c.Password, c.Host, c.DBname)
}
