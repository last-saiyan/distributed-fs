FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go mod download

CMD [ "go", "run", "./dataNode/start/start.go"]

EXPOSE 8010/tcp