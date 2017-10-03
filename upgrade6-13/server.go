package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "io/ioutil"
)

func handleUpgrade(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
        w.WriteHeader(400)
        return
    }
    fmt.Println("Upgrade to MyProtocol")

    hijacker := w.(http.Hijacker)
    conn, readWriter, err := hijacker.Hijack()
    if err != nil {
        panic(err)
        return
    }
    defer conn.Close()
}
