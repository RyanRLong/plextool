package plextool

import (
	"encoding/json"
	"log"
)

// Listen listens for incoming exchange data and
// displays a Windows 10 toast message on receipt.
func Listen() {
	receiver, err := NewReceiver()
	if err != nil {
		log.Panic(err.Error())
	}
	defer receiver.Channel.Close()
	defer receiver.Connection.Close()

	msgs := receiver.Listen()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var data IncomingData
			log.Printf("%v", string(d.Body))
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Panic(err)
			}
			film := data.GetFilmData()
			DisplayToast(film, data)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
