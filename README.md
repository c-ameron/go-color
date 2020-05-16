# go-color

A simple go app to convert a hex color to RGB

```bash
$ curl localhost:8080/convert?hex=ff0000
RGB(255, 0, 0)
```

Any query which isn't the example format will result in a `HTTP 400` error
```bash
$ curl localhost:8080/convert?bad=query -i
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Fri, 14 Jun 2019 09:09:05 GMT
Content-Length: 50

Bad query, please use format: /convert?hex=ff0000
```


There's also a `/status` endpoint

``` bash
$ curl localhost:8080/status -i
HTTP/1.1 200 OK
Date: Fri, 14 Jun 2019 08:55:58 GMT
Content-Length: 3
Content-Type: text/plain; charset=utf-8

OK
```

## How To Build & Run 

### Golang

```bash
go build
./go-color
```

### Docker

```bash
docker build --pull --force-rm -t go-color .
docker run -p 8080:8080 go-color
```

### Configuration

To change the listening port, set the `LISTEN_PORT` environment variable. The default is `8080`
```bash
LISTEN_PORT=9000 ./go-color
```
