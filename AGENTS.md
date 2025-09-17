# Booru Server Agent Guide

## Project Overview

This project is a Go server that acts as an aggregator for multiple booru-style image APIs. It provides a single, unified search endpoint to query across all configured modules. The server is designed to be extensible, allowing for new modules to be added easily.

The project is configured with a full CI/CD pipeline for automated builds, testing, containerization, and releases.

## Development

### Building the Application

To build the application locally, use the `make build` command. This will produce a binary named `booru-server` in the root directory.

```bash
make build
```

The build process automatically embeds version information into the binary.

### Running the Application

To run the application after building it, use the `make run` command.

```bash
make run
```

The server will start on the host and port specified in `config.yaml` (default: `127.0.0.1:8080`).

### Running Tests

To run the test suite, use the `make test` command. This will run all unit tests in the project.

```bash
make test
```

## CI/CD Pipeline

The project has a comprehensive CI/CD setup using GitHub Actions.

### Docker Builds

On every push to the `main` branch, a Docker image is automatically built and pushed to the GitHub Container Registry (GHCR). The image is tagged with `latest` and the commit SHA.

### Automated Releases

This project uses `semantic-release` to automate the release process. When commits are pushed to the `main` branch, the following happens automatically:

1.  Commit messages are analyzed to determine the next semantic version number (following the [Conventional Commits](https://www.conventionalcommits.org/) specification).
2.  Binaries are cross-compiled for Linux, Windows, and macOS (amd64 and arm64).
3.  A new Git tag is created for the version.
4.  A GitHub Release is created with automatically generated release notes.
5.  The compiled binaries are uploaded as assets to the GitHub Release.

To trigger a release, push a commit to `main` with a message like `feat: Add new feature` or `fix: Correct a bug`.
