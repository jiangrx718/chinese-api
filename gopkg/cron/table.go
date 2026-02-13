package cron

import (
	rxCron "crm/gopkg/cron/base"
	"crm/gopkg/log"
)

type TableStatus struct {
}

func NewTableStatus() rxCron.Cron {
	return &TableStatus{}
}

func (ts *TableStatus) Spec() string {
	return "* * * * *"
}

func (ts *TableStatus) Run() {
	log.Sugar().Info("每分钟执行任务")
	// 执行处理业务逻辑
}
