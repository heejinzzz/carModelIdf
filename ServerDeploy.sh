#!/usr/bin/env bash

parameters=$(getopt -a -o h:p:a: -l ip:,port:,access_token: -- "$@" | sed $'s/\'//g')
set -- $parameters

ip=""
port="7180"

while [[ $1 ]]
do
  if [[ $1 == "-h" ]] || [[ $1 == "--ip" ]]
  then
    ip=$2
    shift
  elif [[ $1 == "-p" ]] || [[ $1 == "--port" ]]
  then
    port=$2
    shift
  elif [[ $1 == "-a" ]] || [[ $1 == "--access_token" ]]
  then
    access_token=$2
    shift
  elif [[ $1 == "--" ]]
  then
    shift
    break
  else
    echo "Unknown option: $1"
  fi
  shift
done

if [[ $access_token == "" ]]
then
  echo "Error: no access_token provide. Please provide your Baidu access_token by using '-a' or '--access_token' option."
  exit 7
fi

if [[ $ip ]]
then
  echo "set server ip: $ip"
fi
echo "set server port: $port"

sed -i -e "/var accessToken = /c\var accessToken = \"$access_token\"" server/server.go
sed -i -e "/var ip = /c\var ip = \"$ip\"" server/server.go
sed -i -e "/var port = /c\var port = \"$port\"" server/server.go

echo "CarModelIdf Server Start"
go run server/server.go