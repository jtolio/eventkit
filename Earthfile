VERSION 0.6
FROM golang:1.21
WORKDIR /go/eventkit

lint:
    RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go install honnef.co/go/tools/cmd/staticcheck@2023.1.6
    RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
    RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go install github.com/storj/ci/check-cross-compile@latest
    COPY . .

    RUN golangci-lint run
    RUN staticcheck ./...
    RUN check-cross-compile ./...

    WORKDIR /go/eventkit/tools
    RUN golangci-lint run
    RUN staticcheck ./...
    RUN check-cross-compile ./...

    WORKDIR /go/eventkit/eventkitd
    RUN golangci-lint run
    RUN staticcheck ./...
    RUN check-cross-compile ./...

    WORKDIR /go/eventkit/eventkitd-bigquery
    RUN golangci-lint run
    RUN staticcheck ./...
    RUN check-cross-compile ./...

test:
   COPY . .
   RUN --mount=type=cache,target=/root/.cache/go-build \
       --mount=type=cache,target=/go/pkg/mod \
       go test ./...
   WORKDIR /go/eventkit/tools
   RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go test ./...
   WORKDIR /go/eventkit/eventkitd
   RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go test ./...
   WORKDIR /go/eventkit/eventkitd-bigquery
   RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go test ./...

check-format:
   COPY . .
   RUN bash -c 'mkdir build || true'
   RUN bash -c '[[ $(git status --short) == "" ]] || (echo "Before formatting, please commit all your work!!! (Formatter will format only last commit)" && exit -1)'
   RUN git show --name-only --pretty=format: | grep ".go" | xargs -n1 gofmt -s -w
   RUN git diff > build/format.patch
   SAVE ARTIFACT build/format.patch

format:
   LOCALLY
   COPY +check-format/format.patch build/format.patch
   RUN git apply --allow-empty build/format.patch
   RUN git status

build:
    WORKDIR /usr/src/eventkit
    COPY . .
    ENV CGO_ENABLED=false
    WORKDIR /usr/src/eventkit/eventkitd-bigquery
    RUN  --mount=type=cache,target=/root/.cache/go-build \
         --mount=type=cache,target=/go/pkg/mod \
         go install
    SAVE ARTIFACT /go/bin/eventkitd-bigquery

build-image:
    COPY .git .git
    ARG TAG=$(git rev-parse --short HEAD)
    ARG IMAGE=img.dev.storj.io/nightly/eventkitd
    BUILD +build-tagged-image --TAG=$TAG --IMAGE=$IMAGE

build-tagged-image:
    ARG --required TAG
    ARG --required IMAGE
    COPY +build/eventkitd-bigquery /opt/eventkit/bin/eventkitd-bigquery
    SAVE IMAGE --push $IMAGE:$TAG $IMAGE:latest
