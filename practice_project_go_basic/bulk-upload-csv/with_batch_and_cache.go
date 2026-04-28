package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// // Cấu trúc Cache
// type MasterCache struct {
// 	Products   map[string]int
// 	Categories map[string]int
// 	Warehouses map[string]int
// }

// // Cấu trúc lỗi chi tiết
// type RowError struct {
// 	Row     int    `json:"row"`
// 	Field   string `json:"field"`
// 	Value   string `json:"value"`
// 	Message string `json:"message"`
// }

// func with_batch_and_cache_main() {
// 	db, err := sql.Open("mysql", "admin:user_password_123@tcp(127.0.0.1:3306)/ban_hang_db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.SetMaxOpenConns(50)

// 	r := gin.Default()

// 	r.POST("/api/inventory/upload-csv-fast", func(c *gin.Context) {
// 		// 1. Khởi tạo Cache (Lookup ID cực nhanh)
// 		cache := loadMasterCache(db)

// 		fileHeader, _ := c.FormFile("file")
// 		f, _ := fileHeader.Open()
// 		defer f.Close()

// 		reader := csv.NewReader(f)

// 		// Đọc Header
// 		header, err := reader.Read()
// 		if err != nil || !validateHeader(header) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Header không hợp lệ"})
// 			return
// 		}

// 		var (
// 			allErrors []RowError
// 			batch     [][]interface{}
// 			// mu        sync.Mutex
// 			rowIdx    = 1
// 			batchSize = 1000
// 		)

// 		for {
// 			rowIdx++
// 			record, err := reader.Read()
// 			if err == io.EOF {
// 				break
// 			}

// 			// 2. Validate & Lookup
// 			lineErrors, pID, cID, wID, qty := validateRow(rowIdx, record, cache)

// 			if len(lineErrors) > 0 {
// 				allErrors = append(allErrors, lineErrors...)
// 				continue
// 			}

// 			// 3. Gom vào Batch
// 			batch = append(batch, []interface{}{pID, cID, wID, qty, record[4]})

// 			// 4. Thực thi Batch Insert khi đủ số lượng
// 			if len(batch) >= batchSize {
// 				executeBatch(db, batch)
// 				batch = nil
// 			}
// 		}

// 		// Insert nốt số record còn lại
// 		if len(batch) > 0 {
// 			executeBatch(db, batch)
// 		}

// 		status := http.StatusOK
// 		if len(allErrors) > 0 {
// 			status = http.StatusMultiStatus
// 		}

// 		c.JSON(status, gin.H{
// 			"status": "completed",
// 			"errors": allErrors,
// 		})
// 	})

// 	r.Run(":8082")
// }

// // --- HÀM HỖ TRỢ (HELPERS) ---

// func loadMasterCache(db *sql.DB) MasterCache {
// 	cache := MasterCache{
// 		Products:   make(map[string]int),
// 		Categories: make(map[string]int),
// 		Warehouses: make(map[string]int),
// 	}
// 	// Tải sản phẩm
// 	rows, _ := db.Query("SELECT id, sku FROM products")
// 	for rows.Next() {
// 		var id int
// 		var code string
// 		rows.Scan(&id, &code)
// 		cache.Products[code] = id
// 	}
// 	// Tải danh mục
// 	rows, _ = db.Query("SELECT id, code FROM categories")
// 	for rows.Next() {
// 		var id int
// 		var code string
// 		rows.Scan(&id, &code)
// 		cache.Categories[code] = id
// 	}
// 	// Tải kho
// 	rows, _ = db.Query("SELECT id, code FROM warehouses")
// 	for rows.Next() {
// 		var id int
// 		var code string
// 		rows.Scan(&id, &code)
// 		cache.Warehouses[code] = id
// 	}
// 	return cache
// }

// func validateRow(idx int, rec []string, cache MasterCache) ([]RowError, int, int, int, int) {
// 	var errs []RowError
// 	pID, pOk := cache.Products[rec[0]]
// 	if !pOk {
// 		errs = append(errs, RowError{idx, "product_sku", rec[0], "product not found"})
// 	}

// 	cID, cOk := cache.Categories[rec[1]]
// 	if !cOk {
// 		errs = append(errs, RowError{idx, "category_code", rec[1], "category not found"})
// 	}

// 	wID, wOk := cache.Warehouses[rec[2]]
// 	if !wOk {
// 		errs = append(errs, RowError{idx, "warehouse_code", rec[2], "warehouse not found"})
// 	}

// 	qty, _ := strconv.Atoi(rec[3])
// 	tType := rec[4]

// 	if tType == "IN" && qty <= 0 {
// 		errs = append(errs, RowError{idx, "quantity", rec[3], "quantity must be > 0 for IN transaction"})
// 	}
// 	if tType == "OUT" && qty >= 0 {
// 		errs = append(errs, RowError{idx, "quantity", rec[3], "quantity must be < 0 for OUT transaction"})
// 	}

// 	return errs, pID, cID, wID, qty
// }

// func executeBatch(db *sql.DB, batch [][]interface{}) {
// 	if len(batch) == 0 {
// 		return
// 	}

// 	valueStrings := make([]string, 0, len(batch))
// 	valueArgs := make([]interface{}, 0, len(batch)*5)

// 	for _, row := range batch {
// 		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
// 		valueArgs = append(valueArgs, row...)
// 	}

// 	stmt := fmt.Sprintf("INSERT INTO inventory_transactions (product_id, category_id, warehouse_id, quantity, transaction_type) VALUES %s",
// 		strings.Join(valueStrings, ","))

// 	_, err := db.Exec(stmt, valueArgs...)
// 	if err != nil {
// 		fmt.Println("Lỗi Batch Insert:", err)
// 	}
// }

// func validateHeader(h []string) bool {
// 	expected := []string{"product_sku", "category_code", "warehouse_code", "quantity", "transaction_type"}
// 	if len(h) != len(expected) {
// 		return false
// 	}
// 	for i := range h {
// 		if h[i] != expected[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
