FROM golang:1.13.0-alpine3.10 as BUILDER

# install dependencies
RUN apk --update add \
#    postgresql-client \
    curl \
    git \
    gcc \
    libc-dev \
	wget \
	tar \
	bash \
	npm \
	yarn \
	make

COPY . /go/src/go-server

WORKDIR /go/src/go-server/src

RUN make init
RUN make test
RUN make build

RUN chmod a+x go-server


#########
####  the final image
FROM scratch 
# FROM alpine 

COPY --from=BUILDER /go/src/go-server/src/go-server /go/bin/go-server

ENTRYPOINT ["/go/bin/go-server"]