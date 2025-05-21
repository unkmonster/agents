package paging

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type Paging struct {
	Offset  int64
	Limit   int64
	OrderBy string
	Sort    string
}

func (p *Paging) Make(b sq.SelectBuilder) sq.SelectBuilder {
	if p.OrderBy != "" {
		b = b.OrderBy(fmt.Sprintf("%s %s", p.OrderBy, p.Sort))
	}
	if p.Limit != 0 {
		b = b.Limit(uint64(p.Limit))
	}
	if p.Offset != 0 {
		b = b.Offset(uint64(p.Offset))
	}
	return b
}
