package handler

import (
	"net/http"
	"strconv"

	"github.com/example/robot-manage/internal/model"
	"github.com/example/robot-manage/internal/service"
	"github.com/gin-gonic/gin"
)

type ProjectPlanHandler struct {
	service *service.ProjectPlanService
}

func NewProjectPlanHandler(service *service.ProjectPlanService) *ProjectPlanHandler {
	return &ProjectPlanHandler{service: service}
}

func (h *ProjectPlanHandler) GetAllModulesWithItems(c *gin.Context) {
	modules, err := h.service.GetAllModulesWithItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, modules)
}

func (h *ProjectPlanHandler) CreateModule(c *gin.Context) {
	var module model.ProjectModule
	if err := c.ShouldBindJSON(&module); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateModule(&module); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}

func (h *ProjectPlanHandler) UpdateModule(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var module model.ProjectModule
	if err := c.ShouldBindJSON(&module); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module.ID = uint(id)
	if err := h.service.UpdateModule(&module); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}

func (h *ProjectPlanHandler) DeleteModule(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeleteModule(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ProjectPlanHandler) CreateItem(c *gin.Context) {
	var item model.ProjectItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ProjectPlanHandler) UpdateItem(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var item model.ProjectItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := h.service.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ProjectPlanHandler) DeleteItem(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ProjectPlanHandler) BatchUpdateItems(c *gin.Context) {
	var items []model.ProjectItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.BatchUpdateItems(items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
