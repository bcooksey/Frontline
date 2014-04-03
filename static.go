package main

import (
    "net/http"
    "path/filepath"
)

func serveFileFromDir(w http.ResponseWriter, r *http.Request, dir, fileName string) {
    filePath := filepath.Join(dir, fileName)
    // TODO: Do I want? May be a better way to handle
    // if !PathExists(filePath) {
    // 	logger.Noticef("serveFileFromDir() file '%s' doesn't exist, referer: '%s'", fileName, getReferer(r))
    // }
    http.ServeFile(w, r, filePath)
}

func handleStaticJs(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Path[len("/static/js/"):]
    serveFileFromDir(w, r, "static/js", file)
}

func handleStaticImg(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Path[len("/static/img/"):]
    serveFileFromDir(w, r, "static/img", file)
}
