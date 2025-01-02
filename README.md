<img src="https://cdn.pnck.nl/imagevault/imagevault_logo.jpeg" width="250" alt=""/>

**Release:**

[![Release Version](https://img.shields.io/github/v/release/pnck-projects/imagevault?label=imagevault)](https://github.com/pnck-projects/imagevault/releases/latest)

**Last build:**

![Last build](https://github.com/pnck-projects/imagevault/actions/workflows/go-build.yml/badge.svg)

**Last publish:**

![Last publish](https://github.com/pnck-projects/imagevault/actions/workflows/docker-publish.yml/badge.svg)

# imagevault
Secure Image Vault

# Getting started
ImageVault can be run as a stand-alone application on a server, or run as a container in Docker or Kubernetes.

Make sure you always run the latest release version.

The entire application is built stateless and supports multiple replicas for load balancing and high-availability purposes.

# Generate documentation for local development
```shell
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```
