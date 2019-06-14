FROM golang:1.12 as builder
WORKDIR /go/src/go-color
COPY main.go .
RUN go build

FROM debian:stretch-slim
WORKDIR /app
RUN groupadd -g 1000 appuser && \
	  useradd -ms /bin/bash -u 1000 -g 1000 appuser && \
      chown -R appuser:appuser /app
USER appuser
COPY --from=builder /go/src/go-color/go-color ./
ENV LISTEN_PORT=8080
EXPOSE 8080
CMD ["./go-color"]