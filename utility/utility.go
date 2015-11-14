package utility

import (
    "fmt"
    "log"
    "os"
    "github.com/mitchellh/go-homedir"
    "gopkg.in/ini.v1"
)

type Util struct {
    ConfigPath string
    Apps map[string]*[]string
}

func NewUtil() *Util {
    home, err := homedir.Dir()
    if err != nil {
        log.Fatalln(err)
    }
    util := &Util {
        ConfigPath: fmt.Sprintf(
            "%s%s.twhelp.ini",
            home,
            string(os.PathSeparator),
        ),
        Apps: make(map[string]*[]string, 8),
    }
    util.Apps["android"] = &[]string {
        "3nVuSoBZnx6U4vzUxf5w",
        "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
    }
    util.Apps["win"] = &[]string {
        "TgHNMa7WZE7Cxi1JbkAMQ",
        "SHy9mBMBPNj3Y17et9BF4g5XeqS4y3vkeW24PttDcY",
    }
    util.Apps["wp"] = &[]string {
        "yN3DUNVO0Me63IAQdhTfCA",
        "c768oTKdzAjIYCmpSNIdZbGaG0t6rOhSFQP0S5uC79g",
    }
    util.Apps["google"] = &[]string {
        "iAtYJ4HpUVfIUoNnif1DA",
        "172fOpzuZoYzNYaU3mMYvE8m8MEyLbztOdbrUolU",
    }
    util.Apps["iphone"] = &[]string {
        "IQKbtAYlXLripLGPWd0HUA",
        "GgDYlkSvaPxGxC4X8liwpUoqKwwr3lCADbz8A7ADU",
    }
    util.Apps["ipad"] = &[]string {
        "CjulERsDeqhhjSme66ECg",
        "IQWdVyqFxghAtURHGeGiWAsmCAGmdW3WmbEx6Hck",
    }
    util.Apps["mac"] = &[]string {
        "3rJOl1ODzm9yZy63FACdg",
        "5jPoQ5kQvMJFDYRNE8bQ4rHuds4xJqhvgNJM4awaE8",
    }
    util.Apps["deck"] = &[]string {
        "yT577ApRtZw51q4NPMPPOQ",
        "3neq3XqN5fO3obqwZoajavGFCUrC42ZfbrLXy5sCv8",
    }
    if _, err := os.Stat(util.ConfigPath); err == nil {
        cfg, err := ini.Load(util.ConfigPath)
        if err != nil {
            log.Fatalln(err)
        }
        for _, name := range cfg.SectionStrings() {
            if name == "DEFAULT" {
                continue
            }
            section := cfg.Section(name)
            ck, ckerr := section.GetKey("ck")
            if ckerr != nil {
                log.Fatalln(fmt.Sprintf(`"ck" for %s does not exist`, name))
            }
            cs, cserr := section.GetKey("cs")
            if cserr != nil {
                log.Fatalln(fmt.Sprintf(`"cs" for %s does not exist`, name))
            }
            util.Apps[name] = &[]string{ck.String(), cs.String()}
        }
    }
    return util
}

func (util *Util) Usage() {
os.Stderr.WriteString(fmt.Sprintf(`Usage: %s [options]
Options:
  -h, --help          Show help.

[ Output Format ]

  Default             Output line by line.
  -t, --twist         Output as TwistOAuth-style constrctive code.
  -v, --var           Output as variable line by line.

[ OAuth Process ]

  Default             DirectOAuth. (xAuth manipulation with OAuth)
  -x, --xauth         Pure xAuth. Only available with official keys.
  -o, --oauth         Pure OAuth. You have to authorize via web browser.

[ OAuth Credentials ]

  Insufficient components are required to input via STDIN.
  Password is masked.

  --ck  <value>       Specify consumer_key in advance.
  --cs  <value>       Specify consumer_secret in advance.
  --sn  <value>       Specify screen_name or email in advance.
  --pw  <value>       Specify password in advance. (DEPRECATED)
  --app <value>       Speficy consumer_key and consumer_secret with app name.

                      app name | full name
                      ------------------------------------
                      android  | Twitter for Andriod
                      win      | Twitter for Andriod
                      wp       | Twitter for Windows Phone
                      google   | Twitter for Google TV
                      iphone   | Twitter for iPhone
                      ipad     | Twitter for iPad
                      mac      | Twitter for Mac
                      deck     | TweetDeck

Your own applications also can be defined in %s
Refer to the documentation.
`, os.Args[0], util.ConfigPath))
}
