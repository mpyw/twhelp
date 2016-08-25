_twhelp() {
    local context curcontext=$curcontext state line ret=1 capps=''
    typeset -A opt_args
    if [[ `2>&1 twhelp -h` =~ 'Custom Apps: (.*)' ]]; then
        capps=${match[1]//,/ }
    fi
    _arguments -C \
        '(-h --help -i -y -a -A -j --ini --yaml --array --assoc --json --ck --cs --sn --pw --app)'{-h,--help}'[Show help.]' \
        '(-i --ini -y -a -A -j -h --yaml --array --assoc --json --help)'{-i,--ini}'[Output as INI.]' \
        '(-y --yaml -i -a -A -j -h --ini --array --assoc --json --help)'{-y,--yaml}'[Output as YAML.]' \
        '(-a --array -i -y -A -j -h --ini --yaml --assoc --json --help)'{-a,--array}'[Output as array that compatible with most languages]' \
        '(-A --assoc -i -y -a -j -h --ini --yaml --array --json --help)'{-A,--assoc}'[Output as PHP-style associative array.]' \
        '(-j --json -i -y -a -A -h --ini --yaml --array --assoc --help)'{-j,--json}'[Output as JSON.]' \
        '(-x --xauth -o -h --oauth --help)'{-x,--xauth}'[Pure xAuth. Only available with official keys.]' \
        '(-o --oauth -x -h --xauth --help)'{-o,--oauth}'[Pure OAuth. You have to authorize via web browser.]' \
        '(--ck -h --app --help)--ck=[Specify consumer_key in advance.]: : ' \
        '(--cs -h --app --help)--cs=[Specify consumer_secret in advance.]: : ' \
        '(--sn -h --help)--sn=[Specify screen_name or email in advance.]: : ' \
        '(--pw -h --help)--pw=[Specify password in advance. (Not masked, DEPRECATED)]: : ' \
        '(--app -h --ck --cs --help)--app=[Speficy consumer_key and consumer_secret with app name.]: :(android win wp google iphone ipad mac deck '$capps')' \
        '(-)*:arguments:->args' \
        && ret=0
    return ret
}
compdef _twhelp twhelp
