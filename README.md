# color-info

A simple golang server that responds with the RGB equivalent of hex color strings.

## Installation

Installation is as easy as cloning this repo locally. It doesn't even have to be in your GOPATH because this project uses Go Modules. This does mean that you have to have Go 1.12+ installed on your machine before using this project.

## Building

Running a simple `make build` will build and mark as executable a binary that will show up in the `bin` folder. It can then be executed by running `./bin/color-info`.

## Running

There are a couple of ways to run this project. You can run a `make run` to run it using just `go run` or you can build your own binary using `make build` and then run it using the command above _or_ you can run `make run-bin`. This command will build your binary and run that.

## API listing

This application is really simple. It has 3 endpoints, `GET /`, `GET /ping` and `POST /color`.

### `GET /`

This endpoint just responds back with the server ID and the status of the server.

### `GET /ping`

This endpoint responds back with text `pong`, mostly meant as a health check but it's fun to ping/pong the server.

### `POST /color`

This endpoint is more interesting. When POSTed to with a request body like

```json
{ "color": "#000000" }
```

It will return the same color in the RGB space. For the above request,

```json
{
  "rgbColor": {
    "R": 0,
    "G": 0,
    "B": 0
  },
  "rgbColorString": "rgb(0, 0, 0)"
}
```
