package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// type RowError struct {
// 	Row     int    `json:"row"`
// 	Field   string `json:"field"`
// 	Value   string `json:"value"`
// 	Message string `json:"message"`
// }

// type Job struct {
// 	Index  int
// 	Record []string
// }

// func main() {
// 	db, _ := sql.Open("mysql", "admin:user_password_123@tcp(127.0.0.1:3306)/ban_hang_db")
// 	db.SetMaxOpenConns(20)
// 	r := gin.Default()

// 	r.POST("/api/inventory/upload-csv", func(c *gin.Context) {
// 		fileHeader, _ := c.FormFile("file")
// 		f, _ := fileHeader.Open()
// 		defer f.Close()

// 		reader := csv.NewReader(f)
// 		records, _ := reader.ReadAll()

// 		// 1. Kiểm tra Header
// 		expectedHeader := []string{"product_sku", "category_code", "warehouse_code", "quantity", "transaction_type"}
// 		if len(records) == 0 || !reflect.DeepEqual(records[0], expectedHeader) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Cấu trúc Header không đúng. Yêu cầu: product_sku, category_code, warehouse_code, quantity, transaction_type"})
// 			return
// 		}

// 		// 2. Setup Worker Pool & Error Channel
// 		numRecords := len(records) - 1
// 		jobs := make(chan Job, numRecords)
// 		errorsChan := make(chan RowError, numRecords)
// 		var wg sync.WaitGroup

// 		// Chạy 5 workers
// 		for w := 1; w <= 5; w++ {
// 			wg.Add(1)
// 			go worker(db, jobs, errorsChan, &wg)
// 		}

// 		// 3. Đẩy dữ liệu vào
// 		for i := 1; i < len(records); i++ {
// 			jobs <- Job{Index: i + 1, Record: records[i]}
// 		}
// 		close(jobs)

// 		// Đợi và đóng channel lỗi
// 		go func() {
// 			wg.Wait()
// 			close(errorsChan)
// 		}()

// 		// 4. Thu thập kết quả lỗi
// 		var allErrors []RowError
// 		for e := range errorsChan {
// 			allErrors = append(allErrors, e)
// 		}

// 		if len(allErrors) > 0 {
// 			c.JSON(http.StatusMultiStatus, gin.H{"status": "completed_with_errors", "errors": allErrors})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Tất cả dòng đã được xử lý!"})
// 		}
// 	})

// 	r.Run(":8081")
// }

// func worker(db *sql.DB, jobs <-chan Job, errorsChan chan<- RowError, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		sku, catCode, whCode := j.Record[0], j.Record[1], j.Record[2]
// 		qtyStr, tType := j.Record[3], j.Record[4]
// 		qty, _ := strconv.Atoi(qtyStr)

// 		// Kiểm tra Logic: Quantity >= 0 nếu là IN
// 		if tType == "IN" && qty < 0 {
// 			errorsChan <- RowError{j.Index, "quantity", qtyStr, "quantity must be >= 0 for IN transaction"}
// 			continue
// 		}

// 		var pID, cID, wID int
// 		if err := db.QueryRow("SELECT id FROM products WHERE sku = ?", sku).Scan(&pID); err != nil {
// 			errorsChan <- RowError{j.Index, "product_sku", sku, "product not found"}
// 			continue
// 		}
// 		if err := db.QueryRow("SELECT id FROM categories WHERE code = ?", catCode).Scan(&cID); err != nil {
// 			errorsChan <- RowError{j.Index, "category_code", catCode, "category not found"}
// 			continue
// 		}
// 		if err := db.QueryRow("SELECT id FROM warehouses WHERE code = ?", whCode).Scan(&wID); err != nil {
// 			errorsChan <- RowError{j.Index, "warehouse_code", whCode, "warehouse not found"}
// 			continue
// 		}

// 		db.Exec(`INSERT INTO inventory_transactions (product_id, category_id, warehouse_id, quantity, transaction_type) VALUES (?, ?, ?, ?, ?)`,
// 			pID, cID, wID, qty, tType)
// 	}
// }
