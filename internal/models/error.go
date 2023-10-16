package models

import "errors"

var (
	ErrorOrderNotFound = errors.New("order with this id not found")
)
