package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/securecookie"
    _ "log"
    "net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
    model := struct{}{}
    ExecTemplate(w, tmplMain, model)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
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
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(fmt.Sprintf(`{"user": {"id": %d, "name": "%s"}}`, id, username)))
}
