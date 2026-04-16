package internal

const (
	clear = "\033[0J"
	reset = "\033[0m"
	red   = "\033[31m"
	blue  = "\033[34m"
	cyan  = "\033[36m"
	gray  = "\033[90m"

	save    = "\033[s"
	restore = "\033[u"

	up   = "\033[1A"
	back = "\033[%dD"

	quit_key = '!'
)
