package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接配置
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/port?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rootDir := "/Users/7cru/data" // 修改为你的数据根目录

	// 遍历子目录1到12
	for i := 1; i <= 12; i++ {
		dirPath := filepath.Join(rootDir, fmt.Sprintf("%d", i))
		filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
				log.Printf("Processing file: %s\n", path)
				processCSV(db, path, i)
			}
			return nil
		})
	}
}

// 处理单个CSV文件
func processCSV(db *sql.DB, filePath string, month int) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("无法打开文件:", filePath, err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Println("无法读取CSV:", filePath, err)
		return
	}

	// 跳过表头，从第二行开始处理
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		if len(row) < 4 {
			log.Printf("跳过无效数据行：%v", row)
			continue
		}

		departurePortCode := row[1]
		arrivalPortCode := row[5]

		voyageCount, err1 := strconv.Atoi(row[8])
		containerTonnage, err2 := strconv.ParseFloat(row[9], 64)

		if err1 != nil || err2 != nil {
			log.Printf("数据解析失败: %v, 错误: %v %v", row, err1, err2)
			continue
		}

		// 月份以 YYYY-MM-01 格式保存
		dateStr := fmt.Sprintf("2024-%02d-01", month)
		insertData(db, departurePortCode, arrivalPortCode, dateStr, voyageCount, containerTonnage)
	}
}

// 数据插入MySQL函数
func insertData(db *sql.DB, departurePortCode, arrivalPortCode, month string, voyageCount int, containerTonnage float64) {
	stmt, err := db.Prepare(`
		INSERT INTO port_traffic_monthly 
			(departurePortCode, arrivalPortCode, month, voyageCount, containerTonnage) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println("Prepare error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(departurePortCode, arrivalPortCode, month, voyageCount, containerTonnage)
	if err != nil {
		log.Println("Insert error:", err)
	} else {
		log.Printf("Inserted: %s -> %s [%s], voyages: %d, tonnage: %.2f",
			departurePortCode, arrivalPortCode, month, voyageCount, containerTonnage)
	}
}
