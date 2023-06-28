FROM golang:alpine

WORKDIR /app

COPY ./static ./static
COPY go.mod .
#COPY go.sum .
#RUN go get
COPY main.go .
RUN go build .
RUN chmod +x whoami
CMD ["./whoami"]
