package paging

type Paging struct {
	Offset  int64
	Limit   int64
	OrderBy string
}
