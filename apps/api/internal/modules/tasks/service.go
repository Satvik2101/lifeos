package tasks

import (
	"time"

	"github.com/google/uuid"

	api "lifeos-api/internal/server/api"
)

type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) resolveTags(ids []string) []api.Tag {
	tags := make([]api.Tag, 0, len(ids))

	for _, id := range ids {
		tags = append(tags, api.Tag{
			Id:   id,
			Name: "placeholder",
		})
	}

	return tags
}

func (s *TaskService) Create(req api.CreateTaskRequest) api.Task {
	now := time.Now()
	tags := []api.Tag{}
	if req.TagIds != nil {
		tags = s.resolveTags(*req.TagIds)
	}
	task := api.Task{
		Id:        uuid.NewString(),
		Title:     req.Title,
		Notes:     req.Notes,
		Status:    api.Inbox,
		ProjectId: req.ProjectId,
		Tags:      tags,
		DueAt:     req.DueAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.repo.Create(task)
	return task
}

func (s *TaskService) List() []api.Task {
	return s.repo.List()
}

func (s *TaskService) Get(id string) (api.Task, error) {
	task, ok := s.repo.Get(id)

	if !ok {
		return api.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (s *TaskService) Update(id string, req api.UpdateTaskRequest) (api.Task, error) {
	task, err := s.Get(id)

	if err != nil {
		return api.Task{}, err
	}

	if req.Title != nil {
		task.Title = *req.Title
	}

	if req.Notes != nil {
		task.Notes = req.Notes
	}

	if req.Status != nil {
		task.Status = *req.Status
	}

	if req.ProjectId != nil {
		task.ProjectId = req.ProjectId
	}

	if req.TagIds != nil {
		task.Tags = s.resolveTags(*req.TagIds)
	}

	task.UpdatedAt = time.Now()

	s.repo.Update(task)
	return task, nil
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}
