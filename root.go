package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/securecookie"
    "net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
    ExecTemplate(w, tmplMain)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ExecTemplate(w, "login.html")
        return
    }

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/frontline", config.Dbuser, config.Dbpass))
    if err != nil {
        logger.Error(err.Error())
        serveErrorMsg(w, "Could not login")
    }

    err = r.ParseForm()
    if err != nil {
        logger.Error(err.Error())
        serveErrorMsg(w, "Could not login")
    }

    var id int
    var username string
    err = db.QueryRow("SELECT id, name FROM users WHERE email = ?", r.FormValue("email")).Scan(&id, &username)

    switch {
    case err == sql.ErrNoRows:
        logger.Error("No user with that ID.")
        serveErrorMsg(w, "Could not login")
        return
    case err != nil:
        logger.Errorf("%v", err)
        serveErrorMsg(w, "Could not login")
        return
    default:
        logger.Noticef("User %d logged in", id)
    }

    session, _ := sessionStore.Get(r, fmt.Sprintf("user-%d", id))
    session.Values["token"] = securecookie.GenerateRandomKey(256)
    session.Values["userId"] = id
    session.Values["userName"] = username

    session.Save(r, w)
    http.Redirect(w, r, "/app", 301)
}

func handleApp(w http.ResponseWriter, r *http.Request) {
    ExecTemplate(w, "app.html")
    return
}
