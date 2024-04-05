FROM golang:1.22-alpine as builder
RUN apk --update add build-base

WORKDIR /src/app
ADD go.mod .
RUN go mod download

ADD . .
RUN go build -o bin/db ./cmd/db
RUN go run ./cmd/build

FROM alpine
RUN apk add --no-cache tzdata ca-certificates
WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/app/bin/app .
COPY --from=builder /src/app/bin/db .

RUN chmod +x /bin/app
RUN chmod +x /bin/db

# Running the app
CMD /bin/db migrate; /bin/app;
