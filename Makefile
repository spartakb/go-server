default:
	@echo must specify rule

build:
	docker build -t go-server:local .

run:
	docker run \
		--rm \
		--name go-server \
		-p 80:80 \
		go-server:local 
