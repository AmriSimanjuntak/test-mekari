FROM golang:alpine as builder

# Install git + SSL ca certificates
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /opt/test-mekari/

# use modules
COPY go.mod .


# RUN go mod download
RUN go mod verify

COPY . .
#COPY config.production.yml config.yml

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/test-mekari .


############################
# STEP 2 build a small image
############################
FROM golang:alpine

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/test-mekari .


COPY --from=builder /opt/test-mekari/docker-compose.yml .

ENV TZ=Asia/Jakarta

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8080
# Run binary.
ENTRYPOINT ["./test-mekari"]
