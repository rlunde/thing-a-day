#!/bin/bash
# run mongo and map ports
# make sure environment has MONGO_DATA_PATH set to absolute path of data directory here
docker run -d -p 37017:27017 --rm --name tadmongo -v $MONGO_DATA_PATH:/data/db mongo
# connect to this via mongo localhost:37017/thing-a-day
# it doesn't work to run both mongo and thing-a-day in separate containers on my mac
# (long detailed network explanation is elsewhere) -- for now, just going to run
# mongo as local service and run thing-a-day as the container
