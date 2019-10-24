/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package helpers

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

func ExtractValidationErrors(err error) map[string]string {

	validationErrors := make(map[string]string)
	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		validationErrors["object"] = fmt.Sprintf("%s", err)
		return validationErrors
	}

	for _, err := range err.(validator.ValidationErrors) {
		validationErrors[err.Field()] = fmt.Sprintf("Failed on %s", err.ActualTag())

	}
	return validationErrors
}
