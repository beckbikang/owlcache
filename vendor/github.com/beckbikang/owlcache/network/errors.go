package network

import (
	"errors"
)

var (
	ErrorKeyNotFound       = errors.New("Not found key in cache")
	ErrorKeyNotExist       = errors.New("Key does not exist")
	ErrorOpenHttpsSelected = errors.New("Configuration file Open_Https filled in error")
	ErrorCloseTcp          = errors.New("Configuration file CloseTcp filled in error")
	ErrorGroupWorkMode     = errors.New("Configuration file GroupWorkMode filled in error")
	ErrorCors              = errors.New("Configuration file Cors filled in error")
)
