package ves

import (
	"fmt"
	"strconv"
)

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Status Status `json:"status,omitempty"`
	Code   Status `json:"code,omitempty"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}

type Status string

func (s Status) Int() (int, error) {
	return strconv.Atoi(string(s))
}

func (e Error) Error() string {
	return fmt.Sprintf("Error %s (%s) %s: %s", e.Status, e.Code, e.Title, e.Detail)
}
