package server

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)

type mockRouter struct {
	called    bool
	port      string
	returnErr error
}

func (m *mockRouter) Run(addr ...string) error {
	m.called = true
	if len(addr) > 0 {
		m.port = addr[0]
	}
	return m.returnErr
}

func TestServerWithCorrectPort(t *testing.T) {
	mock := &mockRouter{}
	err := Server(mock)
	assert.NoError(t, err)
	assert.True(t, mock.called)
	assert.Equal(t, GetServerPort(), mock.port)
}

func TestServerWithError(t *testing.T) {
	mock := &mockRouter{returnErr: errors.New("failed to start")}
	err := Server(mock)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to start")
}
