package models

import "errors"

var (
	ErrNotFound = errors.New("models: resource not found")
)

type DataObject map[string]string

type Contact struct {
	Phone string
	Name  string
}
