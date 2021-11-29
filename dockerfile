FROM golang:1.17-alphine3.14

WORKDIR /ecofriends

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o mainfile

EXPOSE 8000

CMD ["mainfile"]