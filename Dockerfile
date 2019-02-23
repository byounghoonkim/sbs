FROM golang:latest
RUN mkdir /app 
ADD . /app/
WORKDIR /app 

ADD ./go.mod ./
RUN go mod download

RUN go build .
EXPOSE 2018
CMD ["./sbs", "serve"]
