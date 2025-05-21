package data

import (
	"agents/app/stats/service/internal/biz"
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var _ biz.StatsRepo = (*statsRepo)(nil)

type statsRepo struct {
	data *Data
}

type DomainStats struct {
	Domain         string    `db:"domain"`
	Date           time.Time `db:"date"`
	RegisterCount  *int64    `db:"register_count"`
	RechargeAmount *int64    `db:"recharge_amount"`
}

func (s *statsRepo) isTodayDomainExists(ctx context.Context, domain string) (bool, error) {
	stmt := `
		SELECT 1
		FROM domain_daily_stats
		WHERE LOWER(domain) = LOWER(?)
		AND date = CURRENT_DATE;
	`
	row := s.data.db.QueryRowContext(ctx, stmt, domain)

	var dummy int
	err := row.Scan(&dummy)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *statsRepo) createDomainRecord(ctx context.Context, stats *DomainStats) error {
	columns := []string{"domain"}
	values := []any{stats.Domain}

	if stats.RechargeAmount != nil {
		columns = append(columns, "recharge_amount")
		values = append(values, stats.RechargeAmount)
	}
	if stats.RegisterCount != nil {
		columns = append(columns, "register_count")
		values = append(values, stats.RegisterCount)
	}

	query, args := sq.Insert("domain_daily_stats").
		Columns(columns...).
		Values(values...).
		MustSql()

	_, err := s.data.db.ExecContext(ctx, query, args...)
	return err
}

// AddRecharge implements biz.StatsRepo.
func (s *statsRepo) AddRecharge(ctx context.Context, domain string, amount int64) error {
	query := `
		INSERT INTO domain_daily_stats (
			domain,
			date,
			recharge_amount
		) VALUES (
			?,
			CURRENT_DATE,
			? 
		) ON DUPLICATE KEY 
		UPDATE recharge_amount = recharge_amount + ?;
	`
	_, err := s.data.db.ExecContext(ctx, query, domain, amount, amount)
	return err
}

// AddRegister implements biz.StatsRepo.
func (s *statsRepo) AddRegister(ctx context.Context, domain string) error {
	// exists, err := s.isTodayDomainExists(ctx, domain)
	// if err != nil {
	// 	return err
	// }

	// if !exists {
	// 	var dummy int64 = 1
	// 	return s.createDomainRecord(ctx, &DomainStats{
	// 		Domain:        domain,
	// 		RegisterCount: &dummy,
	// 	})
	// }

	query := `
		INSERT INTO domain_daily_stats (
			domain, date, register_count
		) VALUES (
			?, CURRENT_DATE, 1 
		) ON DUPLICATE KEY UPDATE 
		 	register_count = register_count + 1;
	`
	_, err := s.data.db.ExecContext(ctx, query, domain)
	return err
}

func NewStatsRepo(data *Data) biz.StatsRepo {
	return &statsRepo{
		data: data,
	}
}
