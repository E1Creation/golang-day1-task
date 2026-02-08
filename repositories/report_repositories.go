package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetTodayReport() (*models.DailyReportResponse, error) {
	var report models.DailyReportResponse

	// 1. Total revenue & total transaksi hari ini
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// 2. Produk terlaris hari ini
	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) AS qty
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)

	// kalau hari ini belum ada transaksi
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.BestProductDTO{}
		return &report, nil
	}
	if err != nil {
		return nil, err
	}

	return &report, nil
}
