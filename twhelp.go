package main

import(
    "fmt"
    "os"
    "log"
    "regexp"
    "github.com/pborman/getopt"
    "net/url"
    "./utility"
    "./prompt"
    "./oauth"
    "./session"
)

func main() {
    var (
        format string
        t *oauth.Credential
        help bool
        twist bool
        var_ bool
        xauth bool
        oauth_ bool
        ck string
        cs string
        sn string
        pw string
        app string
    )
    util := utility.NewUtil()
    prompter := prompt.NewPrompter()
    getopt.BoolVarLong(&help, "help", 'h')
    getopt.BoolVarLong(&twist, "twist", 't')
    getopt.BoolVarLong(&var_, "var", 'v')
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
            t = (&oauth.Credential{ck, cs, "", ""}).RenewWithAccessToken(map[string]string {
                "x_auth_mode": "client_auth",
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
                    "session[password]": {pw},
                }))
                if matches == nil {
                    log.Fatalln("Wrong username or password")
                }
                verifier := string(matches[1])
                t = t.RenewWithAccessToken(map[string]string {}, &verifier)
            }
        }
    } else {
        t = (&oauth.Credential{ck, cs, "", ""}).RenewWithRequestToken()
        uri := "https://api.twitter.com/oauth/authorize?force_login=1&oauth_token=" + t.OAuthToken
        os.Stderr.WriteString("Access here for authorization: " + uri + "\n")
        verifier := prompter.PromptTrimmed("PIN: ")
        t = t.RenewWithAccessToken(map[string]string {}, &verifier)
        fmt.Println(t)
    }

if (twist) {
format = `$to = new TwistOAuth(
    "%s",
    "%s",
    "%s",
    "%s"
);
`
} else if (var_) {
format = `$ck = "%s";
$cs = "%s";
$ot = "%s";
$os = "%s";
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
