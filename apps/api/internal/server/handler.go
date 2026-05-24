package server

import (
	api "lifeos-api/internal/server/api"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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

}

// (POST /tasks)
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

}

// (DELETE /tasks/{taskId})
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {

}

// (GET /tasks/{taskId})
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {

}

// (PATCH /tasks/{taskId})
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request, taskId api.TaskId) {

}
