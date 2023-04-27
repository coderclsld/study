package main

import (
	"fmt"
	. "read-excel/data"
	mysql "read-excel/pkg/mysql"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

var host = "root:123456@tcp(192.168.217.129:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

func main() {

	// var res = BusinessKnowledge{}
	db = mysql.NewMysqlConnect(host)
	// db.First(&res)
	// fmt.Println(res)
	excel := ReadExcel("C://Users/zxc80/Documents/testing_data.xlsx")
	// fmt.Println(excel)
	// processExecl(db, excel)
	reserveExcelToDb(excel)
}

func ReadExcel(filename string) [][]string {
	ret := [][]string{}
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println("read excel error", err.Error())
		return ret
	}
	sheets := f.GetSheetMap()
	// fmt.Println(sheets)
	sheet0 := sheets[2]
	// fmt.Println("first one tab", sheet0)
	rows, _ := f.GetRows(sheet0)
	return rows
}

func processExecl(db *gorm.DB, arr [][]string) {
	var first, second, third = 0, 0, 0
	// var first = 0
	for i := 0; i < len(arr); i++ {
		if arr[i][0] != "" {
			//insert first level
			bus := &BusinessKnowledge{Business_pid: 0, Name: arr[i][0]}
			if err := db.Create(&bus).Error; err != nil {
				fmt.Println("插入失败", err)
				return
			}
			for m1 := i + 1; m1 < len(arr); m1++ {
				//find first index
				if arr[m1][0] != "" {
					first = m1
					for j := i; j < first; j++ {
						//insert second level
						if arr[j][1] != "" {
							bus1 := &BusinessKnowledge{Business_pid: bus.Business_id, Name: arr[j][1]}
							if err := db.Create(&bus1).Error; err != nil {
								fmt.Println("插入失败", err)
								return
							}
							//find second index
							for m2 := j + 1; m2 < first; m2++ {
								if arr[m2][1] != "" {
									second = m2
									for k := j; k < second; k++ {
										//insert third level
										if arr[k][2] != "" {
											bus2 := &BusinessKnowledge{Business_pid: bus1.Business_id, Name: arr[k][2]}
											if err := db.Create(&bus2).Error; err != nil {
												fmt.Println("插入失败", err)
												return
											}
											for m3 := k + 1; m3 < second; m3++ {
												//find thrid index
												if arr[m3][2] != "" {
													third = m3
													for l := k; l < third; l++ {
														//insert four level
														bus3 := &BusinessKnowledge{Business_pid: bus2.Business_id, Name: arr[l][3]}
														if err := db.Create(&bus3).Error; err != nil {
															fmt.Println("插入失败", err)
															return
														}
														// m3 = third
													}
												}
											}
										}
										// m2 = second
									}
								}
							}
						}
						// m1 = first
					}
				}
			}
		}
	}
}

func reserveExcelToDb(arr [][]string) {
	for i := 0; i < len(arr); i++ {
		if arr[i][0] != "" {
			Len := searchLen(arr, i, 0)
			if Len != -1 {
				deep(arr, i, 0, Len, 0)
			}
		}
	}
}

func deep(arr [][]string, x, y int, Len int, pid int) {
	bus := &BusinessKnowledge{
		Business_pid: pid,
		Name:         arr[x][y],
	}
	if err := db.Create(&bus).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	if y == 3 || x == len(arr)-1 {
		return
	}
	for i := x; x < Len; i++ {
		if arr[i][y+1] != "" {
			for j := x + 1; j < Len; j++ {
				if arr[j][y+1] != "" {
					deep(arr, i, y+1, j, bus.Business_id)
				}
			}
		}
	}
}

func searchLen(arr [][]string, x, y int) int {
	for i := x; i < len(arr); i++ {
		if arr[i][y] != "" {
			return i
		}
	}
	return -1
}
