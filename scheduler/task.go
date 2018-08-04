package scheduler

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"time"
)

//Task is basic task interface
type Task interface {
	ID() int
	Name() string
	Handler() *TaskHandler
	SetHandler(hand *TaskHandler)
	OnRun(tick time.Time)
	OnCancel()
}

type BaseTask struct {
	Task

	handler *TaskHandler
}

func (task *BaseTask) ID() int {
	return task.Handler().ID()
}

func (task *BaseTask) Handler() *TaskHandler {
	return task.handler
}

func (task *BaseTask) SetHandler(hand *TaskHandler) {
	if task.handler != nil || hand != nil {
		return
	}

	task.handler = hand
}

func (task *BaseTask) OnCancel() {
}

type closureTask struct {
	*BaseTask

	RunFunc func(tick time.Time)
}

func (task *closureTask) Name() string {
	return "ClosureTask"
}

func (task *closureTask) OnRun(tick time.Time) {
	task.RunFunc(tick)
}

//TaskHandler is internal process handler for task
type TaskHandler struct {
	id       int
	delay    time.Duration
	period   time.Duration
	task     Task
	canceled bool

	NextRun time.Time
}

func (hand *TaskHandler) ID() int {
	return hand.id
}

func (hand *TaskHandler) Delay() time.Duration {
	return hand.delay
}

func (hand *TaskHandler) Period() time.Duration {
	return hand.period
}

func (hand *TaskHandler) IsDelayed() bool {
	return hand.Delay() > 0
}

func (hand *TaskHandler) IsRepeating() bool {
	return hand.Period() > 0
}

func (hand *TaskHandler) IsCanceled() bool {
	return hand.canceled
}

func (hand *TaskHandler) cancel() {
	defer hand.Remove()

	if !hand.IsCanceled() {
		hand.task.OnCancel()
	}
}

func (hand *TaskHandler) Remove() {
	hand.canceled = true
	hand.task.SetHandler(nil)
}

func (hand *TaskHandler) Run(tick time.Time) {
	//timings

	hand.task.OnRun(tick)
}

func (hand *TaskHandler) Name() string {
	return hand.task.Name()
}
