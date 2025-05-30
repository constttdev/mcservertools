package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"

	"github.com/constt/mcservertools/cmd"
)

func main() {
	figure := figure.NewFigure("MC Server Tools", "slant", true)
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue(figure.String()))

	cmd.Execute()
}
