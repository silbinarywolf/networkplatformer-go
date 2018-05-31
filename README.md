# Networked Platformer Proof of Concept in Golang

This is a thrown together remix of the [Ebiten platformer example](https://github.com/hajimehoshi/ebiten/tree/master/examples/platformer).
The difference is that this allows you to host a Websocket server and connect with native or web clients.

**!!!Warning!!!**

This has bugs, sub-optimal networking code and poor error handling. This was hastily put together to show that this
is possible with Go. Use this more as a loose guide to getting started!

**!!!Warning!!!**

## Install

```
go get github.com/silbinarywolf/networkplatformer-go
```

## Requirements

* Golang 1.10+
* [Gorilla Websockets](https://github.com/gorilla/websocket) (for native client and server)
* [GopherJS Websockets](https://github.com/gopherjs/websocket)  (for JavaScript client)

## How to use

These commands were all tested on Windows, running via Git Bash.

Build and run server
```
go build && ./networkplatformer-go.exe --server
```

Build and run client
```
go build && ./networkplatformer-go.exe
```

Build web client (requires GopherJS is installed)
```
GOOES=linux gopherjs build
```
Then open "index.html" in your browser of choice to run it. (Tested Chrome and Firefox)
