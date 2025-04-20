package config

import (
	"os"
	"strconv"
)

var User = os.Getenv("USER")
var Host = os.Getenv("HOST")
var Password = os.Getenv("PASSWD")
var DatabaseName = os.Getenv("DATABASE")
var Port, err = strconv.Atoi(os.Getenv("PORT"))
