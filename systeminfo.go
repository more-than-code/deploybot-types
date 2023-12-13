package types

type DiskInfo struct {
	TotalSize uint64 `json:"totalSize"`
	AvailSize uint64 `json:"availSize"`
	Path      string `json:"path"`
}
