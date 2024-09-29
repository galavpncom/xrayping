package utils

const (
	ColorGreen  = "\033[32m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// WrapTextWithColor wraps the given text with the specified ANSI color code.
func WrapTextWithColor(text, color string) string {
	return color + text + ColorReset
}
