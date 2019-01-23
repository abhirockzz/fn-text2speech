FROM fnproject/go:dev as build-stage
WORKDIR /function
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/func/
RUN cd /go/src/func/ && dep ensure
RUN cd /go/src/func/ && go build -o func

FROM fnproject/go
RUN apk add --no-cache flite
WORKDIR /function
COPY --from=build-stage /go/src/func/func /function/
ENTRYPOINT ["./func"]