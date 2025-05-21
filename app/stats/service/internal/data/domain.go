package data

import (
	v1 "agents/api/user/service/v1"
	"agents/app/stats/service/internal/biz"
	"context"
)

type domainRepo struct {
	data *Data
}

// Get implements biz.DomainRepo.
func (d *domainRepo) Get(ctx context.Context, domain string) (*biz.User, error) {
	reply, err := d.data.uc.GetUserByDomain(ctx, &v1.GetUserByDomainRequest{
		Domain: domain,
	})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id: reply.Id,
	}, nil
}

func NewDomainRepo(data *Data) biz.DomainRepo {
	return &domainRepo{
		data: data,
	}
}
