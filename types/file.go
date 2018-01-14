package types

type File struct {
	Path     string
	Hash     string
	DiskPath string `json:"-"`
}
