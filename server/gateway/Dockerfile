# add the necessary instructions
# to create a Docker container image
# for your Go API gateway server
FROM alpine
COPY gateway /gateway
RUN apk add --no-cache ca-certificates
EXPOSE 443
ENTRYPOINT ["/gateway"]