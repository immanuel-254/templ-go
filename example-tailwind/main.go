package main

import (
	"context"
	"os"
	"templgo"
)

func main() {
	templgo.GenerateTailwindInputCss("index.css", ".btn{@apply flex justify-center items-center p-4 bg-gray-100 rounded-lg shadow-md;}", ".mn{@apply flex justify-center items-center;}")

	comp := templgo.GenerateComponent(nil, Home())
	comp.Render(context.Background(), os.Stdout)
}

//tailwindcss -i index.css -o static/styles.css --minify
