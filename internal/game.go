package internal

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type game struct {
	time   time.Time
	typed  int
	errors int

	fdin int
	old  *term.State

	line string
}

func getWords() []string {
	all_words := strings.Split(words(), " ")
	var words []string
	for range 10 {
		rnd := rand.IntN(1000)
		words = append(words, all_words[rnd])
	}

	return words
}

func is_key(key rune) bool {
	return ('a' <= key && key <= 'z') || ('A' <= key && key <= 'Z') || key == ' ' || key == 127
}

func (g *game) Exit(code int, message *string) {
	msg := ""
	if message != nil {
		msg = *message
	}
	fmt.Print(restore, up, clear, msg)
	os.Exit(code)
}

func Init() *game {
	words := getWords()
	line := strings.Join(words, " ")

	fmt.Print(save)

	fdin := int(os.Stdin.Fd())
	old, err := term.MakeRaw(fdin)
	if err != nil {
		panic(err)
	}

	g := game{
		fdin: fdin,
		old:  old,
		line: line,

		time:   time.Now(),
		typed:  0,
		errors: 0,
	}
	return &g
}

func (g *game) Run() {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	i := 0
	drawWithBorder("", g.line, gray+g.line)
	for i != len(g.line) {
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if key == quit_key {
			quit := red + "DNF" + reset
			g.Exit(0, &quit)
		}
		if !is_key(key) {
			continue
		}

		if key == 127 { // backspace key
			if i > 0 {
				i -= 1
				input = input[:i]
			}
		} else {
			g.typed += 1
			if key != rune(g.line[i]) {
				g.errors += 1
			}

			input += string(key)
			i += 1
		}
		print(input, g.line)
	}
	defer g.Exit(0, g.getStats())
}

func (g *game) getStats() *string {
	t := time.Since(g.time).String()
	stats := fmt.Sprintf("Ended in: "+cyan+"%s"+reset+"\nTyped "+blue+"%d"+reset+" symbols with "+red+"%d"+reset+" errors in total", t, g.typed, g.errors)
	return &stats
}
