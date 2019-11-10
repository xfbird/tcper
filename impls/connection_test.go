package impls

import "testing"

func TestNewConnection(t *testing.T) {
	c := NewConnection(nil, 0, nil)

	if c == nil {
		panic("New connection method not create the connection.")
	}
}

func TestConnection_GetConnectionId(t *testing.T) {
	cid := 10

	c := NewConnection(nil, uint32(cid), nil)

	if c == nil {
		panic("New connection method not create the connection.")
	}

	if c.GetConnectionId() != uint32(cid) {
		panic("GetConnectionId Method broken.")
	}
}
