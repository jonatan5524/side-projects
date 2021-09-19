package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen -out objectbox

type Project struct {
	Id   uint64
	Name string
}
