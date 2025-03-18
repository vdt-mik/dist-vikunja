FROM --platform=$BUILDPLATFORM ghcr.io/techknowlogick/xgo:go-1.23.x AS builder

RUN go install github.com/magefile/mage@latest && \
    mv /go/bin/mage /usr/local/go/bin

WORKDIR /go/src/code.vikunja.io/api

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ARG RELEASE_VERSION=dev
ENV RELEASE_VERSION=$RELEASE_VERSION

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p frontend/dist

COPY main.go magefile.go ./
COPY pkg/ ./pkg/
COPY frontend/embed.go ./frontend/
COPY rest/ ./rest/

RUN echo "<!-- Placeholder -->" > frontend/dist/index.html

ENV GOPROXY=https://goproxy.kolaente.de

RUN export PATH=$PATH:$GOPATH/bin && \
    mage build:clean && \
    mage release:xgo "${TARGETOS}/${TARGETARCH}/${TARGETVARIANT}"

FROM scratch AS app
LABEL org.opencontainers.image.authors='maintainers@vikunja.io'
LABEL org.opencontainers.image.url='https://vikunja.io'
LABEL org.opencontainers.image.documentation='https://vikunja.io/docs'
LABEL org.opencontainers.image.source='https://code.vikunja.io/vikunja'
LABEL org.opencontainers.image.licenses='AGPLv3'
LABEL org.opencontainers.image.title='Vikunja API'

WORKDIR /app/vikunja
ENTRYPOINT [ "/app/vikunja/vikunja" ]
EXPOSE 3456
USER 1000

ENV VIKUNJA_SERVICE_ROOTPATH=/app/vikunja/

COPY --from=builder /build/vikunja-* vikunja
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ 