package internal

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"golang.org/x/term"
)

type game struct {
	fdin int
	old  *term.State

	line string
}

func get_words() []string {
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
	fmt.Print(restore, up, clear)
	os.Exit(code)
}

func Init() *game {
	words := get_words()
	line := strings.Join(words, " ")

	fmt.Print(save)

	fdin := int(os.Stdin.Fd())
	old, err := term.MakeRaw(fdin)
	if err != nil {
		panic(err)
	}

	g := game{}
	g.fdin = fdin
	g.old = old
	g.line = line
	// defer term.Restore(fdin, old)
	// g.start_game_loop(line)
	return &g
}

func (g *game) Run() {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	i := 0
	draw_with_border("", g.line, gray+g.line)
	for i != len(g.line) {
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if key == quit_key {
			break
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
			input += string(key)
			i += 1
		}
		print(input, g.line)
	}
	// TODO: print stats
	defer g.Exit(0, nil)
}
