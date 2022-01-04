# Mục đích
Thử nghiệm tạo microservices bằng Nameko

# Hướng dẫn
```bash
pip install nameko

# start service
nameko run --config config.yaml helloworld

# consume service
nameko shell --config config.yaml
>>> n.rpc.greeting_service.hello(name='Tuan Nguyen')
>>> exit()
```

# Tài liệu tham khảo:
- https://daydreamer3d.github.io/tutorials/nameko-rabbitmq/service_in_action.html

