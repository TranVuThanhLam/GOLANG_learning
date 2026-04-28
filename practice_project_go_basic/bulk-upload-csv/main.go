package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const batch = 3000

type MasterCache struct {
	Products, Categories, Warehouses map[string]int
}

type RowError struct {
	Row     int    `json:"row"`
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

type Job struct {
	Index  int
	Record []string
}

func main() {
	db, _ := sql.Open("mysql", "admin:user_password_123@tcp(127.0.0.1:3306)/ban_hang_db")
	// Tăng giới hạn kết nối cho đa luồng
	db.SetMaxOpenConns(100)

	r := gin.Default()
	r.POST("/api/inventory/upload-hybrid", func(c *gin.Context) {
		cache := loadMasterCache(db)
		fileHeader, _ := c.FormFile("file")
		f, _ := fileHeader.Open()
		defer f.Close()

		reader := csv.NewReader(f)
		header, _ := reader.Read()
		if !validateHeader(header) {
			c.JSON(400, gin.H{"error": "Header sai cấu trúc"})
			return
		}

		// Thiết lập kênh truyền dữ liệu và thu thập lỗi
		jobs := make(chan Job, batch)
		errorsChan := make(chan []RowError, batch)
		var wg sync.WaitGroup
		workerCount := 8 // Tùy số nhân CPU của máy Ubuntu

		// Khởi chạy Worker Pool
		for w := 0; w < workerCount; w++ {
			wg.Add(1)
			go worker(db, jobs, errorsChan, &wg, cache)
		}

		// Luồng đọc file (Producer)
		go func() {
			rowIdx := 1
			for {
				rowIdx++
				record, err := reader.Read()
				if err == io.EOF {
					break
				}
				jobs <- Job{Index: rowIdx, Record: record}
			}
			close(jobs)
		}()

		// Luồng thu thập lỗi
		var allErrors []RowError
		errorsDone := make(chan bool)
		go func() {
			for errs := range errorsChan {
				allErrors = append(allErrors, errs...)
			}
			errorsDone <- true
		}()

		wg.Wait()
		close(errorsChan)
		<-errorsDone

		c.JSON(http.StatusOK, gin.H{"status": "finished", "total_errors": len(allErrors), "errors": allErrors})
	})
	r.Run(":8083")
}

func worker(db *sql.DB, jobs <-chan Job, errChan chan<- []RowError, wg *sync.WaitGroup, cache MasterCache) {
	defer wg.Done()
	const batchSize = batch
	var batch [][]interface{}

	for j := range jobs {
		lineErrs, pID, cID, wID, qty := validateHybrid(j.Index, j.Record, cache)
		if len(lineErrs) > 0 {
			errChan <- lineErrs
			continue
		}
		batch = append(batch, []interface{}{pID, cID, wID, qty, j.Record[4]})
		if len(batch) >= batchSize {
			executeBatch(db, batch)
			batch = nil
		}
	}
	if len(batch) > 0 {
		executeBatch(db, batch)
	} // Flush dữ liệu thừa
}

func validateHybrid(idx int, rec []string, cache MasterCache) ([]RowError, int, int, int, int) {
	var errs []RowError
	pID, pOk := cache.Products[rec[0]]
	if !pOk {
		errs = append(errs, RowError{idx, "product_sku", rec[0], "product not found"})
	}
	cID, cOk := cache.Categories[rec[1]]
	if !cOk {
		errs = append(errs, RowError{idx, "category_code", rec[1], "category not found"})
	}
	wID, wOk := cache.Warehouses[rec[2]]
	if !wOk {
		errs = append(errs, RowError{idx, "warehouse_code", rec[2], "warehouse not found"})
	}

	qty, _ := strconv.Atoi(rec[3])
	tType := rec[4]
	if tType == "IN" && qty <= 0 {
		errs = append(errs, RowError{idx, "quantity", rec[3], "quantity must be > 0 for IN"})
	}
	if tType == "OUT" && qty >= 0 {
		errs = append(errs, RowError{idx, "quantity", rec[3], "quantity must be < 0 for OUT"})
	}

	return errs, pID, cID, wID, qty
}

func executeBatch(db *sql.DB, batch [][]interface{}) {
	valueStrings := make([]string, 0, len(batch))
	valueArgs := make([]interface{}, 0, len(batch)*5)
	for _, row := range batch {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, row...)
	}
	stmt := fmt.Sprintf("INSERT INTO inventory_transactions (product_id, category_id, warehouse_id, quantity, transaction_type) VALUES %s", strings.Join(valueStrings, ","))
	db.Exec(stmt, valueArgs...)
}

func loadMasterCache(db *sql.DB) MasterCache {
	c := MasterCache{make(map[string]int), make(map[string]int), make(map[string]int)}
	fetch := func(q string, m map[string]int) {
		rows, _ := db.Query(q)
		for rows.Next() {
			var id int
			var code string
			rows.Scan(&id, &code)
			m[code] = id
		}
	}
	fetch("SELECT id, sku FROM products", c.Products)
	fetch("SELECT id, code FROM categories", c.Categories)
	fetch("SELECT id, code FROM warehouses", c.Warehouses)
	return c
}

func validateHeader(h []string) bool {
	return len(h) == 5 && h[0] == "product_sku" && h[3] == "quantity"
}
