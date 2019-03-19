export TOURNAMENTADDR=":80"
export RABBITADDR="amqp://guest:guest@rabbit:5672/"
export MYSQL_ROOT_PASSWORD="root"
export DSN="root:$MYSQL_ROOT_PASSWORD@tcp(mydb:3306)/mysql"


docker rm -f tournaments
docker pull cjzhang/tournaments
docker run -d \
--name tournaments \
--network smashgg \
-e RABBITADDR=$RABBITADDR \
-e TOURNAMENTADDR=$TOURNAMENTADDR \
-e DSN=$DSN \
cjzhang/tournaments