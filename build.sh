#!/bin/bash

CGO_ENABLED=0
go build -o app -tags netgo .
