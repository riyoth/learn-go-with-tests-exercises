package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/riyoth/learn-go-with-tests-exercises/18_html_templates/blogrenderer/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			body:        "this is a post",
			Description: "This is a description",
			tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}
