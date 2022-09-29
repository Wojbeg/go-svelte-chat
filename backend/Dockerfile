FROM golang:alpine

RUN mkdir /backend
COPY . /backend/
WORKDIR /backend

RUN go build -o main .

EXPOSE 5000
CMD ["/backend/main"]