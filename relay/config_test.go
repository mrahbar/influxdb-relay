package relay

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestFullConfigJson(t *testing.T) {
	assert := assert.New(t)

	var configJson = `{
			  "http": [{
			    "name": "example-http",
			    "bind-addr": "127.0.0.1:9096",
			    "output": [
			      { "name": "local1", "location": "http://127.0.0.1:8086/write" }
			    ]
			  }],
			  "udp": [{
			    "name": "example-udp",
			    "bind-addr": "127.0.0.1:9096",
			    "read-buffer": 0,
			    "output": [
			       { "name":"local1", "location": "127.0.0.1:8089", "mtu":512 },
			    { "name":"local2", "location": "127.0.0.1:7089", "mtu":1024 }
			    ]
			  }]
			}`

	cfg, err := LoadConfigJson(configJson)
	if err != nil {
		fmt.Println("error:", err)
		t.Fail()
	}

	assert.Equal(len(cfg.HTTPRelays), 1, "HTTPRelays should only have 1 element")
	assert.Equal(cfg.HTTPRelays[0].Name, "example-http", "Name must be example-http")
	assert.Equal(cfg.HTTPRelays[0].Addr, "127.0.0.1:9096", "Addr must be 127.0.0.1:9096")
	assert.Equal(len(cfg.HTTPRelays[0].Outputs), 1, "HTTPRelay Outputs should have 1 element")

	assert.Equal(cfg.UDPRelays[0].Name, "example-udp", "Name must be example-udp")
	assert.Equal(len(cfg.UDPRelays), 1, "UDPRelays should only have 1 element")
	assert.Equal(cfg.UDPRelays[0].Addr, "127.0.0.1:9096", "Addr must be 127.0.0.1:9096")
	assert.Equal(len(cfg.UDPRelays[0].Outputs), 2, "UDPRelay Outputs should have 2 element")
}

func TestHttpOnlyConfigJson(t *testing.T) {
	assert := assert.New(t)

	var configJson = `{
			  "http": [{
			    "name": "example-http",
			    "bind-addr": "127.0.0.1:9096",
			    "output": [
			      { "name": "local1", "location": "http://127.0.0.1:8086/write" }
			    ]
			  }]
			}`

	cfg, err := LoadConfigJson(configJson)
	if err != nil {
		fmt.Println("error:", err)
		t.Fail()
	}

	assert.Equal(len(cfg.HTTPRelays), 1, "HTTPRelays should only have 1 element")
	assert.Equal(cfg.HTTPRelays[0].Name, "example-http", "Name must be example-http")
	assert.Equal(cfg.HTTPRelays[0].Addr, "127.0.0.1:9096", "Addr must be 127.0.0.1:9096")
	assert.Equal(len(cfg.HTTPRelays[0].Outputs), 1, "HTTPRelay Outputs should have 1 element")

	assert.Equal(len(cfg.UDPRelays), 0, "UDPRelays should only have 1 element")
}
