package color

const (
	Gray  = "\033[90m"
	Reset = "\033[0m"
)

func Paint(ansi, content string) string {
	return ansi + content + Reset
}
