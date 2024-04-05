FROM golang:1.22-alpine as builder
RUN apk --update add build-base

WORKDIR /src/app
ADD go.mod .
RUN go mod download

ADD . .
RUN go build -o bin/db ./cmd/db
RUN go build -o bin/cronjobs ./cmd/cronjobs
RUN go run ./cmd/build

FROM alpine
RUN apk add --no-cache tzdata ca-certificates
RUN apk add --no-cache dcron

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/app/bin/app .
COPY --from=builder /src/app/bin/db .
COPY --from=builder /src/app/bin/cronjobs .

RUN chmod +x /bin/app
RUN chmod +x /bin/db
RUN chmod +x /bin/cronjobs

# Copying crontab file
COPY crontab /etc/cron.d/cronjobs
RUN chmod 644 /etc/cron.d/cronjobs

# Running the app and cronjobs
CMD crond -f -L /dev/null && /bin/db migrate && /bin/app