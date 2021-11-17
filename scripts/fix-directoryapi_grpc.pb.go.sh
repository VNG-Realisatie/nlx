#!/bin/bash
# This will replace the service names in the generated directory api grpc code
# So that the directory api proto is backwards compatible

PROTO_LOCATION="${1}"
if [ "${1}" = "" ]; then
  echo "Missing proto file location argument"
  exit 1
fi

sed -i '/var DirectoryRegistration_ServiceDesc = grpc.ServiceDesc{/{n;s/.*/	ServiceName: "registrationapi.DirectoryRegistration",/}' "$PROTO_LOCATION"
grep "registrationapi.DirectoryRegistration" < "$PROTO_LOCATION" || exit 1
sed -i '/var DirectoryInspection_ServiceDesc = grpc.ServiceDesc{/{n;s/.*/	ServiceName: "inspectionapi.DirectoryInspection",/}' "$PROTO_LOCATION"
grep "inspectionapi.DirectoryInspection" < "$PROTO_LOCATION" || exit 1
