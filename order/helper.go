package order

import (
	"fmt"
	"restaurant-app/item"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

type Report struct {
	CreatedAt  time.Time
	TotalPrice float64
}

func (oa *OrderAuth) CreateReceipt(o Order, oi []OrderItem) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "Struk Pembelian", "0", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "No. Pesanan", "0", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, fmt.Sprint(o.ID), "0", 1, "C", false, 0, "")
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Tanggal Transaksi:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, fmt.Sprint(o.CreatedAt.Format("2006-01-02")), "0", 1, "C", false, 0, "")

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(85, 5, "Nama Item", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 5, "Jumlah", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Harga", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Total Harga", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	for _, v := range oi {
		item := item.Item{Id: v.ItemId}
		oa.DB.First(&item)
		pdf.CellFormat(85, 5, item.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 5, fmt.Sprintf("%d", v.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(item.Price)+",00"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(float64(v.Quantity)*item.Price)+",00"), "1", 1, "C", false, 0, "")
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Total Pembayaran:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(o.TotalPrice)+",00"), "0", 1, "C", false, 0, "")

	pdf.OutputFileAndClose("receipt/receipt-" + fmt.Sprint(o.ID) + ".pdf")
}

func (oa *OrderAuth) CreateReport(week, month int) string {
	orders := []Report{}
	if week == 0 {
		oa.DB.Raw("SELECT DATE(created_at) CreatedAt, SUM(total_price) TotalPrice FROM orders o WHERE MONTH(created_at) = ? GROUP BY DATE(created_at)", month).Scan(&orders)
	} else if month == 0 {
		oa.DB.Raw("SELECT DATE(created_at) CreatedAt, SUM(total_price) TotalPrice FROM orders o WHERE WEEK(created_at) = ? GROUP BY DATE(created_at)", week).Scan(&orders)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "Laporan Penghasilan", "0", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	if week == 0 {
		pdf.CellFormat(40, 5, "Periode Bulan:", "0", 0, "L", false, 0, "")
		pdf.CellFormat(50, 5, fmt.Sprint(month, " (", time.Month(month).String(), ")"), "0", 1, "C", false, 0, "")
	} else if month == 0 {
		pdf.CellFormat(40, 5, "Periode Minggu:", "0", 0, "L", false, 0, "")
		pdf.CellFormat(50, 5, fmt.Sprint(week), "0", 1, "C", false, 0, "")
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(95, 5, "Tanggal", "1", 0, "C", false, 0, "")
	pdf.CellFormat(90, 5, "Jumlah Penghasilan", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	income := 0.0
	for _, o := range orders {
		pdf.CellFormat(95, 5, fmt.Sprint(o.CreatedAt.Format("2006/01/02")), "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(o.TotalPrice)+",00"), "1", 1, "C", false, 0, "")
		income += o.TotalPrice
	}

	pdf.Ln(5)
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Total Penghasilan:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(income)+",00"), "0", 1, "C", false, 0, "")

	if month == 0 {
		pdf.OutputFileAndClose("report/week-" + fmt.Sprint(week) + ".pdf")
		return "report/week-" + fmt.Sprint(week) + ".pdf"
	} else if week == 0 {
		pdf.OutputFileAndClose("report/month-" + fmt.Sprint(month) + ".pdf")
		return "report/month-" + fmt.Sprint(month) + ".pdf"
	}
	return ""
}
