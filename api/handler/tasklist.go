package handler

import (
	"net/http"
	"todo/api/models"

	"github.com/gin-gonic/gin"
)

// CreateTaskList godoc
// @Security ApiKeyAuth
// @Router		/task-list [POST]
// @Summary		create task list
// @Description This api creates task list and returns task list
// @Tags		task-list
// @Accept		json
// @Produce		json
// @Param		task-list body models.CreateTaskList true "Task List data"
// @Success		201  {object}  models.TaskList
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) CreateTaskList(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	req := models.CreateTaskList{}
	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}
	req.UserID = authInfo.UserID

	taskList, err := h.Services.Tasklist().CreateTaskList(c.Request.Context(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating task list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, taskList)

}

// GetTaskList godoc
// @Security ApiKeyAuth
// @Router		/task-list/{id} [GET]
// @Summary		get task list by id
// @Description This api gets task list by its id and returns task list
// @Tags		task-list
// @Accept		json
// @Produce		json
// @Param		id path string true "Task List ID"
// @Success		200  {object}  models.GetTaskListResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetTaskList(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskListID := c.Param("id")
	if taskListID == "" {
		handleResponseLog(c, h.Log, "task list id is empty", http.StatusBadRequest, "")
		return
	}

	taskList, err := h.Services.Tasklist().GetTaskList(c.Request.Context(), taskListID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting task list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, taskList)
}

// UpdateTaskList godoc
// @Security ApiKeyAuth
// @Router		/task-list/{id} [PATCH]
// @Summary		update task list by id
// @Description This api updates task list by its id and returns task list
// @Tags		task-list
// @Accept		json
// @Produce		json
// @Param		id path string true "Task List ID"
// @Param		task-list body models.UpdateTaskList true "Task List data"
// @Success		200  {object}  models.TaskList
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) UpdateTaskList(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskListID := c.Param("id")
	if taskListID == "" {
		handleResponseLog(c, h.Log, "task list id is empty", http.StatusBadRequest, "")
		return
	}

	updateReq := models.UpdateTaskList{}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != updateReq.UserID {
		handleResponseLog(c, h.Log, "user can't update another user's task list", http.StatusUnauthorized, "")
		return
	}

	taskList, err := h.Services.Tasklist().UpdateTaskList(c.Request.Context(), updateReq)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating task list", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Succes", http.StatusOK, taskList)
}

// DeleteTaskList godoc
// @Security ApiKeyAuth
// @Router		/task-list/{id} [DELETE]
// @Summary		delete task list by id
// @Description This api deletes task list by its id and returns message
// @Tags		task-list
// @Accept		json
// @Produce		json
// @Param		id path string true "Task List ID"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) DeleteTaskList(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	taskListID := c.Param("id")
	if taskListID == "" {
		handleResponseLog(c, h.Log, "task list id is empty", http.StatusBadRequest, "")
		return
	}

	msg, err := h.Services.Tasklist().DeleteTaskList(c.Request.Context(), taskListID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting task list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "task list was successfully deleted", http.StatusOK, msg)
}

// GetAllTaskLists godoc
// @Security ApiKeyAuth
// @Router		/task-list [GET]
// @Summary		get all task lists
// @Description This api gets all task lists and returns task list
// @Tags		task-list
// @Accept		json
// @Produce		json
// @Param		search query string false "Search by name"
// @Param		page   query int   false "Page Number"
// @Param		limit  query int   false "Limit"
// @Success		200  {object}  models.GetAllTaskListsResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetAllTaskLists(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page query param", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit query param", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	allTaskLists, err := h.Services.Tasklist().GetAllTaskLists(c.Request.Context(), authInfo.UserID, search, page, limit)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting task lists", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, allTaskLists)

}
