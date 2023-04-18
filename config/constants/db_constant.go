package constant

import (
	"fmt"
	"strconv"
)

const (
	DbUser     = "developmentuser1"
	DbPassword = "password1"
	DbName     = "pos-online"
	DbPort     = "5432"
	DbTimezone = "Asia/Jakarta"

	ServicePort = ":8080"
	SectetKey   = "401fb76e-50ab-402d-a1cc-21261d2bfc4a"
)

var constantMap = map[string]string{
	"DbUser":      DbUser,
	"DbPassword":  DbPassword,
	"DbName":      DbName,
	"DbPort":      DbPort,
	"DbTimezone":  DbTimezone,
	"ServicePort": ServicePort,
	"SecretKey":   SectetKey,
}

func GetString(constantName string) (*string, error) {
	value, exists := constantMap[constantName]
	if !exists {
		return nil, fmt.Errorf("unknown constant: %s", constantName)
	}
	return &value, nil
}

func GetInt(constantName string) (*int, error) {
	value, exists := constantMap[constantName]
	if !exists {
		return nil, fmt.Errorf("unknown constant: %s", constantName)
	}
	intValue, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func GetAllConstants() map[string]string {
	return constantMap
}
