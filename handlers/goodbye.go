package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Hello from goodbye handler")
	w.Write([]byte("sup brah"))
	// d, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	// fmt.Fprintf(w, "Goodbye %s", d)
}
