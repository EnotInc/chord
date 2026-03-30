package internal

import (
	"fmt"
	"strings"
)

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

func build_border(message string, line_len int) string {
	border := strings.Repeat("-", (line_len-len(message))/2)
	shift_border := ""
	if line_len%2 == 0 {
		shift_border = "-"
	}
	return reset + "+" + border + " " + message + " " + border + shift_border + "+"
}

func draw_with_border(input string, line string, output string) {
	move_back := fmt.Sprintf(back, len(line)-len(input)+2)
	upper_border := build_border("[ chord ]", len(line))
	lower_border := build_border("press <"+string(quit_key)+"> to quit", len(line))
	fmt.Printf("%s %s\n %s| %s%s %s|\n %s%s", restore, upper_border, reset, blue, output, reset, lower_border, move_back)
}
