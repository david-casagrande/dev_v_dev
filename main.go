package main

import "fmt"
import "code.google.com/p/go-uuid/uuid"
import "net/http"
import "encoding/json"
import (
	_ "github.com/lib/pq"
	"database/sql"
)
type Game struct {
    Id    string `json:"id"`
    Title string `json:"title"`
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fn(w, r)
    }
}

func createGame(w http.ResponseWriter, r *http.Request) {
    g := Game{ Id: uuid.New(), Title: "New Game" }
    j, _ := json.Marshal(g)
    w.Write(j)
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("sup"))
}

func main() {
  db, err := sql.Open("postgres", "user=postgres dbname=ceremony_develop sslmode=disable")
	if err != nil {
		//log.Fatal(err)
	}
	rows, _ := db.Query("SELECT * FROM users")
  for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil {
        //log.Fatal(err)
    }
    fmt.Printf("%s", name)
  }


  http.HandleFunc("/create_game", addDefaultHeaders(createGame))
  http.HandleFunc("/", home)
  http.ListenAndServe(":8080", nil)
}
