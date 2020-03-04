package ves

import "fmt"

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Error %s %s: %s", e.Code, e.Title, e.Detail)
}
