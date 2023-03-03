package order

import (
	"fmt"
	"restaurant-app/item"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

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
