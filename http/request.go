package http

import (
	"github.com/anaskhan96/soup"
)

func Get(url string) string {
	resp, err := soup.Get(url)
	if err != nil {
		panic(err.Error())
	}
	return resp
}
