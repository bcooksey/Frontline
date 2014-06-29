package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
)

var (
    tmplMain        = "main.html"
    templatePaths   []string
    templates       *template.Template
    reloadTemplates = true
    logger          *ServerLogger
    inProduction    = flag.Bool("production", false, "are we running in production")
    configPath      = flag.String("config", "config.json", "Path to configuration file")
    sessionStore    *sessions.CookieStore

    config = struct {
        Dbuser           string
        Dbpass           string
        sessionSecretKey string
    }{
        "",
        "",
        "",
    }
)

func main() {
    if *inProduction {
        reloadTemplates = false
    }

    useStdout := !*inProduction

    logger = NewServerLogger(256, 256, useStdout)
    logger.Noticef("Starting Frontline Server...\n")

    r := mux.NewRouter()

    if err := readConfig(*configPath); err != nil {
        log.Fatalf("Failed reading config file %s. %s\n", *configPath, err.Error())
    }

    sessionStore = sessions.NewCookieStore([]byte(config.sessionSecretKey))

    // Root View
    r.HandleFunc("/", handleIndex)
    r.HandleFunc("/login", handleLogin)
    http.HandleFunc("/static/compiled/js/", handleStaticJs)
    http.HandleFunc("/static/img/", handleStaticImg)

    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8082", nil))
}

func GetTemplates() *template.Template {
    if reloadTemplates || (nil == templates) {
        templates = template.Must(template.ParseGlob("tmpl/*"))
    }
    return templates
}

func ExecTemplate(w http.ResponseWriter, templateName string) bool {
    data := struct{}{}
    return ExecTemplateWithContext(w, templateName, data)
}

func ExecTemplateWithContext(w http.ResponseWriter, templateName string, data interface{}) bool {
    var buf bytes.Buffer
    if err := GetTemplates().ExecuteTemplate(&buf, templateName, data); err != nil {
        logger.Errorf("Failed to execute template '%s', error: %s", templateName, err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return false
    } else {
        // at this point we ignore error
        w.Write(buf.Bytes())
    }
    return true
}

func serveErrorMsg(w http.ResponseWriter, msg string) {
    http.Error(w, msg, http.StatusBadRequest)
}

func readConfig(configFile string) error {
    b, err := ioutil.ReadFile(configFile)
    if err != nil {
        return err
    }
    err = json.Unmarshal(b, &config)
    if err != nil {
        return err
    }
    return nil
}
