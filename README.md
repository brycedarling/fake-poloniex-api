# Fake Poloniex API

Simulates the [Poloniex real-time API](https://poloniex.com/support/api/).

Creates a local [WAMP (Web Application Messaging Protocol)](http://wamp-proto.org/) server for communicating via websockets between clients.

Creates a local client for publishing cryptocurrency tick data that currently uses a small ticks.json file from actual ticker data and only publishes to the "ticker" topic.


## Getting Started

These instructions will help you get a fake Poloniex API running on your local machine.


### Prerequisites

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Node.js](https://nodejs.org/)
* [Yarn](https://yarnpkg.com/)

On my Mac, I had to first install the Command Line Tools:

1. Launch the `Terminal` application, found in `/Applications/Utilities/`
2. Run the following: `xcode-select --install`
3. You'll be waiting a hot second for this to complete so go get a coffee, make a sandwich, or take a nap, and by then it'll probably be done and you'll be ready to move on to the next step.

After installing the Command Line Tools, then I was able to install [Homebrew](https://brew.sh/) by running the following in Terminal:

```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Once that's done, you're finally ready to install Go, Node.js, and Yarn from Terminal by running the following:

```
brew install go node yarn
```


### Installation

In order to build the Go server program, run:

```
make build
```

That should output a `fake-poloniex-api` program in the same directory.

In order to run the client, you'll need to install it's dependencies:

```
yarn
```

That should install the dependencies in to `node_modules` folder in the same directory.


## Running the fake Poloniex API

Now that you have all the prequisites, you can start the WAMP server by running:

```
make run
```

You should see the server startup:

    $ make run
    ./fake-poloniex-api
    2017/09/01 11:27:56 websocket_server.go:60: NewBasicWebsocketServer
    2017/09/01 11:27:56 websocket_server.go:47: NewWebsocketServer
    2017/09/01 11:27:56 realm.go:43: Established internal session: 4354497965326435
    2017/09/01 11:27:56 router.go:91: registered realm: realm1
    2017/09/01 11:27:56 websocket_server.go:78: RegisterProtocol: wamp.2.json
    2017/09/01 11:27:56 websocket_server.go:78: RegisterProtocol: wamp.2.msgpack
    2017/09/01 11:27:56 turnpike server starting on port 8000
    2017/09/01 11:27:56 realm.go:150: [4354497965326435] PUBLISH: &{Request:7554373065801303 Options:map[] Topic:wamp.session.on_join Arguments:[map[]] ArgumentsKw:map[]}


## Publishing fake tick data

Now that the server is running, you can start publishing ticks to all connected clients by running:

```
yarn start
```

You should see the Node program start running and tell you it is publishing ticks:

    $ yarn start
    yarn start v0.27.5
    $ node fake-poloniex-ticks.js
    Ticks: 647
    Websocket connect opened!
    Publishing to 'ticker' topic, 646 ticks remaining
    Publishing to 'ticker' topic, 645 ticks remaining
    ...


## Built with

Standing on the shoulders of giants, this project was made easy thanks to a couple libraries:

* [turnpike](https://github.com/jcelliott/turnpike) - on the Go server for WAMP support.
* [autobahn-js](https://github.com/crossbario/autobahn-js) - on the JavaScript client exactly like you would to connect to Poloniex.


## TODO

* Randomly generate the ticks using a valid range of possible values.
* Add Dockerfile and compose to make installation and running quicker and easier.


## Authors

* [Bryce Darling](http://github.com/brycedarling)


## License

This project is licensed under the [WTFPL](http://www.wtfpl.net/) - see the [LICENSE](http://github.com/brycedarling/fake-poloniex-api/LICENSE) file for details.


## Acknowledgements

* Hi mom

