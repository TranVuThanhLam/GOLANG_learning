package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const RowNumber = 1000000

func main() {
	// Cấu hình các giá trị hợp lệ
	skus := []string{"FS56", "IP15", "TS01"}
	categories := []string{"ELEC", "FURN", "FASH"}
	warehouses := []string{"DN", "HCM", "HN"}
	types := []string{"IN", "OUT"}

	// Tạo file CSV
	file, err := os.Create("inventory_transactions.csv")
	if err != nil {
		fmt.Println("Lỗi tạo file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 1. Viết Header (phải khớp chính xác với yêu cầu API)
	header := []string{"product_sku", "category_code", "warehouse_code", "quantity", "transaction_type"}
	writer.Write(header)

	// Khởi tạo seed cho random
	rand.Seed(time.Now().UnixNano())

	// 2. Tạo 1000000 dòng dữ liệu
	for i := 0; i < RowNumber; i++ {
		sku := skus[rand.Intn(len(skus))]
		cat := categories[rand.Intn(len(categories))]
		wh := warehouses[rand.Intn(len(warehouses))]
		tType := types[rand.Intn(len(types))]

		// Quantity ngẫu nhiên từ 1 đến 100
		var qty int
		if tType == "IN" {
			qty = rand.Intn(100) + 1
		} else {
			qty = -(rand.Intn(100) + 1) // OUT sẽ có quantity âm
		}

		record := []string{
			sku,
			cat,
			wh,
			strconv.Itoa(qty),
			tType,
		}
		writer.Write(record)
	}

	fmt.Printf("✅ Đã tạo thành công file inventory_transactions.csv với %d dòng dữ liệu.\n", RowNumber)
}
