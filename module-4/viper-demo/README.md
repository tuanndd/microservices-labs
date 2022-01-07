# Mục đích
- Thử nghiệm dùng viper để đọc cấu hình từ local file và Consul

# Hướng dẫn
### Đọc config từ local file
```bash
go mod tidy

go run local.go
```

### Đọc config từ Consul
```bash
# chuẩn bị
# start consul
consul agent -dev -client 0.0.0.0

# tạo KV trên consul để lưu config
http://vagrant-ip:8500/ui/dc1/kv/create

key=config

value=
{
  "port": 10666,
  "mysql":{
  	"url":"(127.0.0.1:3306)/dbname",
    "username":"root",
    "password":"123456"
  },
  "redis":["127.0.0.1:6377", "127.0.0.1:6378", "127.0.0.1:6379"],
  "smtp":{
  	"enable":true,
    "addr":"mail_addr",
    "username":"mail_user",
    "password":"mail_password",
    "to":["a@gmail.com","b@gmail.com"]
  }
}

# đọc config từ consul
go run consul.go

# TODO: watch config changed (https://madflojo.medium.com/using-viper-with-consul-to-configure-golang-applications-eaa84394b8de)