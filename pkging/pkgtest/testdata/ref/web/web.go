package web

import (
	"net/http"

	"github.com/al45tair/pkger"
)

func Serve() {
	pkger.Stat("github.com/al45tair/pkger:/README.md")
	dir := http.FileServer(pkger.Dir("/public"))
	http.ListenAndServe(":3000", dir)
}
