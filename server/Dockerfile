
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

# ! Copy ENV to /app
COPY /config/.env .env

# ! Eport ENV file ***(Cloud Compute Without .env)
RUN export $(cat .env | xargs) && go get && go build -o main .

FROM alpine

COPY --from=builder /app/main /src/main

COPY --from=builder /app/config /config

EXPOSE 8443

CMD /src/main
