FROM golang:1.24-alpine AS build

RUN apk add --no-cache git build-base && \
  git clone --branch v0.144.0  https://github.com/gohugoio/hugo.git && \
  cd hugo/ && \
  CGO_ENABLED=0 go build -ldflags "-s -w -X github.com/gohugoio/hugo/common/hugo.buildDate=$(date +%Y-%m-%dT%H:%M:%SZ) -X github.com/gohugoio/hugo/common/hugo.commitHash=$(git rev-parse --short HEAD)" -o /tmp/hugo . && \
  CGO_ENABLED=1 go build -tags extended -ldflags "-s -w -X github.com/gohugoio/hugo/common/hugo.buildDate=$(date +%Y-%m-%dT%H:%M:%SZ) -X github.com/gohugoio/hugo/common/hugo.commitHash=$(git rev-parse --short HEAD)" -o /tmp/hugo-extended  .

COPY . woodpecker-hugo
RUN cd woodpecker-hugo && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -v -a -tags netgo -o /tmp/woodpecker-hugo

FROM plugins/base:linux-amd64

LABEL maintainer="Florian Maurer <fmaurer@disroot.org>" \
  org.label-schema.name="Woodpecker Hugo" \
  org.label-schema.vendor="Community" \
  org.label-schema.schema-version="1.1"

RUN apk --no-cache add git libc6-compat libstdc++

ENV PLUGIN_HUGO_ARCH=64bit
ENV PLUGIN_HUGO_SHIPPED_VERSION=0.144.0

COPY --from=build /tmp/hugo /bin/hugo
COPY --from=build /tmp/hugo-extended /bin/hugo-extended
COPY --from=build /tmp/woodpecker-hugo /bin/woodpecker-hugo

ENTRYPOINT ["/bin/woodpecker-hugo"]
