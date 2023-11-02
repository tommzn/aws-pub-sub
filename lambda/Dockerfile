FROM gcr.io/distroless/static:nonroot

WORKDIR /go

COPY build_artifact_bin lambdahandler

ENTRYPOINT ["/go/lambdahandler"]
