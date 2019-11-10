package prompt

import (
	"bufio"
	"github.com/howeyc/gopass"
	"log"
	"os"
	"strings"
)

type Prompter struct {
	br *bufio.Reader
}

func NewPrompter() *Prompter {
	return &Prompter{bufio.NewReader(os.Stdin)}
}

func (pr *Prompter) PromptTrimmed(caption string) string {
	os.Stderr.WriteString(caption)
	text, err := pr.br.ReadString('\n')
	if err != nil {
		os.Stderr.WriteString("\n")
		log.Fatalln(err)
	}
	return strings.TrimSpace(text)
}

func (pr *Prompter) PromptMasked(caption string) string {
	os.Stderr.WriteString(caption)
	text, err := gopass.GetPasswd()
	os.Stderr.WriteString("\r")
	if err != nil {
		log.Fatalln(err)
	}
	return string(text)
}
