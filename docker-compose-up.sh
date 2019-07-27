#!/bin/sh

cp .realize.normal.yaml .realize.yaml

docker-compose up --build -d
