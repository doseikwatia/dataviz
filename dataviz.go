package main

import (
	"flag"
	"net/http"
)

const (
	apiprefix    = "/api/"
	viewprefix   = "/"
	assetsprefix = "/assets/"
)

var (
	addr = flag.String("addr", ":8080", "")
)

//go:generate go-bindata-assetfs -pkg main -o resource.go assets/js assets/css assets/images views/
func main() {
	flag.Parse()
	http.HandleFunc(apiprefix, apihandlefunc)
	http.Handle(assetsprefix,http.StripPrefix(assetsprefix,http.FileServer(assetFS())))
	http.HandleFunc(viewprefix, viewhandlefunc)
	http.ListenAndServe(*addr, nil)
}
