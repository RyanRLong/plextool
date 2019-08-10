package main

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/streadway/amqp"
)

// holds a request type
type status struct {
	request *http.Request
}

// getJSON returns the json body of the request as []byte.
// if the request in status is multipart, only the first part is
// processed.
func (status status) getJSON() []byte {
	params, err := status.getParams()
	if err != nil {
		log.Panic(err)
	}
	mr := multipart.NewReader(status.request.Body, params["boundary"])
	p, err := mr.NextPart()
	if err == io.EOF {
		return []byte{}
	}
	if err != nil {
		log.Fatal(err)
	}
	json, err := ioutil.ReadAll(p)
	if err != nil {
		log.Fatal(err)
	}
	return json
}

// getParams returns all the header paramteres as map[string]string
func (status status) getParams() (map[string]string, error) {
	_, params, err := mime.ParseMediaType(status.request.Header.Get("Content-Type"))
	return params, err
}

// sendMessage sends a message to the rabbitmq queue or exchange
func sendMessage(message []byte) {
	exchangeActivityURL := os.Getenv("EXCHANGE_ACTIVITY_URL")
	if exchangeActivityURL == "" {
		log.Panic("Environment variable \"EXCHANGE_ACTIVITY_URL\" not found.")
		return
	}
	conn, err := amqp.Dial(exchangeActivityURL)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicln(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"activity", // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Panicln(err)
	}
	err = ch.Publish(
		"activity", // exchange
		"",         // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
	if err != nil {
		log.Panicln(err)
	}
}

// requestHandler handles routing requests
func requestHandler(w http.ResponseWriter, r *http.Request) {
	status := status{r}
	json := status.getJSON()
	sendMessage(json)
}

// listenForRequests listens for requests
func listenForRequests() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// main execution
func main() {
	listenForRequests()
}
