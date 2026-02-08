package models

type DailyReportResponse struct {
	TotalRevenue    int            `json:"total_revenue"`
	TotalTransaksi  int            `json:"total_transaksi"`
	ProdukTerlaris  BestProductDTO `json:"produk_terlaris"`
}

type BestProductDTO struct {
	Nama        string `json:"nama"`
	QtyTerjual  int    `json:"qty_terjual"`
}
