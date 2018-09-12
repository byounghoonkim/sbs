
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    fallocate -l 5M testfile
elif [[ "$OSTYPE" == "darwin"* ]]; then
    mkfile -n 5M testfile
else
    echo "unknwon os type. SKIP this test."
    exit
fi


TLS="--tls"
SERVER_OVERRIDE="--server_host_override sbs.test.youtube.com"

go run main.go serve $TLS &
SERVER_PID=$!

while ! nc -z localhost 2018 ; do sleep 1 ; done

go run main.go put $TLS $SERVER_OVERRIDE testfile
go run main.go get $TLS $SERVER_OVERRIDE testfile

rm testfile
pkill -P $SERVER_PID
