package impls

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()

	if server == nil {
		panic("Server is not create.")
	}
}

func TestServer_listenTCP(t *testing.T) {
	server := NewServer()

	if server == nil {
		panic("Server is not create.")
	}

	server.listenTCP()
}
