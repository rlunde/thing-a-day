#!/bin/bash
# run mongo and map ports
cd ./db
source source_me.sh
./dockermongorun.sh
cd ..
# todo: map mongo ports in Dockerfile for thing-a-day
docker run -p 8084:8084  --name thing-a-day --rm thing-a-day
