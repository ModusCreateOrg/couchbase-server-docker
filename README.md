![No longer maintained](https://img.shields.io/badge/Maintenance-OFF-red.svg)
### [DEPRECATED] This repository is no longer maintained
> While this project is fully functional, the dependencies are no longer up to date. You are still welcome to explore, learn, and use the code provided here.
>
> Modus is dedicated to supporting the community with innovative ideas, best-practice patterns, and inspiring open source solutions. Check out the latest [Modus Labs](https://labs.moduscreate.com?utm_source=github&utm_medium=readme&utm_campaign=deprecated) projects.

[![Modus Labs](https://res.cloudinary.com/modus-labs/image/upload/h_80/v1531492623/labs/logo-black.png)](https://labs.moduscreate.com?utm_source=github&utm_medium=readme&utm_campaign=deprecated)

---
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

