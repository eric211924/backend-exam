package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ==================== Item Interface & Items ====================

type Item interface {
	Process(workerID int)
	TypeName() string
}

type Item1 struct{}
type Item2 struct{}
type Item3 struct{}

func (i Item1) Process(workerID int) {
	fmt.Printf("[員工 %d] 開始處理 Item1...\n", workerID)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("[員工 %d] 結束處理 Item1!\n", workerID)
}

func (i Item1) TypeName() string { return "Item1" }

func (i Item2) Process(workerID int) {
	fmt.Printf("[員工 %d] 開始處理 Item2...\n", workerID)
	time.Sleep(350 * time.Millisecond)
	fmt.Printf("[員工 %d] 結束處理 Item2!\n", workerID)
}

func (i Item2) TypeName() string { return "Item2" }

func (i Item3) Process(workerID int) {
	fmt.Printf("[員工 %d] 開始處理 Item3...\n", workerID)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[員工 %d] 結束處理 Item3!\n", workerID)
}

func (i Item3) TypeName() string { return "Item3" }

// ==================== Employee ====================

type Employee struct {
	ID    int
	Count int
}

// ==================== Main ====================

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	const workerCount = 5
	const itemsPerType = 10

	// 建立物品（30 個）
	var items []Item
	for i := 0; i < itemsPerType; i++ {
		items = append(items, Item1{}, Item2{}, Item3{})
	}

	// 隨機打亂順序
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	// channel 當作流水線
	itemChan := make(chan Item)

	// 員工統計
	employees := make([]*Employee, workerCount)
	for i := range employees {
		employees[i] = &Employee{ID: i + 1}
	}

	var wg sync.WaitGroup

	// 啟動 5 名員工
	for _, emp := range employees {
		wg.Add(1)
		go func(e *Employee) {
			defer wg.Done()
			for item := range itemChan {
				item.Process(e.ID)
				e.Count++
			}
		}(emp)
	}

	// 投放物品
	for _, item := range items {
		itemChan <- item
	}
	close(itemChan)

	// 等待所有員工完成
	wg.Wait()

	// 統計
	totalTime := time.Since(startTime)

	fmt.Println("\n=========== 統計結果 ===========")
	fmt.Printf("總處理時間：%v\n", totalTime)

	for _, emp := range employees {
		fmt.Printf("員工 %d 共處理 %d 件物品\n", emp.ID, emp.Count)
	}
}
