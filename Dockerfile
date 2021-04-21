FROM golang:1.12-alpine
LABEL maintainer="jpmenezes@gmail.com" \
    description="This image contains a golang image for development purposes."

RUN mkdir -p /jpmenezes.com/idebo
WORKDIR /jpmenezes.com/idebo

EXPOSE 40000 8888

RUN apk add --no-cache git mercurial gcc musl-dev
RUN go get github.com/derekparker/delve/cmd/dlv

RUN go mod init jpmenezes.com/idebo

RUN export GO111MODULE=on && \
go get -u goa.design/goa/v3 && \
go get -u goa.design/goa/v3/...

RUN go get -v github.com/jinzhu/gorm
RUN go get -v github.com/lib/pq

COPY . .

# These were giving trouble on startup. Without them, modules are downloaded on startup but it works
# RUN goa gen jpmenezes.com/idebo/design
# RUN goa example jpmenezes.com/idebo/design

# Go running without delve
CMD [ "go", "run", "jpmenezes.com/idebo/cmd/bo" ]

# Debugging with delve
# CMD [ "dlv", "debug", "jpmenezes.com/idebo/cmd/bo", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]
