FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache git make

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN make build



FROM alpine AS runner

COPY --from=builder /usr/local/src/.bin/api /bin/api

CMD [ "/bin/api" ]
