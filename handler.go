package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("hogefugastring"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index")

	session, err := store.Get(r, "twitter")
	if err != nil {
		// TODO: error handling
	}

	_, ok_ := session.Values["token"].(string)
	_, ok := session.Values["secret"].(string)
	// TODO
	if ok_ && ok {
		log.Println("logined")
	} else {
		log.Println("login")
		w.Header().Set("Pragma", "no-cache")
		http.ServeFile(w, r, "views/signin.html")
	}
}

func twitterGetTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("get token")

	tokenUrl := fmt.Sprintf("http://%s/callback", r.Host)
	token, reqUrl, err := twitter.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Fatal(err)
	}

	tokens[token.Token] = token

	http.Redirect(w, r, reqUrl, http.StatusTemporaryRedirect)
}

func twitterCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("callback")

	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	accessToken, err := twitter.AuthorizeToken(tokens[tokenKey], verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: error handling
	session, _ := store.Get(r, "twitter")
	session.Values["token"] = accessToken.Token
	session.Values["secret"] = accessToken.Secret
	session.Save(r, w)

	url := fmt.Sprintf("http://%s", r.Host)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
