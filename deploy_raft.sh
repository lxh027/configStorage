#!/bin/bash
# shellcheck disable=SC2143
# shellcheck disable=SC2164
# shellcheck disable=SC2103

ver=$(cat /proc/sys/kernel/random/uuid | md5sum |cut -c 1-12)
raft_id="raft000"
raft_host="172.17.0.1"

if [[ -n $(docker ps -a | grep raft_peer1) ]]; then
  docker stop raft_peer1 && docker rm raft_peer1
fi

if [[ -n $(docker ps -a | grep raft_peer2) ]]; then
  docker stop raft_peer2 && docker rm raft_peer2
fi

if [[ -n $(docker ps -a | grep raft_peer3) ]]; then
  docker stop raft_peer3 && docker rm raft_peer3
fi

if [[ -n $(docker image ls | grep raft_peer) ]]; then
  echo 'y' | docker container prune
fi

cp ./docker/raft/Dockerfile ./Dockerfile

docker build -t raft_peer:$ver .

docker run -itd \
  -p 2001:2000 \
  -p 3001:3000 \
  -e RFID=$raft_id \
  -e RFHOST=$raft_host \
  -e RFPT="2001" \
  -e CPT="3001" \
  --name raft_peer1 \
  raft_peer:$ver

docker run -itd \
  -p 2002:2000 \
  -p 3002:3000 \
  -e RFID=$raft_id \
  -e RFHOST=$raft_host \
  -e RFPT="2002" \
  -e CPT="3002" \
  --name raft_peer2 \
  raft_peer:$ver


docker run -itd \
  -p 2003:2000 \
  -p 3003:3000 \
  -e RFID=$raft_id \
  -e RFHOST=$raft_host \
  -e RFPT="2003" \
  -e CPT="3003" \
  --name raft_peer3 \
  raft_peer:$ver

rm ./Dockerfile