# twhelp-go

Twitter OAuth CLI Helper distributed by Golang cross-compilation.
It has migrated from [mpyw/twhelp](https://github.com/mpyw/twhelp).

## Requirements

**Nothing**.  
Feel free to download from [releases](https://github.com/mpyw/twhelp-go/releases).

- **[mpyw/twhelp-go](https://github.com/mpyw/twhelp-go)**

## Usage

```ShellSession
mpyw@localhost:~$ ./twhelp -h
Usage: ./twhelp [options]
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

Your own applications also can be defined in /Users/mpyw/.twhelp.ini
Refer to the documentation.
mpyw@localhost:~$
```

## `~/.twhelp.ini` schema

```ini
[my_app_01]
ck = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
cs = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

[my_app_02]
ck = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
cs = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

...
```
