package item

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func (ia *ItemAuth) CreateStockReport() string {
	items := []Item{}
	ia.DB.Raw("SELECT * FROM items").Scan(&items)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "Laporan Stock", "0", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 5, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(40, 5, "Tanggal:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(50, 5, fmt.Sprint(time.Now().Format("2006-01-02")), "0", 1, "C", false, 0, "")

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(105, 5, "Nama Item", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 5, "Stok", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 5, "Harga", "1", 1, "C", false, 0, "")
	// pdf.CellFormat(40, 5, "Total Harga", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	for _, item := range items {
		pdf.CellFormat(105, 5, item.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 5, fmt.Sprintf("%d", item.Stock), "1", 0, "C", false, 0, "")
		pdf.CellFormat(50, 5, fmt.Sprintf("Rp. %s", humanize.Commaf(item.Price)+",00"), "1", 1, "C", false, 0, "")
		// pdf.CellFormat(40, 5, "Total Harga", "1", 1, "C", false, 0, "")
	}

	pdf.OutputFileAndClose("stock/" + fmt.Sprint(time.Now().Format("2006-01-02")) + ".pdf")

	return "stock/" + fmt.Sprint(time.Now().Format("2006-01-02")) + ".pdf"
}
