# protobuf

### Install

- Install protobuf by following below link:

https://github.com/golang/protobuf

### How to build

- Run:
    Single file
    ```
        protoc -I. --go_out=plugins=grpc:. your_file_want_to_build.proto
    ```
    Build all
    ```
        make
    ```

### How to test:
- Install prototool (https://github.com/uber/prototool#installation)
    ```
    brew install prototool
    prototool lint
    ```

### How to auto align:
- Install clang-format
    ```
    brew install clang-format
    ```
- Generate .clang-format file
    ```
    clang-format -assume-filename=<any file *.proto you want> -style=llvm -dump-config > .clang-format
    ```
    EX: clang-format -assume-filename=tnx.proto -style=llvm -dump-config > .clang-format
- Config clang format:
    Add or update below line to .clang-format which had been generated prev step.
    ```
    AlignConsecutiveAssignments: true
    AlignConsecutiveDeclarations: true
    ```
- Add the following to your settings.json file:
    ```
    "[proto3]": {
        "editor.formatOnSave": true,
    },
    ```
- Double save(Cmd + S) to see the magic