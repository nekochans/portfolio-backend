#!/bin/sh

cp .realize.debug.yaml .realize.yaml

docker-compose up --build -d
