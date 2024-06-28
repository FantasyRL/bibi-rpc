package main

import (
	"bibi/pkg/constants"
	"github.com/cloudwego/kitex/pkg/limit"
)

type LimiterUpdater struct {
	updater limit.Updater
}

type LimitReporter interface {
	ConnOverloadReport()
	QPSOverloadReport()
}

func (lu *LimiterUpdater) Update() {
	// your logic: set new option as needed
	newOpt := &limit.Option{
		MaxConnections: constants.UpdateConnections,
		MaxQPS:         constants.UpdateQPS,
	}
	// update limit config
	lu.updater.UpdateLimit(newOpt)
	// your logic
}

func (lu *LimiterUpdater) UpdateControl(u limit.Updater) {
	u = lu.updater
}
