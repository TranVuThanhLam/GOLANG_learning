package prices

import (
	"fmt"

	"example.com/practice/convertion"
	"example.com/practice/filemanager"
)

type mapstring map[string]string

type TaxIncludedPriceJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	InputPrice       []float64               `json:"input_price"`
	TaxIncludedPrice map[string]string       `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData(filePath string) {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := convertion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrice = prices
}

func (job *TaxIncludedPriceJob) Process(filePath string) {
	job.LoadData(filePath)
	result := make(mapstring)
	for _, price := range job.InputPrice {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrice = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(f filemanager.FileManager, taxRace float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:  f,
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRace,
	}
}
