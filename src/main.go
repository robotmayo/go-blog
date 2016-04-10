package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// BlogPost HOT MEMES 2016
type BlogPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// ResultResponse I DONT KNOW LOL
type ResultResponse struct {
	Error string `json:"error"`
	Ok    bool   `json:"ok"`
}

var posts []BlogPost

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", postsHandler)
	mux.HandleFunc("/new/post", newPostHandler)
	mux.HandleFunc("/post/", getSinglePostHandler)
	http.Handle("/", http.FileServer(http.Dir("/home/robtmayo/golang-work/blog/static")))
	http.ListenAndServe(":8000", mux)
}

func mess() {
	temp := []BlogPost{BlogPost{Title: "My Title", Body: "Some body"}}
	mrjson, err := json.Marshal(temp)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("IM MR JSON %s", string(mrjson))
}

func getSinglePostHandler(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	w.Header().Set("Content-Type", "application/javascript")
	var parts []string
	for _, s := range splitPath {
		if s != "" {
			parts = append(parts, s)
		}
	}

	fmt.Printf("%v %d %s", parts, len(parts), r.URL.Path)
	responseStruct := ResultResponse{Error: "", Ok: true}
	if len(parts) < 2 {
		responseStruct.Error = "Missing Id"
		responseStruct.Ok = false
		s, _ := json.Marshal(responseStruct)
		fmt.Fprintf(w, "%s", s)
		return
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		responseStruct.Error = "No Post Found"
		responseStruct.Ok = false
		s, _ := json.Marshal(responseStruct)
		fmt.Fprintf(w, "%s", s)
		return
	}
	if id >= len(posts) {
		responseStruct.Error = "No Post Found"
		responseStruct.Ok = false
		s, _ := json.Marshal(responseStruct)
		fmt.Fprintf(w, "%s", s)
		return
	}
	post := posts[id]
	responseJSON, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStruct.Error = "Internal Error"
		responseStruct.Ok = false
		s, _ := json.Marshal(responseStruct)
		fmt.Fprintf(w, "%s", s)
		return
	}
	fmt.Fprintf(w, "%s", string(responseJSON))
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	decoderRing := json.NewDecoder(r.Body)
	var newPost BlogPost
	err := decoderRing.Decode(&newPost)
	w.Header().Set("Content-Type", "application/javascript")
	resultResp := ResultResponse{Error: "", Ok: true}
	fmt.Printf("My error : %v %T", err, newPost.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resultResp.Error = "Malformed JSON"
		resultResp.Ok = false
		s, _ := json.Marshal(resultResp)
		fmt.Fprintf(w, "%s", string(s))
		return
	}

	if newPost.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		resultResp.Error = "Missing Title"
		resultResp.Ok = false
		s, _ := json.Marshal(resultResp)
		fmt.Fprintf(w, "%s", string(s))
		return
	}

	if newPost.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
		resultResp.Error = "Missing Body"
		resultResp.Ok = false
		s, _ := json.Marshal(resultResp)
		fmt.Fprintf(w, "%s", string(s))
		return
	}

	posts = append(posts, newPost)
	s, _ := json.Marshal(resultResp)
	fmt.Fprintf(w, "%s", string(s))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	if len(posts) == 0 {
		fmt.Fprint(w, "[]")
		return
	}
	postsList, err := json.Marshal(posts)
	if err != nil {
		fmt.Fprint(w, "[]")
		return
	}
	fmt.Fprintf(w, "%s", string(postsList))
}

func createPost(title, body string) BlogPost {
	post := BlogPost{Title: title, Body: body}
	posts = append(posts, post)
	return post
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fuck da police")
}
