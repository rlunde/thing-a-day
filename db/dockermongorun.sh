#!/bin/bash
# run mongo and map ports
# make sure environment has MONGO_DATA_PATH set to absolute path of data directory here
docker run -d -p 37017:27017 --rm --name tadmongo -v $MONGO_DATA_PATH:/data/db mongo
# connect to this via mongo localhost:37017/thing-a-day
