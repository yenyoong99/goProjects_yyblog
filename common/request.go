package common

// update module
// 1. update all: object overwrite
// 2. update some: (old obj) --patch--> new obj ---> save

type UpdateMode int

const (
	UPDATE_MODE_PUT UpdateMode = iota
	UPDATE_MODE_PATCH
)

type Pager struct {
	PageSize   int
	PageNumber int
}
