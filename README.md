# Simple `context` package usage demo

Trying to understand how Golang `context` works, applying it to on a middleware use case.

```bash
# Get deps
$ go get -v .

# Launch server
$ go run main.go

# Assert that first call to /status endpoint will serve 'Not Logged in' as response body
curl localhost:8085/status # will ouput 'Not Logged in'

# Login and save email  in cookies.txt
curl -c cookies.txt -H 'username:julglotain@github.com' localhost:8085/login

# Now check /status again 
curl -b cookies.txt localhost:8085/status # will ouput 'Hello julglotain@github.com'
```

Sources:
* https://blog.golang.org/context
* https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/
