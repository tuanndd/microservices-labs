# Mục đích
Thử nghiệm triển khai WordPress bằng docker-compose và Kubernetes

# Hướng dẫn

### Deploy Wordpress bằng docker-compose
```bash
cd wp-compose
docker-compose

curl http://vagrant-ip:8000
```

### Deploy WordPress bằng Kubernetes
```bash
cd wp-k8s

# start minikube
minikube start

alias kubectl='minikube kubectl --'

# deploy wordpress
kubectl apply -f mysql.yml
kubectl apply -f wordpress.yml

kubectl get pods
kubectl get svc

# test
kubectl port-forward svc/wordpress 8000:80 --address='0.0.0.0'

curl http://vagrant-ip:8000

# delete deployment
kubectl delete -f wordpress.yml
kubectl delete -f mysql.yml
```

# Tài liệu tham khảo
- https://docs.docker.com/samples/wordpress/
- https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/
- https://minikube.sigs.k8s.io/docs/start/

