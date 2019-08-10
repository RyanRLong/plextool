# plextool
[![CircleCI](https://circleci.com/gh/SaltyCatFish/plextool/tree/master.svg?style=svg)](https://circleci.com/gh/SaltyCatFish/plextool/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/saltycatfish/plextool)](https://goreportcard.com/report/github.com/saltycatfish/plextool)

plextool is a package containing a sender and receiver for interacting with a rabbitMQ database. It listens for webhooks
from a Plex Media server with the sender.

On receipt, it sends the JSON received to a RabbitMQ exchange.

The receiver subscribes to the exchange and processes any JSONs on the exchange and display a Windows 10 toast message.

Any time media is played/paused/stopped on a Plex Media Server, an "event" can be emitted using the webook feature (Pro only at the time of this).

This event contains loads of data, including the current media being played statistics down to other media the user might enjoy.

## Building

Build the sender and receiver the same you would any Go code.

```bash
go build cmd/receiver/main.go
go build cmd/sender/main.go
```

## Usage

Soon...

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
