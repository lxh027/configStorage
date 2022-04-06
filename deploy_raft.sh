#!/bin/bash
# shellcheck disable=SC2143
# shellcheck disable=SC2164
# shellcheck disable=SC2103

ver=$(cat /proc/sys/kernel/random/uuid | md5sum |cut -c 1-12)
raft_id="raft000"

if [[ -n $(docker ps -a | grep raft_peer) ]]; then
  docker stop raft_peer && docker rm raft_peer
fi

if [[ -n $(docker image ls | grep raft_peer) ]]; then
  echo 'y' | docker container prune
fi

cp ./docker/raft/Dockerfile ./Dockerfile

docker build -t raft_peer:$ver .

docker run -itd \
  -p 2001:2000 \
  -p 3001:3000 \
  -v /home/lxh001/Workspace/go/configStorage/logs:/app/logs \
  -e rfid=$raft_id \
  -e rfpt=2001 \
  -e cpt=3001 \
  --name raft_peer1 \
  raft_peer:$ver

docker run -itd \
  -p 2002:2000 \
  -p 3002:3000 \
  -v /home/lxh001/Workspace/go/configStorage/logs:/app/logs \
  -e rfid=$raft_id \
  -e rfpt=2002 \
  -e cpt=3002 \
  --name raft_peer2 \
  raft_peer:$ver


docker run -itd \
  -p 2003:2000 \
  -p 3003:3000 \
  -v /home/lxh001/Workspace/go/configStorage/logs:/app/logs \
  -e rfid=$raft_id \
  -e rfpt=2003 \
  -e cpt=3003 \
  --name raft_peer3 \
  raft_peer:$ver