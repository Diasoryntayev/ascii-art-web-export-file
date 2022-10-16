FROM golang:1.19-alpine

WORKDIR /build

COPY . .

RUN go build -o server cmd/*.go


LABEL author="DiasOryntayev"

EXPOSE 8080

CMD ["./server"]


#command for Docker
#docker images -a   ---> check docker images
#docker ps -a         ---> check docker containers
#docker container stop DOCKERID    ---> stop docker container
#docker build -t <image_name> .  ---> build new image
#docker run -p 8080:8080 <IMAGE ID>  ---> run container
#docker system prune -a ---> remove all images and containers