package main

import (
	"net/http"
	// "strings"
)

const (
	msgTmpl = ``
)

func init() {

}

func gitHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, ``, 405)
}

func processHook() {

}
