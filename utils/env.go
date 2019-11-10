package utils

import (
	"os"
	"strconv"
)

const (
	ServerName = "SERVER_NAME"
	TCPVersion = "TCP_VERSION"
	IpAddr     = "IP_ADDR"
	Port       = "PORT"
	MaxBuffer  = "MAX_BUFFER"
)

func Getenv(key string, default_value ...string) string {
	if r := os.Getenv(key); r != "" {
		return r
	} else {
		if len(default_value) == 0 {
			return ""
		}

		return default_value[0]
	}

	return ""
}

func GetServerName() string {
	return Getenv(ServerName, "Tcper Server")
}

func GetTCPVersion() string {
	return Getenv(TCPVersion, "tcp4")
}

func GetIpAddr() string {
	return Getenv(IpAddr, "0.0.0.0")
}

func GetPort() string {
	return Getenv(Port, "8888")
}

func GetMaxBuffer() int {
	max_buffer, _ := strconv.Atoi(Getenv(MaxBuffer, "8192"))
	return max_buffer
}
