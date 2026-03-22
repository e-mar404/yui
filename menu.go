package main

type menu int

const (
	all menu = iota
	explicit
	aur
)

func (m menu) String() string {
	switch m {
	case 0:
		return "All (a)"

	case 1:
		return "Explicit (e)"

	case 2:
		return "AUR (u)"

	default:
		return ""
	}
}
