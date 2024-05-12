package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
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
	t.Run("read files from filesystem", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
			secondBody = `Title: Post 1
Description: Description 1
Tags: rust, borrow-checker`
		)
		fs := fstest.MapFS{
			"hellow-world.md":  {Data: []byte(firstBody)},
			"hellow-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal("this should not fail")
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
		})
	})

	t.Run("returns an error if reading files fail", func(t *testing.T) {
		fs := StubFailingFs{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Errorf("expected an error, got none")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
