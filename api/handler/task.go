package handler

import (
	"net/http"
	"todo/api/models"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Security ApiKeyAuth
// @Router		/task [POST]
// @Summary		create task
// @Description This api creates task and returns task
// @Tags		task
// @Accept		json
// @Produce		json
// @Param		task body models.CreateTask true "Task data"
// @Success		201  {object}  models.Task
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) CreateTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	req := models.CreateTask{}
	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}
	req.UserID = authInfo.UserID

	task, err := h.Services.Task().CreateTask(c.Request.Context(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating task", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, task)

}

// GetTask godoc
// @Security ApiKeyAuth
// @Router		/task/{id} [GET]
// @Summary		get task by id
// @Description This api gets task by its id and returns task
// @Tags		task
// @Accept		json
// @Produce		json
// @Param		id path string true "Task ID"
// @Success		200  {object}  models.Task
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetTask(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskID := c.Param("id")
	if taskID == "" {
		handleResponseLog(c, h.Log, "task id is empty", http.StatusBadRequest, "")
		return
	}

	task, err := h.Services.Task().GetTask(c.Request.Context(), taskID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting task", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, task)
}

// UpdateTask godoc
// @Security ApiKeyAuth
// @Router		/task/{id} [PATCH]
// @Summary		update task by id
// @Description This api updates task by its id and returns task
// @Tags		task
// @Accept		json
// @Produce		json
// @Param		id path string true "Task ID"
// @Param		task body models.UpdateTask true "Task data"
// @Success		200  {object}  models.Task
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) UpdateTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskID := c.Param("id")
	if taskID == "" {
		handleResponseLog(c, h.Log, "task id is empty", http.StatusBadRequest, "")
		return
	}

	updateReq := models.UpdateTask{}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != updateReq.UserID {
		handleResponseLog(c, h.Log, "user can't update another user's task", http.StatusUnauthorized, "")
		return
	}

	task, err := h.Services.Task().UpdateTask(c.Request.Context(), updateReq)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating task", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Succes", http.StatusOK, task)
}

// DeleteTask godoc
// @Security ApiKeyAuth
// @Router		/task/{id} [DELETE]
// @Summary		delete task by id
// @Description This api deletes task by its id and returns message
// @Tags		task
// @Accept		json
// @Produce		json
// @Param		id path string true "Task ID"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) DeleteTask(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskID := c.Param("id")
	if taskID == "" {
		handleResponseLog(c, h.Log, "task id is empty", http.StatusBadRequest, "")
		return
	}

	msg, err := h.Services.Task().DeleteTask(c.Request.Context(), taskID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting task", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "task was successfully deleted", http.StatusOK, msg)
}

