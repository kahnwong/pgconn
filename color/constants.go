package color

import "github.com/fatih/color"

var (
	Blue  = color.New(color.FgBlue).SprintFunc()
	Green = color.New(color.FgHiGreen).SprintFunc()
	Red   = color.New(color.FgRed).SprintFunc()
)
