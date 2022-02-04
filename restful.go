package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct { // https://mholt.github.io/json-to-go/
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getPost() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// bodyStr := string(bodyBytes)
	// fmt.Println(bodyStr)

	var post Post
	json.Unmarshal(bodyBytes, &post)

	fmt.Println("GET BODY: ", post.Body)
}

func savePost() {
	post := Post{1, 101, "Merhaba Dünya", "Vücutçu Yalvaç"}
	jsonPost, _ := json.Marshal(post)

	response, _ := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonPost))

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyStr := string(bodyBytes)
	fmt.Println("String JSON: ", bodyStr)

	var savedPost Post
	json.Unmarshal(bodyBytes, &savedPost)

	fmt.Println(savedPost.Body)
}

func main() {
	getPost()
	savePost()
}
