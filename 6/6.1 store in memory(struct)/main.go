package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello hoge", Author: "tanaka"}
	post2 := Post{Id: 2, Content: "foobar", Author: "takahashi"}
	post3 := Post{Id: 3, Content: "good morning", Author: "Bob"}
	post4 := Post{Id: 4, Content: "bye", Author: "tanaka"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostByAuthor["tanaka"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["Bob"] {
		fmt.Println(post)
	}

}
