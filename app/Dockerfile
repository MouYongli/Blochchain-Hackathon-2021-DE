FROM golang:1.17-buster

WORKDIR /backend

COPY . /backend

RUN cd /backend \
    && go mod download \
    && go mod tidy \
    && go build -o app main.go

VOLUME [ "/backend/uploads" ]

EXPOSE 8090

CMD ["/backend/app"]