package blogposts

import (
	"bufio"
	"io"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

type Post struct {
	Title, Description string
}

func newPost(f io.Reader) (Post, error) {
	scanner := bufio.NewScanner(f)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
	}, nil
}
