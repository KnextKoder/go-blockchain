package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	traLocal := tra.(*LocalTransport)
    trbLocal := trb.(*LocalTransport)


	assert.Equal(t, traLocal.peers[trb.Addr()], trbLocal)
    assert.Equal(t, trbLocal.peers[tra.Addr()], traLocal)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <- trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.Addr())
}
