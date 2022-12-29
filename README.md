# RomManager
## Backend
The backend is written in go with gin and gorm

`cd backend`

`go run main.go`

To run it with live reloading use [air](https://github.com/cosmtrek/air)

`air`


For the Images and game information we use the [SteamGridDB Api](https://www.steamgriddb.com/api/v2). Create a key and set it in the `.env`

`STEAMGRIDDBAPIKEY="yourapikey"`