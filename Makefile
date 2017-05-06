build : bin/buildlog

bin/buildlog :
	go build -o ./bin/buildlog ./cmd/buildlog

bin/buildlog.static :
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/buildlog.static ./cmd/buildlog

docker : bin/buildlog.static
	docker build -t fajran/buildlog .

test :
	go test ./pkg/... ./cmd/...

get-deps :
	go get github.com/gorilla/mux
	go get github.com/lib/pq
	go get github.com/mattes/migrate
	go get gopkg.in/yaml.v2

