# go-rest-api-service-legacy development

## Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

## Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-rest-api-service-legacy
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

## Build

1. Build the binaries.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make build

    ```

1. The binaries will be found in ${GIT_REPOSITORY_DIR}/target.
   Example:

    ```console
    tree ${GIT_REPOSITORY_DIR}/target

    ```

## Run

1. Run the binary.
   Examples:

    1. linux

        ```console
        ${GIT_REPOSITORY_DIR}/target/linux-amd64/go-rest-api-service-legacy

        ```

    1. macOS

        ```console
        ${GIT_REPOSITORY_DIR}/target/darwin-amd64/go-rest-api-service-legacy

        ```

    1. windows

        ```console
        ${GIT_REPOSITORY_DIR}/target/windows-amd64/go-rest-api-service-legacy

        ```

1. Clean up.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## Test

1. Run Go tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```

## Documentation

1. Start `godoc` documentation server.
   Example:

    ```console
     cd ${GIT_REPOSITORY_DIR}
     godoc

    ```

1. Visit [localhost:6060](http://localhost:6060)
1. Senzing documentation will be in the "Third party" section.
   `github.com` > `senzing` > `go-rest-api-service-legacy`

1. When a versioned release is published with a `v0.0.0` format tag,
the reference can be found by clicking on the following badge at the top of the README.md page:
[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-rest-api-service-legacy.svg)](https://pkg.go.dev/github.com/senzing-garage/go-rest-api-service-legacy)

## Docker

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make docker-build

    ```

1. Run docker container.
   Example:

    ```console
    docker run \
      --rm \
      senzing/go-rest-api-service-legacy

    ```

## Package

### Package RPM and DEB files

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make package

    ```

1. The results will be in the `${GIT_REPOSITORY_DIR}/target` directory.
   Example:

    ```console
    tree ${GIT_REPOSITORY_DIR}/target

    ```

### Test DEB package on Ubuntu

1. Determine if `go-rest-api-service-legacy` is installed.
   Example:

    ```console
    apt list --installed | grep go-rest-api-service-legacy

    ```

1. :pencil2: Install `go-rest-api-service-legacy`.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}/target
    sudo apt install ./go-rest-api-service-legacy-0.0.0.deb

    ```

1. Run command.
   Example:

    ```console
    go-rest-api-service-legacy

    ```

1. Remove `go-rest-api-service-legacy` from system.
   Example:

    ```console
    sudo apt-get remove go-rest-api-service-legacy

    ```
