package internal

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"golang.org/x/term"
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

func Play() {
	words := get_words()
	line := strings.Join(words, " ")

	fmt.Print(save)

	fdin := int(os.Stdin.Fd())
	old, err := term.MakeRaw(fdin)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fdin, old)
	start_game_loop(line)
}

func start_game_loop(line string) {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	i := 0
	draw_with_border("", line, gray+line)
	for i != len(line) {
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
		print(input, line)
	}
	// TODO: print stats
	fmt.Print("\n\r\n\r")
}
