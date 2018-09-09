
TLS="--tls"
SERVER_OVERRIDE="--server_host_override sbs.test.youtube.com"

go run main.go serve $TLS &
SERVER_PID=$!

while ! nc -z localhost 2018 ; do sleep 1 ; done

fallocate -l 5M testfile

go run main.go put $TLS $SERVER_OVERRIDE testfile
go run main.go get $TLS $SERVER_OVERRIDE testfile

rm testfile
pkill -P $SERVER_PID
