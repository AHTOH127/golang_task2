package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func(taskID int) error

type TaskInfo struct {
	Id        int
	BeginTime int64
	EndTime   int64
	Duration  int64
	Error     error
}

var wg sync.WaitGroup

func ScheduleTasks(tasks []Task) []TaskInfo {
	result := make([]TaskInfo, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		go func(index int, t Task) {
			defer wg.Done()
			beginTime := time.Now()
			err := t(index)
			endTime := time.Now()
			duration := endTime.Sub(beginTime)

			result[index] = TaskInfo{
				Id:        index,
				BeginTime: beginTime.Unix(),
				EndTime:   endTime.Unix(),
				Duration:  duration.Milliseconds(),
				Error:     err,
			}
		}(i, task)
	}
	wg.Wait()
	return result
}

func main() {
	// 创建示例任务
	tasks := []Task{
		// 任务1：模拟耗时200ms
		func(taskID int) error {
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("任务 %d 执行完成\n", taskID)
			return nil
		},
		// 任务2：模拟耗时100ms
		func(taskID int) error {
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("任务 %d 执行完成\n", taskID)
			return nil
		},
		// 任务3：模拟耗时300ms
		func(taskID int) error {
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("任务 %d 执行完成\n", taskID)
			return nil
		},
		// 任务4：模拟一个可能出错的任务
		func(taskID int) error {
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(150 * time.Millisecond)
			fmt.Printf("任务 %d 执行失败\n", taskID)
			return fmt.Errorf("模拟错误：任务执行失败")
		},
	}

	// 执行任务调度
	fmt.Println("开始调度任务...")
	startTime := time.Now()
	results := ScheduleTasks(tasks)
	totalDuration := time.Since(startTime)

	// 输出执行结果
	fmt.Printf("\n所有任务处理完毕，总耗时: %v\n", totalDuration)
	fmt.Println("任务执行详情：")
	for _, result := range results {
		status := "成功"
		if result.Error != nil {
			status = fmt.Sprintf("失败: %v", result.Error)
		}
		fmt.Printf("任务 %d: 耗时 %v, 状态: %s\n",
			result.Id,
			result.Duration,
			status,
		)
	}
}
