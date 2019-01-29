# directory

The directory is a centralized service that provides a public API serving a list of organizations and services registered in the NLX network. It also provides outways with connection details for these services.

## API

The directory API is a gRPC API, it's specification resides in the folder `directoryapi`, which also contains the protobuf/grpc generated code (service interface and client implementation) as well as a generated swagger specification and grpc > http/json gateway. The service impementation resides in `directoryservice`.
