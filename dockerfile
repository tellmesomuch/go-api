FROM sundaeparty/devcontainer:latest AS builder

ENV ENV="/root/.bashrc" \
    TZ=Europe

WORKDIR /build

COPY . /build/

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /build/api-linux-amd64

FROM scratch
WORKDIR /go/bin/
COPY --from=builder /build/api-linux-amd64 /go/bin/api-linux-amd64
ENTRYPOINT ["/go/bin/api-linux-amd64"]