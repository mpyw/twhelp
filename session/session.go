package session

import (
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "io/ioutil"
    "../cabundle"
)

type Session struct {
    Client http.Client
    AuthenticityToken string
}

func NewSession() *Session {
    jar, _ := cookiejar.New(nil)
    return &Session{Client: http.Client{
        Jar: jar,
        Transport: cabundle.GetTransport(),
    }}
}

func (sess *Session) SetAuthenticityToken(token string) {
    sess.AuthenticityToken = token
}

func (sess *Session) Get(uri string) []byte {
    res, err := sess.Client.Get(uri)
    if err != nil {
        log.Fatalln(err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    return body
}

func (sess *Session) Post(uri string, values *url.Values) []byte {
    if sess.AuthenticityToken != "" {
        values.Add("authenticity_token", sess.AuthenticityToken)
    }
    res, err := sess.Client.PostForm(uri, *values)
    if err != nil {
        log.Println(err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    return body
}
