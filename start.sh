#!/bin/bash
set -e
source ./../fabric/scripts/gen-utils.sh

runCMD export GOPATH=$GOPATH:$PWD
runCMD export GOBIN=$PWD/bin
runCMD export FDD_MONGO_ADDR="localhost:37017"
runCMD export FDD_MONGO_DB="FDDApi"
runCMD export FDD_MONGO_USER="FDDUser"
runCMD export FDD_MONGO_PASS="Pass"
runCMD export FDD_SERVER_PORT=4000
runCMD export FDD_APP_VERSION=v2

CURR_DIR=$PWD
runCMD cd ./src
runCMD go build -o go-app-server
mv go-app-server ./../bin/
cd $CURR_DIR
./bin/go-app-server

if [ -f "log.txt" ]; then 
echo file 
rm log.txt

fi