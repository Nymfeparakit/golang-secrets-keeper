# golang-secrets-keeper
Keeper of user secrets, such login and password, bank cards data, secure notes.  

## Client
The project has client-server architecture. Client is a cli program, which allows user to authorize and then add secrets data to the server.
Server and client communicate via RPC protocol.  

The client program has set of commands. The following commands are for user authentication/authorization:  
- `reg` - create new user on server  
- `login` - login existing user  

![Alt text](login_screen.png?raw=true "Login page")

There are also subset of commands to perform crud operations with user's secrets, such as `add`, `list`, `update`, `delete`.  
#### todo: add commands examples  
Complete information about each command and its flags can be obtained using the help command, for example:
```
keeper_client add --help
```

## Server
### Setup
The server requires a postgres database to work. After creating database, run migrations in it using goose:
```
cd server/internal/migrations
goose postgres "user={db_user} password={db_password} dbname={db_name} sslmode=disable" up
```
To run server, build the binary and then run the it. When starting the server, the database URL should be specified using `-d` flag or environment variable `DATABASE_DSN`
```
go build -o keeper_server server/cmd/gophkeeper/main.go
./keeper_server -d "postgresql://{db_user}:{db_password}@localhost:5432/{db_name}"
```
The server has several settings, the values of which can be set using flags or environment variables. For, example, 
