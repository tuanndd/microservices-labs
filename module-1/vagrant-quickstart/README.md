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

### Bước 3: Run Linux VM

```bash
# edit Vagrantfile, thay đổi cấu hình của VM: CPU, memory, disk size và ip address
# up vm
vagrant up h1

# ssh đến vm
vagrant ssh h1

# stop vm
vagrant halt h1

# delete vm
vagrant destroy h1
```