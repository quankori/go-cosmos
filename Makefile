block:
	- go run cmd/block/main.go
event:
	- go run cmd/events/main.go
start:
	- archwayd start --x-crisis-skip-assert-invariants
build:
	- DOCKER_BUILDKIT=0 docker-compose build --no-cache && docker-compose up