FROM golang:1.15 as builder

WORKDIR /usr/src/diodon
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/diodon


FROM alpine

WORKDIR /bin/diodon

COPY --from=builder /bin/diodon .

CMD ["./diodon"]
