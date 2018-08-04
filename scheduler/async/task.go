package async

import (
	"sync"
)

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

type Task interface {
	IsCanceled() bool
	OnRun(val interface{})
	OnCompletion()
	Value() interface{}
	SetValue(data interface{})
	Job() *Job
	SetJob(job *Job)
}

type BaseTask struct {
	Task

	job    *Job
	value  interface{}
	mValue sync.RWMutex
}

func (task *BaseTask) IsCanceled() bool {
	return task.job.IsCanceled()
}

func (task *BaseTask) OnCompletion() {
	//
}

func (task *BaseTask) Value() (val interface{}) {
	task.mValue.RLock()
	val = task.value
	task.mValue.RUnlock()

	return val
}

func (task *BaseTask) SetValue(val interface{}) {
	task.mValue.Lock()
	task.value = val
	task.mValue.Unlock()
}

func (task *BaseTask) Job() *Job {
	return task.job
}

func (task *BaseTask) SetJob(job *Job) {
	if task.job == nil || job == nil {
		task.job = job
	}
}
