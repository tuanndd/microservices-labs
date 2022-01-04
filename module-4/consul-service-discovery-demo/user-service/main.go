package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

type User struct {
	ID       uint64    `json:"id"`
	Username string    `json:"username"`
	Products []product `json:"products"`
}

type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// dang ky service voi consul
func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "user-service"   // id
	registration.Name = "user-service" // name cua service
	address := hostname()
	registration.Address = address // ip
	p, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = p // port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, p) // healthcheck
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

// lookup service tu consul
func lookupServiceWithConsul(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := consul.Agent().Services()
	if err != nil {
		return "", err
	}

	srvc := services[serviceName] // chua thong tin cua "product-service"

	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

func main() {
	registerServiceWithConsul()

	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/user-products", UserProduct)

	fmt.Printf("user service is up on port: %s", port())

	http.ListenAndServe(port(), nil)
}

// web handler "/healthcheck"
func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `user service is good`)
}

// web hander "/user-products"
func UserProduct(w http.ResponseWriter, r *http.Request) {
	p := []product{}

	url, err := lookupServiceWithConsul("product-service")

	fmt.Println("URL: ", url)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Get(url + "/products")
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	u := User{
		ID:       1,
		Username: "didiyudha@gmail.com",
	}
	u.Products = p
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}

// tra ve port tu bien moi truong "USER_SERVICE_PORT" | 8080
func port() string {
	p := os.Getenv("USER_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8080"
	}
	return fmt.Sprintf(":%s", p)
}

// tra ve hostname
func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
