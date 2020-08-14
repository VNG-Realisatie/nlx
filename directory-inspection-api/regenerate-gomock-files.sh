#!/bin/bash

mockgen -source=./pkg/database/database.go -destination=./pkg/database/mock/database.go -package=mock
