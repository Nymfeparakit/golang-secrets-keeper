# golang-secrets-keeper
Keeper of user secrets, such login and password, bank cards data, secure notes.  

## Client
The project has client-server architecture. Client is a cli program, which allows user to authorize and then add secrets data to the server.
Server and client communicate via RPC protocol.  

The client program has set of commands. Here are some of them:  
`reg` - create new user on server
