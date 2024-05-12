package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/rsgarxia/blogposts"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("read files from filysyste", func(t *testing.T) {
		fs := fstest.MapFS{
			"hellow-world.md":  {Data: []byte("hi")},
			"hellow-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal("this should fail")
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("returns an error if reading files fail", func(t *testing.T) {
		fs := StubFailingFs{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Errorf("expected an error, got none")
		}
	})
}
