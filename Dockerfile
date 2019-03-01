FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app 

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch
COPY --from=builder /app/sbs /app/

EXPOSE 2018
CMD ["serve", "--host", "0.0.0.0"]
ENTRYPOINT ["/app/sbs"]
