package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

const queue = "queue-example"

func main() {
	conn, err := amqp.Dial("amqp://rabbitmq")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		ch, err := conn.Channel()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer ch.Close()

		err = ch.Publish(
			"",
			queue,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Go Content"),
			},
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Success")
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
