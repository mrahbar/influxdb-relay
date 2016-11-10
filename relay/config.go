package relay

import (
	"encoding/json"
)

type Config struct {
	HTTPRelays []HTTPConfig `json:"http"`
	UDPRelays  []UDPConfig  `json:"udp"`
}

type HTTPConfig struct {
	// Name identifies the HTTP relay
	Name string `json:"name"`

	// Addr should be set to the desired listening host:port
	Addr string `json:"bind-addr"`

	// Set certificate in order to handle HTTPS requests
	SSLCombinedPem string `json:"ssl-combined-pem"`

	// Default retention policy to set for forwarded requests
	DefaultRetentionPolicy string `json:"default-retention-policy"`

	// Outputs is a list of backed servers where writes will be forwarded
	Outputs []HTTPOutputConfig `json:"output"`
}

type HTTPOutputConfig struct {
	// Name of the backend server
	Name string `json:"name"`

	// Location should be set to the URL of the backend server's write endpoint
	Location string `json:"location"`

	// Timeout sets a per-backend timeout for write requests. (Default 10s)
	// The format used is the same seen in time.ParseDuration
	Timeout string `json:"timeout"`

	// Buffer failed writes up to maximum count. (Default 0, retry/buffering disabled)
	BufferSizeMB int `json:"buffer-size-mb"`

	// Maximum batch size in KB (Default 512)
	MaxBatchKB int `json:"max-batch-kb"`

	// Maximum delay between retry attempts.
	// The format used is the same seen in time.ParseDuration (Default 10s)
	MaxDelayInterval string `json:"max-delay-interval"`

	// Skip TLS verification in order to use self signed certificate.
	// WARNING: It's insecure. Use it only for developing and don't use in production.
	SkipTLSVerification bool `json:"skip-tls-verification"`
}

type UDPConfig struct {
	// Name identifies the UDP relay
	Name string `json:"name"`

	// Addr is where the UDP relay will listen for packets
	Addr string `json:"bind-addr"`

	// Precision sets the precision of the timestamps (input and output)
	Precision string `json:"precision"`

	// ReadBuffer sets the socket buffer for incoming connections
	ReadBuffer int `json:"read-buffer"`

	// Outputs is a list of backend servers where writes will be forwarded
	Outputs []UDPOutputConfig `json:"output"`
}

type UDPOutputConfig struct {
	// Name identifies the UDP backend
	Name string `json:"name"`

	// Location should be set to the host:port of the backend server
	Location string `json:"location"`

	// MTU sets the maximum output payload size, default is 1024
	MTU int `json:"mtu"`
}

// LoadConfigFile parses the specified file into a Config object
func LoadConfigJson(jsonObject string) (cfg Config, err error) {
	return cfg, json.Unmarshal([]byte(jsonObject), &cfg)
}
