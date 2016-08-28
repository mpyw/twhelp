package main

import (
    "./oauth"
    "./prompt"
    "./session"
    "./utility"
    "fmt"
    "github.com/pborman/getopt"
    "log"
    "net/url"
    "os"
    "regexp"
)

func main() {
    var (
        format string
        t      *oauth.Credential
        help   bool
        ini    bool
        yaml   bool
        array  bool
        assoc  bool
        json   bool
        env    bool
        xauth  bool
        oauth_ bool
        ck     string
        cs     string
        sn     string
        pw     string
        app    string
    )
    util := utility.NewUtil()
    prompter := prompt.NewPrompter()
    getopt.BoolVarLong(&help, "help", 'h')
    getopt.BoolVarLong(&ini, "ini", 'i')
    getopt.BoolVarLong(&yaml, "yaml", 'y')
    getopt.BoolVarLong(&array, "array", 'a')
    getopt.BoolVarLong(&assoc, "assoc", 'A')
    getopt.BoolVarLong(&json, "json", 'j')
    getopt.BoolVarLong(&env, "env", 'e')
    getopt.BoolVarLong(&oauth_, "oauth", 'o')
    getopt.BoolVarLong(&xauth, "xauth", 'x')
    getopt.StringVarLong(&ck, "ck", 0, "")
    getopt.StringVarLong(&cs, "cs", 0, "")
    getopt.StringVarLong(&sn, "sn", 0, "")
    getopt.StringVarLong(&pw, "pw", 0, "")
    getopt.StringVarLong(&app, "app", 0, "")
    getopt.SetUsage(util.Usage)
    getopt.CommandLine.Parse(os.Args)
    if help {
        util.Usage()
        os.Exit(0)
    }
    if ck == "" && cs == "" && app != "" {
        appks, ok := util.Apps[app]
        if !ok {
            log.Fatalln("Application not found: " + app)
        }
        ck = (*appks)[0]
        cs = (*appks)[1]
    }
    if ck == "" {
        ck = prompter.PromptTrimmed("consumer_key: ")
    }
    if cs == "" {
        cs = prompter.PromptTrimmed("consumer_secret: ")
    }
    if xauth || !oauth_ {
        if sn == "" {
            sn = prompter.PromptTrimmed("screen_name: ")
        }
        if pw == "" {
            pw = prompter.PromptMasked("password: ")
        }
        if xauth {
            t = (&oauth.Credential{ck, cs, "", ""}).RenewWithAccessToken(map[string]string{
                "x_auth_mode":     "client_auth",
                "x_auth_username": sn,
                "x_auth_password": pw,
            }, nil)
        } else {
            t = (&oauth.Credential{ck, cs, "", ""}).RenewWithRequestToken()
            sess := session.NewSession()
            uri := "https://api.twitter.com/oauth/authorize?force_login=1&oauth_token=" + t.OAuthToken
            {
                expr := `<input name="authenticity_token" type="hidden" value="([^"]+)">`
                matches := regexp.MustCompile(expr).FindSubmatch(sess.Get(uri))
                if matches == nil {
                    log.Fatalln("Could not find authenticity_token")
                }
                sess.SetAuthenticityToken(string(matches[1]))
            }
            {
                expr := `<code>([^<]+)</code>`
                matches := regexp.MustCompile(expr).FindSubmatch(sess.Post(uri, &url.Values{
                    "session[username_or_email]": {sn},
                    "session[password]":          {pw},
                }))
                if matches == nil {
                    log.Fatalln("Wrong username or password. Otherwise, you may have to verify your email address.")
                }
                verifier := string(matches[1])
                t = t.RenewWithAccessToken(map[string]string{}, &verifier)
            }
        }
    } else {
        t = (&oauth.Credential{ck, cs, "", ""}).RenewWithRequestToken()
        uri := "https://api.twitter.com/oauth/authorize?force_login=1&oauth_token=" + t.OAuthToken
        os.Stderr.WriteString("Access here for authorization: " + uri + "\n")
        verifier := prompter.PromptTrimmed("PIN: ")
        t = t.RenewWithAccessToken(map[string]string{}, &verifier)
    }

    if ini {
        format = `consumer_key        = "%s"
consumer_secret     = "%s"
access_token        = "%s"
access_token_secret = "%s"
`
    } else if yaml {
        format = `consumer_key:        "%s"
consumer_secret:     "%s"
access_token:        "%s"
access_token_secret: "%s"
`
    } else if array {
        format = `[
    "%s",
    "%s",
    "%s",
    "%s"
]
`
    } else if assoc {
        format = `[
    "consumer_key"        => "%s",
    "consumer_secret"     => "%s",
    "access_token"        => "%s",
    "access_token_secret" => "%s",
]
`
    } else if json {
        format = `{
    "consumer_key":        "%s",
    "consumer_secret":     "%s",
    "access_token":        "%s",
    "access_token_secret": "%s"
}
`
    } else if env {
        format = `CONSUMER_KEY="%s"
CONSUMER_SECRET="%s"
ACCESS_TOKEN="%s"
ACCESS_TOKEN_SECRET="%s"
`
    } else {
        format = `%s
%s
%s
%s
`
    }

    fmt.Printf(format, t.ConsumerKey, t.ConsumerSecret, t.OAuthToken, t.OAuthTokenSecret)

}
