package publicapi

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// AllowedCORSOrigin is the allowed origin for CORS requests.
	// Used to allow webui to connect to the backend API server.
	// TODO: make this more restricted or configurable
	AllowedCORSOrigin = "*"
)

var WebsocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: make this more restricted or configurable
		return true
	},
}

type Config struct {
	// These are TCP connection deadlines and not HTTP timeouts. They don't control the time it takes for our handlers
	// to complete. Deadlines operate on the connection, so our server will fail to return a result only after
	// the handlers try to access connection properties
	// ReadHeaderTimeout is the amount of time allowed to read request headers
	ReadHeaderTimeout time.Duration
	// ReadTimeout is the maximum duration for reading the entire request, including the body
	ReadTimeout time.Duration
	// WriteTimeout is the maximum duration before timing out writes of the response.
	// It doesn't cancel the context and doesn't stop handlers from running even after failing the request.
	// It is for added safety and should be a bit longer than the request handler timeout for better error handling.
	WriteTimeout time.Duration

	// This represents maximum duration for handlers to complete, or else fail the request with 503 error code.
	RequestHandlerTimeout time.Duration

	// MaxBytesToReadInBody is used by safeHandlerFuncWrapper as the max size of body
	MaxBytesToReadInBody string

	// ThrottleLimit is the maximum number of requests per second
	ThrottleLimit int

	// Protocol
	Protocol string

	// LogLevel is the minimum log level to log requests
	LogLevel string
}

// defaultConfig default values for Config
var defaultConfig = Config{
	ReadHeaderTimeout:     5 * time.Second,
	ReadTimeout:           20 * time.Second,
	WriteTimeout:          45 * time.Second,
	RequestHandlerTimeout: 30 * time.Second,
	MaxBytesToReadInBody:  "10MB",
	ThrottleLimit:         1000,
	Protocol:              "http",
	LogLevel:              "trace",
}

// DefaultConfig returns the default configuration for the public API server.
func DefaultConfig() Config {
	return defaultConfig
}

type Option func(*Config)

func WithReadHeaderTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.ReadHeaderTimeout = t
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.WriteTimeout = t
	}
}

func WithRequestHandlerTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.RequestHandlerTimeout = t
	}
}

func WithMaxBytesToReadInBody(size string) Option {
	return func(c *Config) {
		c.MaxBytesToReadInBody = size
	}
}

func WithThrottleLimit(limit int) Option {
	return func(c *Config) {
		c.ThrottleLimit = limit
	}
}

func WithProtocol(protocol string) Option {
	return func(c *Config) {
		c.Protocol = protocol
	}
}

func WithLogLevel(logLevel string) Option {
	return func(c *Config) {
		c.LogLevel = logLevel
	}
}

func NewConfig(opts ...Option) *Config {
	config := DefaultConfig()

	for _, opt := range opts {
		opt(&config)
	}

	return &config
}
