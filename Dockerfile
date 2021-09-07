FROM golang:1-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go ./
COPY ./pkg ./pkg

COPY ./makefile ./
RUN make
ENTRYPOINT ["./archive-provide"]
