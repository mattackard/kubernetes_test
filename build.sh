CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ip-echo main.go
mv ip-echo ./bin/ip-echo