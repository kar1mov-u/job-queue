package main

func generatePDF(data WorkerData) {
	// fmt.Println("generating job with id: ", data.workID)
	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.SetMargins(10, 10, 10)
	// pdf.AddPage()

	// pdf.SetFont("Arial", "B", 14)
	// pdf.Cell(0, 10, "O'TKAZMA KVITANSIYASI")
	// pdf.Ln(15)

	// pdf.SetFont("Arial", "", 12)

	// addRow := func(label, value string) {
	// 	pdf.SetFont("Arial", "", 12)
	// 	pdf.CellFormat(90, 8, label, "0", 0, "L", false, 0, "")
	// 	pdf.SetFont("Arial", "B", 12)
	// 	pdf.CellFormat(90, 8, value, "0", 1, "R", false, 0, "")
	// }

	// addRow("Sender Name:", data.SenderName)
	// addRow("Sender:", data.SenderCard)
	// pdf.Ln(5)

	// addRow("Reciever Name:", data.ReceiverName)
	// addRow("Reciever:", data.ReceiverCard)
	// pdf.Ln(5)

	// addRow("Transaction Date:", data.TransactionDate)
	// addRow("Reciept Date:", data.ReceiptDate)
	// pdf.Ln(5)

	// addRow("Transaction ID:", data.TransactionId)
	// addRow("Amount:", data.Amout)
	// pdf.Ln(5)

	// addRow("Commision:", data.Commision)
	// addRow("Total:", data.Total)
	// pdf.Ln(5)

	// err := pdf.OutputFileAndClose(data.TransactionId + ".pdf")
	// if err != nil {
	// 	log.Print("Failed on transaction" + data.TransactionId)
	// }

}
