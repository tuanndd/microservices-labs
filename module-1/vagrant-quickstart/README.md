# Mục đích
Run Centos VM bằng Vagrant trên Windows

# Hướng dẫn
### Bước 1: Download và cài đặt VirtualBox
https://www.virtualbox.org/wiki/Downloads

### Bước 2: Download và cài đặt Vagrant 
https://www.vagrantup.com/downloads

### Bước 3: Run CentOS VM

```bash
# up vm
vagrant up h1

# ssh đến vm
vagrant ssh h1

# stop vm
vagrant halt h1

# delete vm
vagrant destroy h1
```