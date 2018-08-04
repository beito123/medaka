package async

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

type Job struct {
	id        int
	task      Task
	canceled  bool
	completed bool
}

func (job *Job) ID() int {
	return job.id
}

func (job *Job) Task() Task {
	return job.task
}

func (job *Job) IsCanceled() bool {
	return job.canceled
}

func (job *Job) IsCompleted() bool {
	return job.completed
}

func newWorker() *Worker {
	return &Worker{
		jobs: make(chan *Job),
	}
}

type Worker struct {
	jobs    chan *Job
	working bool
	count   int
}

func (worker *Worker) AddJob(job *Job) {
	worker.jobs <- job
	worker.count++
}

func (worker *Worker) Len() int {
	ln := worker.count

	if worker.IsWorking() { //
		ln += 1
	}

	return ln
}

func (worker *Worker) IsWorking() bool {
	return worker.working
}

func (worker Worker) Process() {
	var job *Job
	var ok bool
	var shutdown bool
	for {
		select {
		case job, ok = <-worker.jobs:
			if !ok {
				shutdown = true
			}
		}

		if shutdown {
			break
		}

		worker.count--

		if job.IsCanceled() {
			continue
		}

		worker.working = true

		// need recovery?

		task := job.Task()

		task.OnRun(task.Value())

		job.completed = true

		worker.working = false
	}
}
