#!/bin/sh

cp .air.debug.conf .air.conf

docker-compose up --build -d
