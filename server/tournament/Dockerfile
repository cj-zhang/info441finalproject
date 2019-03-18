# add the necessary instructions
# to create a Docker container image
# for your Go API gateway server
FROM alpine
COPY tournaments /tournaments
RUN apk add --no-cache ca-certificates
EXPOSE 80
ENTRYPOINT ["/tournaments"]