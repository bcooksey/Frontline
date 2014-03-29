package main

import (
    "bytes"
    "flag"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

var (
    tmplMain        = "main.html"
    templateNames   = [...]string{tmplMain, "scripts.html", "footer.html"}
    templatePaths   []string
    templates       *template.Template
    reloadTemplates = true
    logger          *ServerLogger
    inProduction    = flag.Bool("production", false, "are we running in production")
)

func main() {
    if *inProduction {
        reloadTemplates = false
    }

    useStdout := !*inProduction

    logger = NewServerLogger(256, 256, useStdout)
    logger.Noticef("Starting Frontline Server...\n")

    r := mux.NewRouter()
    r.HandleFunc("/", handleIndex)
	http.HandleFunc("/static/js/", handleStaticJs)
	http.HandleFunc("/static/img/", handleStaticImg)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8082", nil))
}

func GetTemplates() *template.Template {
    if reloadTemplates || (nil == templates) {
        if 0 == len(templatePaths) {
            for _, name := range templateNames {
                templatePaths = append(templatePaths, filepath.Join("tmpl", name))
            }
        }
        templates = template.Must(template.ParseFiles(templatePaths...))
    }
    return templates
}

func ExecTemplate(w http.ResponseWriter, templateName string, model interface{}) bool {
    var buf bytes.Buffer
    if err := GetTemplates().ExecuteTemplate(&buf, templateName, model); err != nil {
        logger.Errorf("Failed to execute template '%s', error: %s", templateName, err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return false
    } else {
        // at this point we ignore error
        w.Write(buf.Bytes())
    }
    return true
}
