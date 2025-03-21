package main

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

func generatePDF(data WorkerData) (*bytes.Reader, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "O'TKAZMA KVITANSIYASI")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	addRow := func(label, value string) {
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(90, 8, label, "0", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(90, 8, value, "0", 1, "R", false, 0, "")
	}

	addRow("Sender Name:", data.transsaction.SenderName)
	addRow("Sender:", data.transsaction.SenderCard)
	pdf.Ln(5)

	addRow("Reciever Name:", data.transsaction.ReceiverName)
	addRow("Reciever:", data.transsaction.ReceiverCard)
	pdf.Ln(5)

	addRow("Transaction Date:", data.transsaction.TransactionDate)
	addRow("Reciept Date:", data.transsaction.ReceiptDate)
	pdf.Ln(5)

	addRow("Transaction ID:", data.transsaction.TransactionId)
	addRow("Amount:", data.transsaction.Amout)
	pdf.Ln(5)

	addRow("Commision:", data.transsaction.Commision)
	addRow("Total:", data.transsaction.Total)
	pdf.Ln(5)

	// Write PDF to the buffer

	var pdfBuff bytes.Buffer
	err := pdf.Output(&pdfBuff)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(pdfBuff.Bytes())
	return reader, nil

}
