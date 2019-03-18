docker rm -f tournaments
docker pull cjzhang/tournaments
docker run -d \
--name tournaments \
--network smashgg \
cjzhang/tournaments