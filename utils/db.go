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
	url = fmt.Sprintf("%s:%s@/%s:%s?parseTime=true", data.User, data.Password, data.Table, data.Port)
	return url
}
