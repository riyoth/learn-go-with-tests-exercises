package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/riyoth/learn-go-with-tests-exercises/17_read_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("hi")},
		"hello_world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("go %d posts, wanted %d posts", len(posts), len(fs))
	}
}
