# golang-secrets-keeper
Keeper of user secrets, such login and password, bank cards data, secure notes.  

## Secrets
There are 4 types of secrets that the user can manage:
 - login and password pair
 - bank card info, such as card number, cvv code, expiration month, etc
 - secure text note
 - secure binary data uploaded from some file

## Client
The project has client-server architecture. Client is a cli program, which allows user to authorize and then add secrets data to the server.
Server and client communicate via RPC protocol.  

The client program has set of commands. The following commands are for user authentication/authorization:  
- `reg` - create new user on server  
- `login` - login existing user  

For user convenience, the client provides simple tui using [tview](https://github.com/rivo/tview).  

![Alt text](login_screen.png?raw=true "Login page")

There are also subset of commands to perform crud operations with user's secrets, such as `add`, `list`, `update`, `delete`.  
For example, to update existing secret with id="123":  
```
./keeper_client update -t card_info 123
```

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
The server has also some other settings, the values of which can be set using flags or environment variables. The full info about settings can be obtained with `help` command.

## Encryption and security
Client and server can use tls for more secure communication.
To run server and client with tls support:
```
./keeper_server -d "postgresql://{db_user}:{db_password}@localhost:5432/{db_name}" -s -c {path/to/tls/certificate} -k {path/to/tls/key}
./keeper_client -s -c {path/to/tls/certificate}
```
Additionally to this, client sends secrets data encrypted with AES algorithm.

## Local storage
User can manage his secrets without connection to server, all his secrets are stored in temporarily local storage in sqlite database.
