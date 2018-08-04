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
	"errors"
	"math"
	"sort"
	"time"

	"github.com/beito123/medaka"
)

var (
	errNotEnabled = errors.New("not enabled")
)

func NewScheduler(logger medaka.Logger) *Scheduler {
	sche := &Scheduler{
		Logger: logger,
	}

	sche.Init()

	return sche
}

// Scheduler is a simple task manager
type Scheduler struct {
	Logger medaka.Logger

	enable bool

	queue TaskHandlerSlice

	tasks map[int]*TaskHandler

	tick time.Time

	nextID int
}

func (sche *Scheduler) Init() {
	sche.nextID = 0
	sche.queue = []*TaskHandler{}
	sche.tasks = make(map[int]*TaskHandler)
}

func (sche *Scheduler) Shutdonw() {
	sche.enable = false

	sche.CancelAllTasks()
}

func (sche *Scheduler) bumpNextID() (next int) {
	next = sche.nextID
	sche.nextID = (sche.nextID % math.MaxInt32) + 1
	return next
}

func (sche *Scheduler) ScheduleFunc(f func(tick time.Time), delay time.Duration, period time.Duration) (*TaskHandler, error) {
	return sche.addTask(&closureTask{
		BaseTask: new(BaseTask),
		RunFunc:  f,
	}, -1, -1)
}

func (sche *Scheduler) ScheduleTask(task Task) (*TaskHandler, error) {
	return sche.addTask(task, -1, -1)
}

func (sche *Scheduler) ScheduleDelayedTask(task Task, delay time.Duration) (*TaskHandler, error) {
	return sche.addTask(task, delay, -1)
}

func (sche *Scheduler) ScheduleRepeatinTask(task Task, period time.Duration) (*TaskHandler, error) {
	return sche.addTask(task, -1, period)
}

func (sche *Scheduler) ScheduleDelayedRepeatinTask(task Task, delay time.Duration, period time.Duration) (*TaskHandler, error) {
	return sche.addTask(task, delay, period)
}

func (sche *Scheduler) addTask(task Task, delay time.Duration, period time.Duration) (*TaskHandler, error) {
	if !sche.enable {
		return nil, errNotEnabled
	}

	if delay <= 0 {
		delay = -1
	}

	if period < 0 {
		period = -1
	} else if period == 0 {
		period = 1
	}

	return &TaskHandler{
		id:       sche.bumpNextID(),
		delay:    delay,
		period:   period,
		task:     task,
		canceled: false,
	}, nil
}

func (sche *Scheduler) handle(hand *TaskHandler) *TaskHandler {
	nextRun := sche.tick
	if hand.IsDelayed() {
		nextRun = nextRun.Add(hand.Delay())
	}

	hand.NextRun = nextRun
	sche.tasks[hand.ID()] = hand
	sche.queue = append(sche.queue, hand)

	return hand
}

func (sche *Scheduler) CancelTask(id int) {
	hand, ok := sche.tasks[id]
	if !ok {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			sche.Logger.Warnf("Happened a problem while task %s canceling Error:%s", hand.Name(), err)
			sche.Logger.Trace(nil)
		}

		delete(sche.tasks, hand.ID())
	}()

	hand.cancel()
}

func (sche *Scheduler) CancelAllTasks() {
	for id, _ := range sche.tasks {
		sche.CancelTask(id)
	}

	sche.Init()
}

func (sche *Scheduler) removeQueue(i int) {
	if len(sche.queue) <= i {
		return
	}

	sche.queue = remove(sche.queue, i)
}

func (sche *Scheduler) Update(tick time.Time) error {
	if !sche.enable {
		return errNotEnabled
	}

	sort.Sort(sche.queue)

	for i, hand := range sche.queue {
		if tick.Before(hand.NextRun) { // if tick is earlier than hand
			continue
		}

		sche.removeQueue(i)

		if hand.IsCanceled() {
			delete(sche.tasks, hand.ID())
			continue
		}

		crashed := false

		defer func() {
			if err := recover(); err != nil {
				crashed = true
				sche.Logger.Warnf("Happened a problem while task %s running Error:%s", hand.Name(), err)
				sche.Logger.Trace(nil)
			}
		}()

		hand.Run(tick)

		if hand.IsRepeating() {
			if crashed {
				sche.Logger.Debugf("Dropping repeating task %s due to an error happened while running", hand.Name())
			} else {
				hand.NextRun = tick.Add(hand.Period())
				sche.queue = append(sche.queue, hand)
				continue
			}
		}

		hand.Remove()
		delete(sche.tasks, hand.ID())
	}

	return nil
}

// Thank you, Chris(https://stackoverflow.com/users/59198/chris)
// From StackOverflow. Questioner: Jorge Olivero(https://stackoverflow.com/users/2441725/jorge-olivero)
// Link: https://stackoverflow.com/questions/25025409/delete-element-in-a-slice
func remove(a TaskHandlerSlice, i int) TaskHandlerSlice {
	return a[:i+copy(a[i:], a[i+1:])]
}

type TaskHandlerSlice []*TaskHandler

func (s TaskHandlerSlice) Len() int {
	return len(s)
}

func (s TaskHandlerSlice) Less(i, j int) bool {
	return s[i].NextRun.Before(s[j].NextRun)
}

func (s TaskHandlerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
