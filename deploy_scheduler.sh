
ver=$(cat /proc/sys/kernel/random/uuid | md5sum |cut -c 1-12)

if [[ -n $(docker ps -a | grep scheduler) ]]; then
  docker stop scheduler && docker rm scheduler
fi

if [[ -n $(docker image ls | grep scheduler) ]]; then
  echo 'y' | docker container prune
fi

cp ./docker/scheduler/Dockerfile ./Dockerfile

docker build -t scheduler:$ver .

docker run -itd \
  -p 2888:2888 \
  -p 2889:2889 \
  --name scheduler \
  scheduler:$ver

rm ./Dockerfile