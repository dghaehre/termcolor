package termcolor

import (
	"fmt"
)

type Color string

const (
	Black      Color = "\u001b[30m"
	Red              = "\u001b[31m"
	Green            = "\u001b[32m"
	Yellow           = "\u001b[33m"
	Blue             = "\u001b[34m"
	ColorReset       = "\u001b[0m"
)

func Str(color Color, message string) string {
	return fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset))
}

func Println(color Color, message string) {
	s := fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset))
	fmt.Println(s)
}
