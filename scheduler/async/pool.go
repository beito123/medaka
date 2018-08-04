package async

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"math"
	"sort"
)

type Workers []*Worker

func (s Workers) Len() int {
	return len(s)
}

func (s Workers) Less(i, j int) bool {
	return s[i].Len() < s[j].Len()
}

func (s Workers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func NewPool(size int) *Pool {
	if size <= 0 {
		panic("invalid size")
	}

	p := &Pool{
		jobs:    make(map[int]*Job),
		workers: make([]*Worker, size),
	}

	for i := 0; i < size; i++ {
		p.workers[i] = newWorker()

		go p.workers[i].Process()
	}

	return p
}

type Pool struct {
	jobs    map[int]*Job
	workers Workers

	nextID int
}

func (p *Pool) bumpNextID() (next int) {
	next = p.nextID
	p.nextID = (p.nextID % math.MaxInt32) + 1
	return next
}

func (p *Pool) Count() int {
	return len(p.jobs)
}

func (p *Pool) QueueTask(task Task) *Job {
	job := &Job{
		id:        p.bumpNextID(),
		task:      task,
		canceled:  false,
		completed: false,
	}

	p.jobs[job.ID()] = job

	task.SetJob(job)

	worker := p.selectWorker()

	worker.AddJob(job)

	return job
}

func (p *Pool) selectWorker() *Worker {
	if len(p.workers) == 0 {
		return nil
	}

	sort.Sort(p.workers)

	return p.workers[0]
}

func (p *Pool) CancelTask(id int) {
	job, ok := p.jobs[id]
	if !ok {
		return
	}

	job.canceled = true
}

func (p *Pool) CancelAllTask() {
	for id, _ := range p.jobs {
		p.CancelTask(id)
	}
}

func (p *Pool) Shutdown() {
	for _, worker := range p.workers {
		close(worker.jobs)
	}
}

func (p *Pool) Update() {
	for id, job := range p.jobs {
		if job.IsCanceled() || job.IsCompleted() {
			job.Task().OnCompletion()

			delete(p.jobs, id)
		}
	}
}
