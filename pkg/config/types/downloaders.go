package types

import (
	"slices"
	"strings"
)

var _ Provider = (*ResultDownloaders)(nil)

type ResultDownloaders struct {
	// Disabled is a list of downloaders that are disabled.
	Disabled []string `yaml:"Disabled,omitempty"`
	// Timeout specifies the maximum time allowed for a download operation.
	Timeout Duration               `yaml:"Timeout,omitempty"`
	Types   ResultDownloadersTypes `yaml:"Types,omitempty"`
}

type ResultDownloadersTypes struct {
	IPFS IpfsDownloader `yaml:"IPFS,omitempty"`
}

func (r ResultDownloaders) IsNotDisabled(kind string) bool {
	return !slices.ContainsFunc(r.Disabled, func(s string) bool {
		return strings.ToLower(s) == strings.ToLower(kind)
	})
}

type IpfsDownloader struct {
	// Endpoint specifies the multi-address to connect to for IPFS. e.g /ip4/127.0.0.1/tcp/5001
	Endpoint string `yaml:"Endpoint,omitempty"`
}