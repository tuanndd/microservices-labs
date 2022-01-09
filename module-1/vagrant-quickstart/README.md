# Mục đích
Run Linux VM bằng Vagrant

# Hướng dẫn
### Bước 1: Download và cài đặt VirtualBox
https://www.virtualbox.org/wiki/Downloads

### Bước 2: Download và cài đặt Vagrant 
https://www.vagrantup.com/downloads

```bash
vagrant plugin install vagrant-disksize
```

### Bước 3: Test Linux VM

edit *Vagrantfile* - thay đổi CPU, memory, disk size và ip address của máy ảo. Lưu ý IP máy ảo phải thuộc range 10.0.0.0/8, 172.16.0.0/12 hoặc 192.168.0.0/16  và không trùng với range IP của mạng LAN hiện tại.  

```bash
# up vm
vagrant up h1

# ssh đến vm
vagrant ssh h1

# dùng ssh-key
ssh -i insecure_private_key vagrant@[vagrant-ip]

# stop vm
vagrant halt h1

# delete vm
#vagrant destroy h1
```