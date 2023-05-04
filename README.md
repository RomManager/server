# server
Will be using [go.rice](https://github.com/GeertJohan/go.rice) to connect web-client and go server

Automate with npm script or something like that

`docker run --rm -p 8080:8080 ghcr.io/rommanager/server:edge`

## Backend
-- To be extended -- 

The backend is written in go with gin and gorm

`cd backend`

`go run *.go`

To run it with live reloading use [air](https://github.com/cosmtrek/air)

`air`


For the Images and game information we use the [SteamGridDB Api](https://www.steamgriddb.com/api/v2). Create a key and set it in the `.env`

`STEAMGRIDDBAPIKEY="yourapikey"`