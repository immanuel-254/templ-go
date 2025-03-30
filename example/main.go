package main

import (
	"context"
	"os"
	"templgo"
)

func main() {
	err := templgo.GenerateJS("index.js", WithoutParameters(), WithParameters("", "", 0))

	if err != nil {
		panic(err)
	}

	err = templgo.GenerateCSS("index.css", red(), text_white())

	if err != nil {
		panic(err)
	}

	comp := templgo.GenerateComponent(nil, Home())
	Base(comp).Render(context.Background(), os.Stdout)
}
