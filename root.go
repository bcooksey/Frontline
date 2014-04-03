package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
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
    err = db.QueryRow("SELECT id FROM users WHERE email = ?", r.FormValue("email")).Scan(&id)

    switch {
    case err == sql.ErrNoRows:
        logger.Error("No user with that ID.")
        serveErrorMsg(w, "Could not login")
    case err != nil:
        logger.Errorf("%v", err)
        serveErrorMsg(w, "Could not login")
    default:
        logger.Noticef("Username is %d\n", id)
    }

    // TODO: Generate random session token and store in session (everyone else will need to check session token)

    // cb := "http://" + r.Host + "/" + "?" + q
    // logger.Noticef("handleLogin: cb=%s\n", cb)

    // cookie := &SecureCookieValue{TwitterTemp: tempCred.Secret}
    // setSecureCookie(w, cookie)
    // http.Redirect(w, r, oauthClient.AuthorizationURL(tempCred, nil), 302)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
    //deleteSecureCookie(w)
    //http.Redirect(w, r, redirect, 302)
}

func deleteSecureCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "deleted",
		MaxAge: WeekInSeconds,
		Path:   "/",
	}
	http.SetCookie(w, cookie)
}
