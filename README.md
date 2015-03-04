couchbase-server-docker
=======================

This image is currenlty not auto-built, so you will need to clone the repo
to use the image.

To build the docker image, move into the directory of the version you would like
to build and issue the docker build command. For example to build 3.0.1:

```bash
$ cd 3.0.1
$ docker build -t moduscreate/couchbase:3.0.1 .
```

To create a container.

```bash
$ docker run -d -p 8091:8091 --name couchbase moduscreate/couchbase:3.0.1
```

This will expose the web interface on the docker host, this isn't necessary, but
helpful during development.

To link to another container, modify your application code to use environment
variables, by default docker creates environment variables in your container
with the name of the linked container and their port number. So if you linked
this couchbase container with your application like:

```bash
$ docker run -d -p 8080:8080 --name my-app --link couchbase:couchbase
moduscreate/best-app-ever:latest
```

The ports for couchbase would be exposed to your app as:

```bash
COUCHBASE_PORT_8091_TCP_PORT
COUCHBASE_PORT_8091_TCP_ADDR
```

You can use this to construct your connection string.

