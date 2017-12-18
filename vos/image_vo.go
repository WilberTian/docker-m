package vos

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
