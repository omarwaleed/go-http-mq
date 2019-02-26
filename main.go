package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

type Queue struct {
	name    string
	entries [][]byte
	// Consumers []string
}

var queues = []Queue{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/q/{queue}", HandleQueueEntry).Methods("GET", "POST")

	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}

func HandleQueueEntry(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

	case "POST":
	default:
		log.Fatalln("Somehow method is not GET or POST")
	}
}

func getQueueEntry(queueName string) (error, []byte) {
	if len(queues) == 0 {
		return emptyQueueError(), nil
	}
	for index, q := range queues {
		if q.name == queueName {
			toReturn := q.entries[0]
			queues[index].entries = queues[index].entries[1:]
			return nil, toReturn
		}
	}
	return noQueueFoundError(), nil
}

func emptyQueueError() error {
	return errors.New("Queue is empty")
}

func noQueueFoundError() error {
	return errors.New("Queue with given name doesn't exist")
}
