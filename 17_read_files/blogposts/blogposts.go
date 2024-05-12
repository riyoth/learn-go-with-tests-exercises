package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSsytem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSsytem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(fileSsytem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	post, err := newPost(postFile)
	return post, nil
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	title := readLine(scanner)[len(titleSeparator):]
	description := readLine(scanner)[len(descriptionSeparator):]
	tags := strings.Split(readLine(scanner)[len(tagsSeparator):], ", ")
	body := readBody(scanner)

	posts := Post{Title: title, Description: description, Tags: tags, Body: body}
	return posts, nil

}

func readBody(scanner *bufio.Scanner) string {
	readLine(scanner)
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
