# 🚇 sshtunnel

Ultra simple SSH tunnelling for Go programs.

## Fork notes

This is a fork of [elliotchance/sshtunnel](https://github.com/elliotchance/sshtunnel) at v1.4.0:
- the API is backwards incompatible!
- the API will break again. 

To make this more evident, the Go module name is now `github.com/marco-m/sshtunnel`.

## Installation

```bash
go get -u github.com/marco-m/sshtunnel
```

## Example

```go
// Setup the tunnel, but do not yet start it yet.
tunnel := sshtunnel.NewSSHTunnel(
   // User and host of tunnel server, it will default to port 22
   // if not specified.
   "ec2-user@jumpbox.us-east-1.mydomain.com",

   // Pick ONE of the following authentication methods:
   sshtunnel.PrivateKeyFile("path/to/private/key.pem"), // 1. private key
   ssh.Password("password"),                            // 2. password
   sshtunnel.SSHAgent(),                                // 3. ssh-agent

   // The destination host and port of the actual server.
   "dqrsdfdssdfx.us-east-1.redshift.amazonaws.com:5439",
   
   // The local port you want to bind the remote port to.
   // Specifying "0" will lead to an ephemeral port, which can be read
   // either from listener.Addr or from tunnel.Local.Port
   "8443",
)

// You can provide a logger for debugging, or remove this line to
// make it silent.
tunnel.Log = log.New(os.Stdout, "", log.Ldate | log.Lmicroseconds)

listener, err := tunnel.Listen()
if err != nil {
    panic(err)
}
// After having called tunnel.Listen(), there is no need to sleep, the port
// is already allocated and bound. The address is available at
// listener.Addr().String(), or the port at tunnel.Local.Port.  

// Start the server in the background.
go tunnel.Serve(listener)

// You can use any normal Go code to connect to the destination server
// through localhost. You may need to use 127.0.0.1 for some libraries.
//
// Here is an example of connecting to a PostgreSQL server:
conn := fmt.Sprintf("host=127.0.0.1 port=%d username=foo", tunnel.Local.Port)
db, err := sql.Open("postgres", conn)

// ...
```
