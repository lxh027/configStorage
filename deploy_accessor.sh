
ver=$(cat /proc/sys/kernel/random/uuid | md5sum |cut -c 1-12)

if [[ -n $(docker ps -a | grep accessor) ]]; then
  docker stop accessor && docker rm accessor
fi

if [[ -n $(docker image ls | grep accessor) ]]; then
  echo 'y' | docker container prune
fi

cp ./docker/accessor/Dockerfile ./Dockerfile

docker build -t accessor:$ver .

docker run -itd \
  -p 5000:5000 \
  --name accessor \
  accessor:$ver

rm ./Dockerfile