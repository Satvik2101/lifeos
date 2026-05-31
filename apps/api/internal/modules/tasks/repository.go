package tasks

import (
	"sync"

	api "lifeos-api/internal/server/api"
)

type TaskRepository struct {
	mu    sync.Mutex
	tasks map[string]api.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]api.Task),
	}
}

func (r *TaskRepository) Create(task api.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.Id] = task
}

func (r *TaskRepository) List() []api.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]api.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		out = append(out, t)
	}
	return out
}

func (r *TaskRepository) Get(id string) (api.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.tasks[id]
	return t, ok
}

func (r *TaskRepository) Update(updated api.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[updated.Id]; !ok {
		return ErrTaskNotFound
	}

	r.tasks[updated.Id] = updated

	return nil
}

func (r *TaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)

	return nil
}
