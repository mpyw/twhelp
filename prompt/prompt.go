package prompt

import (
    "os"
    "bufio"
    "strings"
    "github.com/howeyc/gopass"
)

type Prompter struct {
    bs *bufio.Scanner
}

func NewPrompter() *Prompter {
    return &Prompter { bufio.NewScanner(os.Stdin) }
}

func (pr *Prompter) PromptTrimmed(caption string) string {
    os.Stderr.WriteString(caption)
    pr.bs.Scan()
    return strings.TrimSpace(pr.bs.Text())
}

func (pr *Prompter) PromptMasked(caption string) string {
    os.Stderr.WriteString(caption)
    return string(gopass.GetPasswd())
}
