package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  
)
type Post struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
}

var posts []Post


func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(posts)
}



func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var post Post
  _ = json.NewDecoder(r.Body).Decode(&post)
  posts = append(posts, post)
  json.NewEncoder(w).Encode(&post)
}



func main() {
  router := mux.NewRouter()
  posts = append(posts, Post{ID: "1", Title: "My first post", Body:      "This is the content of my first post"})
  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  
http.ListenAndServe(":8000", router)
}