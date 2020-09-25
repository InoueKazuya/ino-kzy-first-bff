FROM golang:1.15-alpine
ADD . /go/src/on-boarding-bff
RUN go install on-boarding-bff

FROM alpine:latest
COPY --from=0 /go/bin/on-boarding-bff .
ENV PORT 8443
CMD ["./on-boarding-bff"]