# Mục đích
Thử nghiệm tạo microservices bằng Nameko

# Hướng dẫn
```bash
# start rabbitmq
docker run --rm -it -p 5672:5672 -p 15672:15672 rabbitmq:3-management

pip3 install nameko

# start service
nameko run --config config.yaml helloworld

# consume service
nameko shell --config config.yaml
>>> n.rpc.greeting_service.hello(name='Tuan Nguyen')
>>> exit()

# test stop service
# Ctrl-C
```

# Tài liệu tham khảo:
- https://daydreamer3d.github.io/tutorials/nameko-rabbitmq/service_in_action.html

