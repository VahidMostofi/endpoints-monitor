#!/bin/bash

set -e

docker build ./nginx-gateway/ -t vahidmostofi/nginx-gateway:latest
docker push vahidmostofi/nginx-gateway:latest

docker build ./telegraf-agent/ -t vahidmostofi/telegraf-agent:latest
docker push vahidmostofi/telegraf-agent:latest