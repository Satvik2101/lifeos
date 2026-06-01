package server

import (
	"encoding/json"
	"errors"
	"lifeos-api/internal/modules/tasks"
	api "lifeos-api/internal/server/api"
	"net/http"
)

type Handler struct {
	Tasks *tasks.TaskService
}

func NewHandler() *Handler {
	fakeTaskRepository := tasks.NewFakeTaskRepository()
	tasksService := tasks.NewTaskService(fakeTaskRepository)
	return &Handler{
		Tasks: tasksService,
	}
}

func (h *Handler) HealthCheck(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// (GET /projects)

func (h *Handler) ListProjects(w http.ResponseWriter, r *http.Request) {

}

// (POST /projects)
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
}

// (DELETE /projects/{projectId})
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request, projectId api.ProjectId) {

}

// (GET /projects/{projectId})
func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request, projectId api.ProjectId) {

}

// (PATCH /projects/{projectId})
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request, projectId api.ProjectId) {

}

// (GET /tags)
func (h *Handler) ListTags(w http.ResponseWriter, r *http.Request) {

}

// (POST /tags)
func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {

}

// (DELETE /tags/{tagId})
func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request, tagId api.TagId) {

}

// (GET /tags/{tagId})
func (h *Handler) GetTag(w http.ResponseWriter, r *http.Request, tagId api.TagId) {

}

// (PATCH /tags/{tagId})
func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request, tagId api.TagId) {

}

// (GET /tasks)
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request, params api.ListTasksParams) {

	tasks := h.Tasks.List()
	writeJSON(w, http.StatusOK, tasks)
}

// (POST /tasks)
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req api.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task := h.Tasks.Create(req)
	writeJSON(w, http.StatusCreated, task)
}

// (DELETE /tasks/{taskId})
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {

	if err := h.Tasks.Delete(taskId); err != nil {
		if errors.Is(err, tasks.ErrTaskNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

}

// (GET /tasks/{taskId})
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {
	task, err := h.Tasks.Get(taskId)

	if err != nil {
		if errors.Is(err, tasks.ErrTaskNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

// (PATCH /tasks/{taskId})
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {

	req := api.UpdateTaskRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.Tasks.Update(taskId, req)

	if err != nil {
		if errors.Is(err, tasks.ErrTaskNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, task)
}
