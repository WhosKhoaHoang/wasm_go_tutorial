package main

import (
    "flag"
    "log"
    "net/http"
)

var (
    listen = flag.String("listen", ":8080", "listen address")
    dir    = flag.String("dir", ".", "directory to serve")
)

// You'll get a "main redeclared in this block" warning, but don't sweat it.
func main() {
    flag.Parse()
    log.Printf("listening on %q...", *listen)
    log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
