FROM golang:latest

WORKDIR /goaway

COPY src/go.mod src/go.sum ./src/

RUN cd ./src && go mod download

COPY src ./src

CMD ["go", "run", "./src"]
