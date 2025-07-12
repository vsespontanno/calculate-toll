gate:
	@go build -o bin/gate gateway/main.go
	@./bin/gate


obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

rec:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calc:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

agg:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

prom:
	docker run --name prometheus -d -p 127.0.0.1:9090:9090 -v $(pwd)/.config/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

graf:
	docker run -d -p 3000:3000 --name=grafana grafana/grafana-enterprise
.PHONY: obu