# servegrpc

## :warning: WARNING: servegrpc is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`servegrpc` is a command in the
[senzing-tools](https://github.com/Senzing/senzing-tools)
suite of tools.
This command is a
[gRPC](https://grpc.io/)
server application that supports requests to the Senzing SDK via network access.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/servegrpc.svg)](https://pkg.go.dev/github.com/senzing/servegrpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/servegrpc)](https://goreportcard.com/report/github.com/senzing/servegrpc)

## Overview

`servegrpc` supports the
[Senzing Protocol Buffer definitions](https://github.com/Senzing/g2-sdk-proto).
Under the covers, the gRPC request is translated by the gRPC server into a Senzing Go SDK API call using
[senzing/g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base).
The response from the Senzing Go SDK API is returned to the gRPC client.

Senzing SDKs for accessing the gRPC server:

1. Go: [g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc)
1. Python: [g2-sdk-python-grpc](https://github.com/Senzing/g2-sdk-python-grpc)

## Install

1. The `servegrpc` command is installed with the
   [senzing-tools](https://github.com/Senzing/senzing-tools)
   suite of tools.
   See senzing-tools [install](https://github.com/Senzing/senzing-tools#install).

## Use

```console
export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
senzing-tools servegrpc [flags]
```

1. For options and flags:
    1. [Online documentation](https://hub.senzing.com/senzing-tools/senzing-tools_servegrpc.html)
    1. Runtime documentation:

        ```console
        export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
        senzing-tools servegrpc --help
        ```

1. In addition to the following simple usage examples, there are additional [Examples](docs/examples.md).

### Using command line options

1. :pencil2: Specify database using command line option.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools servegrpc --database-url postgresql://username:password@postgres.example.com:5432/G2
    ```

1. See [Parameters](#parameters) for additional parameters.

### Using environment variables

1. :pencil2: Specify database using environment variable.
   Example:

    ```console
    export SENZING_TOOLS_DATABASE_URL=postgresql://username:password@postgres.example.com:5432/G2
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools servegrpc
    ```

1. See [Parameters](#parameters) for additional parameters.

### Using Docker

This usage shows how to initialze a database with a Docker container.

1. This usage has an SQLite database that is baked into the Docker container.
   The container is mutable and the data in the database is lost when the container is terminated.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db \
        --interactive \
        --publish 8258:8258 \
        --rm \
        --tty \
        senzing/senzing-tools servegrpc

    ```

1. This usage specifies a URL of an external database.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_DATABASE_URL=postgresql://username:password@postgres.example.com:5432/G2 \
        --interactive \
        --publish 8258:8258 \
        --rm \
        --tty \
    senzing/senzing-tools servegrpc

    ```

1. See [Parameters](#parameters) for additional parameters.

### Parameters

- **[SENZING_TOOLS_DATABASE_URL](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_database_url)**
- **[SENZING_TOOLS_ENABLE_G2CONFIG](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_enable_g2config)**
- **[SENZING_TOOLS_ENABLE_G2CONFIGMGR](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_enable_g2configmgr)**
- **[SENZING_TOOLS_ENABLE_G2DIAGNOSTIC](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_enable_g2diagnostic)**
- **[SENZING_TOOLS_ENABLE_G2ENGINE](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_enable_g2engine)**
- **[SENZING_TOOLS_ENABLE_G2PRODUCT](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_enable_g2product)**
- **[SENZING_TOOLS_ENGINE_CONFIGURATION_JSON](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_configuration_json)**
- **[SENZING_TOOLS_ENGINE_LOG_LEVEL](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_log_level)**
- **[SENZING_TOOLS_ENGINE_MODULE_NAME](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_module_name)**
- **[SENZING_TOOLS_GRPC_PORT](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_grpc_port)**
- **[SENZING_TOOLS_LOG_LEVEL](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_log_level)**

## References

- [Command reference](https://hub.senzing.com/senzing-tools/senzing-tools_servegrpc.html)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
