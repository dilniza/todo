package handler

import (
	"net/http"
	"todo/api/models"

	"github.com/gin-gonic/gin"
)

// CreateLabel godoc
// @Security ApiKeyAuth
// @Router		/label [POST]
// @Summary		create label
// @Description This api creates label and returns label
// @Tags		label
// @Accept		json
// @Produce		json
// @Param		label body models.CreateLabelRequest true "Label data"
// @Success		201  {object}  models.Label
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) CreateLabel(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	req := models.CreateLabelRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}
	req.UserID = authInfo.UserID

	label, err := h.Services.Label().CreateLabel(c.Request.Context(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating label", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, label)
}

// GetLabel godoc
// @Security ApiKeyAuth
// @Router		/label/{id} [GET]
// @Summary		get label by id
// @Description This api gets label by its id and returns label
// @Tags		label
// @Accept		json
// @Produce		json
// @Param		id path string true "Label ID"
// @Success		200  {object}  models.Label
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetLabel(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	labelID := c.Param("id")
	if labelID == "" {
		handleResponseLog(c, h.Log, "label id is empty", http.StatusBadRequest, "")
		return
	}

	label, err := h.Services.Label().GetLabel(c.Request.Context(), labelID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting label", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, label)
}

// UpdateLabel godoc
// @Security ApiKeyAuth
// @Router		/label/{id} [PATCH]
// @Summary		update label by id
// @Description This api updates label by its id and returns label
// @Tags		label
// @Accept		json
// @Produce		json
// @Param		id path string true "Label ID"
// @Param		label body models.UpdateLabelRequest true "Label data"
// @Success		200  {object}  models.Label
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) UpdateLabel(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	labelID := c.Param("id")
	if labelID == "" {
		handleResponseLog(c, h.Log, "label id is empty", http.StatusBadRequest, "")
		return
	}

	updateReq := models.UpdateLabelRequest{}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != updateReq.UserID {
		handleResponseLog(c, h.Log, "user can't update another user's label", http.StatusUnauthorized, "")
		return
	}

	label, err := h.Services.Label().UpdateLabel(c.Request.Context(), updateReq)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating label", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, label)
}

// DeleteLabel godoc
// @Security ApiKeyAuth
// @Router		/label/{id} [DELETE]
// @Summary		delete label by id
// @Description This api deletes label by its id and returns message
// @Tags		label
// @Accept		json
// @Produce		json
// @Param		id path string true "Label ID"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) DeleteLabel(c *gin.Context) {
	_, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth info", http.StatusUnauthorized, err.Error())
		return
	}

	labelID := c.Param("id")
	if labelID == "" {
		handleResponseLog(c, h.Log, "label id is empty", http.StatusBadRequest, "")
		return
	}

	msg, err := h.Services.Label().DeleteLabel(c.Request.Context(), labelID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting label", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "label was successfully deleted", http.StatusOK, msg)
}

// GetAllLabels godoc
// @Security ApiKeyAuth
// @Router		/label [GET]
// @Summary		get all labels
// @Description This api gets all labels and returns label list
// @Tags		label
// @Accept		json
// @Produce		json
// @Param		search query string false "Search by name"
// @Param		page   query int   false "Page Number"
// @Param		limit  query int   false "Limit"
// @Success		200  {object}  models.GetAllLabelsResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetAllLabels(c *gin.Context) {
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

	allLabels, err := h.Services.Label().GetAllLabels(c.Request.Context(), authInfo.UserID, search, page, limit)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting labels", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, allLabels)

}
