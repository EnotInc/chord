package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
	blue = "\033[34m"
	red  = "\033[31m"
	gray = "\033[90m"

	save    = "\033[s"
	restore = "\033[u"

	back = "\033[%dD\033[1A"
)

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

func main() {

	words := get_words()
	line := strings.Join(words, " ")

	fmt.Print(save)
	fdin := int(os.Stdin.Fd())

	old, err := term.MakeRaw(fdin)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fdin, old)

	raw_input := ""
	reader := bufio.NewReader(os.Stdin)
	i := 0
	draw_with_border("", line, gray+line)
	for i != len(line) {
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if !is_key(key) {
			continue
		}

		if key == 127 { // backspace key
			if i > 0 {
				i -= 1
				raw_input = raw_input[:i]
			}
		} else {
			raw_input += string(key)
			i += 1
		}
		print(raw_input, line)
	}
	// TODO: print stats
	fmt.Print("\n\n")
}

func print(input string, line string) {
	output := ""
	for i, ch := range input {
		if line[i] != byte(ch) {
			output += red
			if ch == ' ' {
				ch = '_'
			}
		}
		output += string(ch)
		output += blue
	}
	output += gray
	output += line[len(input):]

	draw_with_border(input, line, output)
}

func draw_with_border(input string, line string, output string) {
	move_back := fmt.Sprintf(back, len(line)-len(input)+2)
	border := strings.Repeat("-", len(line)+2)
	fmt.Printf("%s+%s+\n| %s%s |\n+%s+%s", restore, border, blue, output, border, move_back)
}
