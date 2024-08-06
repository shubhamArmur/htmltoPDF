package helper

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func Highli(markdown []byte) string {
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"), // You can choose your preferred style
				highlighting.WithGuessLanguage(true),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert(markdown, &buf); err != nil {
		fmt.Println("Error converting markdown to HTML:", err)
		return ""
	}

	fmt.Println(buf.String())
	return buf.String()
}
