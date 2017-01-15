# drone-spark

[![Build Status](https://travis-ci.org/hhxiao/drone-spark.svg?branch=master)](https://travis-ci.org/hhxiao/drone-spark)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-template?status.svg)](http://godoc.org/github.com/hhxiao/drone-spark)
[![Go Report](https://goreportcard.com/badge/github.com/hhxiao/drone-spark)](https://goreportcard.com/report/github.com/hhxiao/drone-spark)
[![MicroBadger](https://images.microbadger.com/badges/image/hhxiao/drone-spark.svg)](https://microbadger.com/images/hhxiao/drone-spark "Get your own image badge on microbadger.com")

Drone plugin for sending Spark notifications. For the usage information and a
listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build -t hhxiao/drone-spark .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-spark' not found or does not exist..
```

## Usage

Execute from the working directory:

Send notification to individual

```
docker run --rm \
  -e DRONE_REPO_OWNER=hhxiao \
  -e DRONE_REPO_NAME=drone_spark \
  -e DRONE_REPO_LINK=https://github.com/hhxiao/drone_spark \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_REF= refs/heads/master \
  -e DRONE_COMMIT_AUTHOR=hhxiao \
  -e DRONE_COMMIT_AUTHOR_EMAIL=hhxiao@gmail.com \
  -e DRONE_COMMIT_MESSAGE="bug fixing" \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://beta.drone.io/hhxiao/drone-spark/1 \
  -e DRONE_TAG=1.0.0 \
  -e PLUGIN_AUTH_TOKEN=################################################ \
  -e PLUGIN_PERSON_EMAIL="hhxiao@gmail.com" \
  -e PLUGIN_ATTACHMENT=README.md \
  hhxiao/drone-spark
```

Send notification to spark room

```
docker run --rm \
  -e DRONE_REPO_OWNER=hhxiao \
  -e DRONE_REPO_NAME=drone_spark \
  -e DRONE_REPO_LINK=https://github.com/hhxiao/drone_spark \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_REF= refs/heads/master \
  -e DRONE_COMMIT_AUTHOR=hhxiao \
  -e DRONE_COMMIT_AUTHOR_EMAIL=hhxiao@gmail.com \
  -e DRONE_COMMIT_MESSAGE="bug fixing" \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://beta.drone.io/hhxiao/drone-spark/1 \
  -e DRONE_TAG=1.0.0 \
  -e PLUGIN_AUTH_TOKEN=################################################ \
  -e PLUGIN_ROOM_NAME="Build Status Room" \
  -e PLUGIN_ATTACHMENT=README.md \
  hhxiao/drone-spark
```

## Reference
This plugin references a lot from the official **[drone-slack](https://github.com/drone-plugins/drone-slack)** plugin
