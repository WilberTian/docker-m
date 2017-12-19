package vos

import (
	"time"
)

type ImagesVo struct {
	ID          string `json:"Id"`
	RepoTags    []string
	Created     int64
	Size        int64
	VirtualSize int64
	ParentID    string
	RepoDigests []string
	Labels      map[string]string
}

type RootFS struct {
	Type   string   `json:"Type,omitempty"`
	Layers []string `json:"Layers,omitempty"`
}

type ImageDetailVo struct {
	ID              string    `json:"Id"`
	RepoTags        []string  `json:"RepoTags,omitempty"`
	Parent          string    `json:"Parent,omitempty"`
	Comment         string    `json:"Comment,omitempty"`
	Created         time.Time `json:"Created,omitempty"`
	Container       string    `json:"Container,omitempty"`
	ContainerConfig Config    `json:"ContainerConfig,omitempty"`
	DockerVersion   string    `json:"DockerVersion,omitempty"`
	Author          string    `json:"Author,omitempty"`
	Config          *Config   `json:"Config,omitempty"`
	Architecture    string    `json:"Architecture,omitempty"`
	Size            int64     `json:"Size,omitempty"`
	VirtualSize     int64     `json:"VirtualSize,omitempty"`
	RepoDigests     []string  `json:"RepoDigests,omitempty"`
	RootFS          *RootFS   `json:"RootFS,omitempty"`
	OS              string    `json:"Os,omitempty"`
}

type ImageHistoryVo struct {
	ID        string   `json:"Id"`
	Tags      []string `json:"Tags,omitempty"`
	Created   int64    `json:"Created,omitempty"`
	CreatedBy string   `json:"CreatedBy,omitempty"`
	Size      int64    `json:"Size,omitempty"`
}