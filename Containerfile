FROM golang:alpine AS builder

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

# Copy files
WORKDIR /app
ADD . .
# Buiild
RUN go mod tidy
RUN go build -o /app/memesaviour main.go
# RUN CGO_ENABLED=1 GOOS=linux go build -o /app/memesaviour -installsuffix 'static' -a -ldflags '-s -w' cmd/main.go

FROM scratch
LABEL maintainer="contre95"
# USER nonroot:nonroot
# COPY --from=builder --chown=nonroot:nonroot /app/memesaviour /app/memesaviour
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/memesaviour /app/memesaviour
ENTRYPOINT ["/app/memesaviour"]

