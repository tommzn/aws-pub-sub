FROM gcr.io/distroless/static:nonroot

WORKDIR /go

COPY build_artifact_bin lambdahandler

RUN chmod 755 /go/lambdahandler
ENTRYPOINT ["/go/lambdahandler"]
