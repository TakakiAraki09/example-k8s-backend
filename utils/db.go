package utils

import (
	"fmt"
)

type DatabaseMetadata struct {
	User     string
	Password string
	Host     string
	Port     string
	Table    string
}

func CreateDBUrlString(data DatabaseMetadata) (url string) {
	url = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", data.User, data.Password, data.Host, data.Port, data.Table)
	return url
}

// kazuhira:password@(172.17.0.2:3306)/practice
