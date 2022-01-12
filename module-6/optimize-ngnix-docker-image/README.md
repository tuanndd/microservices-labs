# Mục đích
Thử nghiệm giảm kích thước của docker-image nginx

# Hướng dẫn
```bash
# pull ngnix docker-images
docker image pull nginx:latest
docker image pull nginx:alpine

# build nginx.slim docker-image từ nginx:latest
curl -O https://downloads.dockerslim.com/releases/1.37.3/dist_linux.tar.gz
cd dist/
./docker-slim build --target nginx:latest

# so sánh kích thước của các docker-image nginx
docker images | grep nginx

# test các docker-image nginx
docker run -it --rm -d -p 8081:80 nginx:latest
docker run -it --rm -d -p 8082:80 nginx:alpine
docker run -it --rm -d -p 8083:80 nginx.slim

curl http://vagrant_ip:8081
curl http://vagrant_ip:8082
curl http://vagrant_ip:8083

```

