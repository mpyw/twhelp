package utility

import (
    "fmt"
    "github.com/mitchellh/go-homedir"
    "gopkg.in/ini.v1"
    "log"
    "os"
    "strings"
)

type Util struct {
    ConfigPath     string
    CustomAppNames string
    Apps           map[string]*[]string
}

func NewUtil() *Util {
    home, err := homedir.Dir()
    if err != nil {
        log.Fatalln(err)
    }
    util := &Util{
        ConfigPath: fmt.Sprintf(
            "%s%s.twhelp.ini",
            home,
            string(os.PathSeparator),
        ),
        Apps: make(map[string]*[]string, 8),
    }
    util.Apps["android"] = &[]string{
        "3nVuSoBZnx6U4vzUxf5w",
        "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
    }
    util.Apps["win"] = &[]string{
        "TgHNMa7WZE7Cxi1JbkAMQ",
        "SHy9mBMBPNj3Y17et9BF4g5XeqS4y3vkeW24PttDcY",
    }
    util.Apps["wp"] = &[]string{
        "yN3DUNVO0Me63IAQdhTfCA",
        "c768oTKdzAjIYCmpSNIdZbGaG0t6rOhSFQP0S5uC79g",
    }
    util.Apps["google"] = &[]string{
        "iAtYJ4HpUVfIUoNnif1DA",
        "172fOpzuZoYzNYaU3mMYvE8m8MEyLbztOdbrUolU",
    }
    util.Apps["iphone"] = &[]string{
        "IQKbtAYlXLripLGPWd0HUA",
        "GgDYlkSvaPxGxC4X8liwpUoqKwwr3lCADbz8A7ADU",
    }
    util.Apps["ipad"] = &[]string{
        "CjulERsDeqhhjSme66ECg",
        "IQWdVyqFxghAtURHGeGiWAsmCAGmdW3WmbEx6Hck",
    }
    util.Apps["mac"] = &[]string{
        "3rJOl1ODzm9yZy63FACdg",
        "5jPoQ5kQvMJFDYRNE8bQ4rHuds4xJqhvgNJM4awaE8",
    }
    util.Apps["deck"] = &[]string{
        "yT577ApRtZw51q4NPMPPOQ",
        "3neq3XqN5fO3obqwZoajavGFCUrC42ZfbrLXy5sCv8",
    }
    customAppNames := make([]string, 0)
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
            ck, ckerr := section.GetKey("consumer_key")
            if ckerr != nil {
                log.Fatalln(fmt.Sprintf(`"consumer_key" for %s does not exist`, name))
            }
            cs, cserr := section.GetKey("consumer_secret")
            if cserr != nil {
                log.Fatalln(fmt.Sprintf(`"consumer_secret" for %s does not exist`, name))
            }
            util.Apps[name] = &[]string{ck.String(), cs.String()}
            customAppNames = append(customAppNames, name)
        }
    }
    if len(customAppNames) > 0 {
        util.CustomAppNames = fmt.Sprintf(
            "Config File: %s\nCustom Apps: %s\n",
            util.ConfigPath,
            strings.Join(customAppNames, ", "),
        )
    } else {
        util.CustomAppNames = fmt.Sprintf(
            `Your own applications also can be defined in %s
Example:

[my_app_01]
consumer_key    = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
consumer_secret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
`, util.ConfigPath)
    }
    return util
}

func (util *Util) Usage() {
    os.Stderr.WriteString(fmt.Sprintf(`Usage: %s [options]
Options:
  -h, --help          Show help.

[ Output Format ]

  Default             Output line by line.
  -i, --ini           Output as INI.
  -y, --yaml          Output as YAML.
  -j, --json          Output as JSON.
  -a, --array         Output as array that compatible with most languages.
  -A, --assoc         Output as PHP-style associative array.
  -e, --env           Output as environmental uppercase variables.

[ OAuth Process ]

  Default             xAuth manipulation with OAuth scraping.
  -x, --xauth         Pure xAuth. Only available with official keys.
  -o, --oauth         Pure OAuth. You have to authorize via web browser.

[ OAuth Credentials ]

  Insufficient components are required to input via STDIN.
  Password is masked.

  --ck  <value>       Specify consumer_key in advance.
  --cs  <value>       Specify consumer_secret in advance.
  --sn  <value>       Specify screen_name or email in advance.
  --pw  <value>       Specify password in advance. (Not masked, DEPRECATED)
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

%s
`, os.Args[0], util.CustomAppNames))
}
