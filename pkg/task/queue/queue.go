package queue

import (
	"github.com/skygeario/skygear-server/pkg/db"
	"github.com/skygeario/skygear-server/pkg/task"
)

type CaptureTaskContext func() *task.Context

type Executor interface {
	Submit(ctx *task.Context, task task.Spec)
}

type Queue struct {
	TxContext      db.Context
	CaptureContext CaptureTaskContext
	Executor       Executor

	pendingTasks []task.Spec
	hooked       bool
}

func (s *Queue) Enqueue(spec task.Spec) {
	if s.TxContext != nil && s.TxContext.HasTx() {
		s.pendingTasks = append(s.pendingTasks, spec)
		if !s.hooked {
			s.TxContext.UseHook(s)
			s.hooked = true
		}
	} else {
		// No transaction context -> submit immediately.
		s.submit(spec)
	}
}

func (s *Queue) WillCommitTx() error {
	return nil
}

func (s *Queue) DidCommitTx() {
	for _, task := range s.pendingTasks {
		s.submit(task)
	}
	s.pendingTasks = nil
}

func (s *Queue) submit(spec task.Spec) {
	s.Executor.Submit(s.CaptureContext(), spec)
}