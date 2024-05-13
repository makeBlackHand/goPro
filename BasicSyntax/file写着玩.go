package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Server struct {
	IpAddr string `yaml:"ipAddr"`
	Port   int    `yaml:"port"`
}

type RedisServer struct {
	RedisIpAddr string `yaml:"redisIpAddr"`
	RedisPort   int    `yaml:"redisPort"`
}

type Conf struct {
	Server      `yaml:"server"`
	RedisServer `yaml:"redisServer"`
}

func main() {
	server1 := Server{
		IpAddr: "1.1.1.1",
		Port:   9090,
	}

	redisServer := RedisServer{
		RedisIpAddr: "dewred.com",
		RedisPort:   6379,
	}

	conf1 := Conf{
		server1,
		redisServer,
	}

	data, err := yaml.Marshal(conf1)
	if err != nil {
		log.Println(err)
	}

	file, err := os.OpenFile("./conf.yaml", os.O_CREATE|os.O_RDWR, 0777)
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	defer file.Close()

	fmt.Println(string(data))

}
