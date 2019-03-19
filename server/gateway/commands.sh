docker rm -f smash
# docker rm -f mongoInstance
docker pull cjzhang/smash
docker rm -f mydb
docker pull cjzhang/smashdb
docker rm -f sessionstore
docker rm -f rabbit
docker network rm smashgg
docker network create smashgg
export MYSQL_ROOT_PASSWORD="root"
docker run -d \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=mysql \
--name mydb \
--network smashgg \
cjzhang/smashdb
sleep 10
docker run -d --name sessionstore --network smashgg redis
sleep 10

# docker run -d \
# -e MONGO_INITDB_DATABASE=mongodb \
# --name mongoInstance \
# --network smashgg \
# mongo

export TLSCERT=/etc/letsencrypt/live/smash.chenjosephzhang.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/smash.chenjosephzhang.me/privkey.pem
export DSN="root:$MYSQL_ROOT_PASSWORD@tcp(mydb:3306)/mysql"
export REDISADDR="sessionstore:6379"
export SESSIONKEY="placeholder"
export RABBITADDR="amqp://guest:guest@rabbit:5672/"
# export SUMMARYADDR="summary:80"
# export MESSAGESADDR="messaging:80"
docker run -d \
--name rabbit \
--network smashgg \
rabbitmq:3-management

sleep 10
docker run -d \
--name smash \
--network smashgg \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e DSN=$DSN \
-e SUMMARYADDR=$SUMMARYADDR \
-e RABBITADDR=$RABBITADDR \
cjzhang/smash