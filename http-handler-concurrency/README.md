https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/

ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
