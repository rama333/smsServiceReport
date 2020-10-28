# Initial stage: download modules
FROM golang:1.15 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Intermediate stage: Build the binary
FROM golang:1.15 as builder

COPY --from=modules /go/pkg /go/pkg

# add a non-privileged user
RUN useradd -u 10001 smsService


RUN mkdir -p /smsService
ADD . /smsService



WORKDIR /smsService



# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/smsService ./cmd/smsReport


# Final stage: Run the binary
FROM scratch

# don't forget /etc/passwd from previous stage

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Berlin

USER smsService

# and finally the binary
COPY --from=builder /smsService/bin/smsService /smsService

CMD ["/smsService"]