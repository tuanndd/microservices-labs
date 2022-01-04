# Mục đích
Thử nghiệm triển khai WordPress bằng docker-compose và Kubernetes

# Hướng dẫn
### Deploy Wordpress bằng docker-compose
```bash
cd wp-compose
docker-compose

curl http://[ip-address]:8000
```

### Deploy WordPress bằng Kubernetes
```bash
cd wp-k8s

kubectl -- apply -f mysql.yml
kubectl -- apply -f wordpress.yml

kubectl get pods
kubectl -- delete -f wordpress.yml
kubectl -- delete -f mysql.yml

curl http://[ip-address]:30000
```

# Tài liệu tham khảo
- https://docs.docker.com/samples/wordpress/
- https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/
- https://minikube.sigs.k8s.io/docs/start/

## Phụ lục: Install minikube trên vagrant host
```bash
# Start centos box
vagrant up h1
vagrant ssh h1

# Fix bugs trên centos box trước khi install minikube:
sudo apt-get install -y conntrack
sudo swapoff -a
sudo vi /etc/sysctl.conf
 net.bridge.bridge-nf-call-iptables = 1
sudo sysctl -p

# cài đặt minikube
...

# edit path, config sau khi cài đặt minikube
export PATH=$PATH:/usr/local/bin

sudo minikube start --driver=none
alias kubectl='minikube kubectl'

## Copy config from /root/.kube, .minikube to $HOME
cp ...
vi ~/.kube/config (update path to $HOME/...)

# Test
## tạo deployment, expose service
kubectl -- create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.4
kubectl -- expose deployment hello-minikube --type=NodePort --port=8080
kubectl get svc
minikube service hello-minikube
kubectl port-forward service/hello-minikube 7080:8080

## tạo loadbalancer
kubectl -- create deployment balanced --image=k8s.gcr.io/echoserver:1.4
kubectl -- expose deployment balanced --type=LoadBalancer --port=8080
minikube tunnel
kubectl get services balanced

## xóa service, deployment
kubectl -- delete svc
kubectl -- delete deployments

# get pods
kubectl get -- pods --all-namespaces