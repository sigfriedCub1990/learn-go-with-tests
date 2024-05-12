package blogposts

import "io"

type Post struct {
	Title string
}

func newPost(f io.Reader) (Post, error) {
	postData, err := io.ReadAll(f)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
