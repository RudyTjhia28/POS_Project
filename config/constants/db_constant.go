package constant

import (
	"fmt"
)

const (
	DbUser     = "developmentuser1"
	DbPassword = "password1"
	DbName     = "pos-online"
	DbPort     = 5432
	DbTimezone = "Asia/Jakarta"

	ServicePort = 8080
)

func GetDbConfig(constantName string) {
	switch constantName {
	case "DbUser":
		fmt.Println(DbUser)
	case "DbPassword":
		fmt.Println(DbPassword)
	case "DbName":
		fmt.Println(DbName)
	case "DbPort":
		fmt.Println(DbPort)
	case "DbTimezone":
		fmt.Println(DbTimezone)
	case "ServicePort":
		fmt.Println(ServicePort)
	default:
		fmt.Println("Unknown constant")
	}
}

func GetServiceConfig(constantName string) {
	switch constantName {
	case "ServicePort":
		fmt.Println(ServicePort)
	default:
		fmt.Println("Unknown constant")
	}
}
