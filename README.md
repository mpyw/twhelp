# twhelp

Twitter OAuth CLI Helper distributed by Golang cross-compilation.  

## Requirements

**Nothing**.  
Feel free to download from [releases](https://github.com/mpyw/twhelp/releases).

...Oops, x64(amd64) CPU architecture required at least.

## Usage

```ShellSession
mpyw@localhost:~$ ./twhelp -h
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
                      win      | Twitter for Windows
                      wp       | Twitter for Windows Phone
                      google   | Twitter for Google TV
                      iphone   | Twitter for iPhone
                      ipad     | Twitter for iPad
                      mac      | Twitter for Mac
                      deck     | TweetDeck

Your own applications also can be defined in /Users/mpyw/.twhelp.ini
Example:

[my_app_01]
consumer_key    = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
consumer_secret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

mpyw@localhost:~$
```
