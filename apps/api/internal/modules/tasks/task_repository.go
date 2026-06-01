package tasks

import (
	"sync"

	api "lifeos-api/internal/server/api"
)

type TaskRepository interface {
	Create(task api.Task) error
	Get(id string) (api.Task, bool)
	List() []api.Task
	Update(task api.Task) error
	Delete(id string) error
}

var _ TaskRepository = (*FakeTaskRepository)(nil)

type FakeTaskRepository struct {
	mu    sync.Mutex
	tasks map[string]api.Task
}

func NewFakeTaskRepository() *FakeTaskRepository {
	return &FakeTaskRepository{
		tasks: make(map[string]api.Task),
	}
}

func (r *FakeTaskRepository) Create(task api.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.Id] = task
	return nil
}

func (r *FakeTaskRepository) List() []api.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]api.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		out = append(out, t)
	}
	return out
}

func (r *FakeTaskRepository) Get(id string) (api.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.tasks[id]
	return t, ok
}

func (r *FakeTaskRepository) Update(updated api.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[updated.Id]; !ok {
		return ErrTaskNotFound
	}

	r.tasks[updated.Id] = updated

	return nil
}

func (r *FakeTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)

	return nil
}
