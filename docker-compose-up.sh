#!/bin/sh

cp .air.normal.conf .air.conf

docker-compose up --build -d
