package main

import (
    "net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
    model := struct{}{}
    ExecTemplate(w, tmplMain, model)
    // fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    // redirect := strings.TrimSpace(r.FormValue("redirect"))
    // if redirect == "" {
    // 	serveErrorMsg(w, fmt.Sprintf("Missing redirect value for /login"))
    // 	return
    // }
    // q := url.Values{
    // 	"redirect": {redirect},
    // }.Encode()

    // cb := "http://" + r.Host + "/oauthtwittercb" + "?" + q
    // //fmt.Printf("handleLogin: cb=%s\n", cb)
    // tempCred, err := oauthClient.RequestTemporaryCredentials(http.DefaultClient, cb, nil)
    // if err != nil {
    // 	http.Error(w, "Error getting temp cred, "+err.Error(), 500)
    // 	return
    // }
    // cookie := &SecureCookieValue{TwitterTemp: tempCred.Secret}
    // setSecureCookie(w, cookie)
    // http.Redirect(w, r, oauthClient.AuthorizationURL(tempCred, nil), 302)
}

// url: GET /logout?redirect=$redirect
func handleLogout(w http.ResponseWriter, r *http.Request) {
    // redirect := strings.TrimSpace(r.FormValue("redirect"))
    // if redirect == "" {
    // 	serveErrorMsg(w, fmt.Sprintf("Missing redirect value for /logout"))
    // 	return
    // }
    // deleteSecureCookie(w)
    // http.Redirect(w, r, redirect, 302)
}
