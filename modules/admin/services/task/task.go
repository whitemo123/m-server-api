package taskService

import (
	"fmt"
	"m-server-api/utils/timeTask"
	"time"

	"github.com/spf13/cast"
)

type task struct {
	TaskID      int64
	TaskName    string
	Status      bool
	Cron        string
	TaskFunc    func()
	LastRunTime *time.Time
	RunTimes    int
}

func helloTask() {
	fmt.Println("hello go task")
}

var tasks = []task{
	// {
	// 	TaskID:      1,
	// 	TaskName:    "测试定时任务（每10秒触发一次）",
	// 	Cron:        "*/10 * * * * *",
	// 	Status:      true,
	// 	TaskFunc:    helloTask,
	// 	LastRunTime: nil,
	// 	RunTimes:    0,
	// },
}

func init() {
	for _, task := range tasks {
		fmt.Println("加载定时任务：" + task.TaskName)
		err := timeTask.TimeTask.AddTask(cast.ToString(task.TaskID), task.Cron, task.TaskFunc, task.Status)
		if err != nil {
			panic(err)
		}
	}
}

func GetTaskList() []interface{} {
	var taskList []interface{}
	for _, task := range tasks {
		taskList = append(taskList, map[string]interface{}{
			"id":          task.TaskID,
			"name":        task.TaskName,
			"status":      task.Status,
			"cron":        task.Cron,
			"lastRunTime": task.LastRunTime,
			"runTimes":    task.RunTimes,
		})
	}
	return taskList
}

func SwitchTaskStatus(id int64) {
	for i, task := range tasks {
		if task.TaskID == id {
			if task.Status {
				timeTask.TimeTask.StopTask(cast.ToString(id))
				tasks[i].Status = false
			} else {
				timeTask.TimeTask.StartTask(cast.ToString(id))
				tasks[i].Status = true
			}
		}
	}
}
