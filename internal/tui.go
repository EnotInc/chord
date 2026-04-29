package internal

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	upper int = iota
	lower
)

func print(input string, line string) {
	var output strings.Builder
	for i, ch := range input {
		if line[i] != byte(ch) {
			output.WriteString(red)
			if ch == ' ' {
				ch = '\u00b7'
			}
		}
		output.WriteString(string(ch))
		output.WriteString(blue)
	}
	output.WriteString(gray)
	output.WriteString(line[len(input):])

	drawWithBorder(input, line, output.String())
}

func buildBorder(message string, line_len int, pos int) string {
	border := strings.Repeat("─", (line_len-len(message))/2)
	shift_border := ""
	if line_len%2 == 0 {
		shift_border = "─"
	}

	var l_corner = "┼"
	var r_corner = "┼"

	switch pos {
	case upper:
		l_corner = "┌"
		r_corner = "┐"
	case lower:
		l_corner = "└"
		r_corner = "┘"
	}
	return reset + l_corner + border + "─" + message + "─" + border + shift_border + r_corner
}

func drawWithBorder(input string, line string, output string) {
	move_back := fmt.Sprintf(back, len(line)-len(input)+2) + up
	upper_border := buildBorder(" chord ", len(line), upper)
	lower_border := buildBorder(" press <"+string(quit_key)+"> to quit ", len(line), lower)
	fmt.Printf("%s %s\n\r %s│ %s%s %s│\n\r %s%s", restore, upper_border, reset, blue, output, reset, lower_border, move_back)
}

func (g *game) getStats() *string {
	t := time.Since(g.time).Seconds()
	round := math.Round(t*100) / 100

	result := fmt.Sprintf("%.2f", round)
	speed := fmt.Sprintf("%.2f", float64(g.typed)/round*60)

	stats := fmt.Sprintf("Ended in: "+cyan+"%ss"+reset+"\nTyped "+blue+"%d"+reset+" symbols with "+red+"%d"+reset+" errors in total\nAverage speed: "+cyan+"%s CPM"+reset, result, g.typed, g.errors, speed)
	return &stats
}
