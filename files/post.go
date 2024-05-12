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

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readLine()[len(titleSeparator):]
	descriptionLine := readLine()[len(descriptionSeparator):]

	return Post{Title: titleLine, Description: descriptionLine}, nil
}
