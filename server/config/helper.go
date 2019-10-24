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
	"encoding/json"
	"io/ioutil"
)

// Gets config from file
func LoadConfigFromJSON(filefname string, targetStruct interface{}) error {

	rawData, err := ioutil.ReadFile(filefname)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawData, targetStruct)
	return err
}
