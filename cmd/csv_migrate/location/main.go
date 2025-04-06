package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接信息（请替换为你的数据库信息）
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/port?charset=utf8mb4")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	// 打开CSV文件
	file, err := os.Open("/Users/7cru/marine-backend/port_location.csv")
	if err != nil {
		log.Fatal("无法打开CSV文件:", err)
	}
	defer file.Close()

	// 读取CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("读取CSV失败:", err)
	}

	// 遍历CSV行，跳过第一行（表头）
	for idx, row := range records {
		if idx == 0 {
			continue
		}

		if len(row) < 4 {
			log.Printf("跳过无效行: %v\n", row)
			continue
		}

		portName := row[1]
		portCode := row[0]
		latitude, err1 := strconv.ParseFloat(row[10], 64) // 纬度
		longitude, err2 := strconv.ParseFloat(row[9], 64) // 经度

		if err1 != nil || err2 != nil {
			log.Printf("经纬度解析失败，跳过此行: %v\n", row)
			continue
		}

		err = insertPort(db, portName, portCode, latitude, longitude)
		if err != nil {
			log.Printf("插入数据库失败: %v, 错误: %v\n", row, err)
		} else {
			fmt.Printf("已插入港口: %s (%s)\n", portName, portCode)
		}
	}
}

// 插入单条港口数据到数据库
func insertPort(db *sql.DB, name, code string, lat, lon float64) error {
	stmt, err := db.Prepare(`
		INSERT INTO port (portName, portCode, latitude, longitude)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, code, lat, lon)
	return err
}
