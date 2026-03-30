package internal

const (
	reset = "\033[0m"
	blue  = "\033[34m"
	red   = "\033[31m"
	gray  = "\033[90m"

	save    = "\033[s"
	restore = "\033[u"

	back = "\033[%dD\033[1A"

	quit_key = '!'
)
