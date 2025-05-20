package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/dictionary-service/internal/insurance"
)

type InsuranceService interface {
	GetById(id *uuid.UUID) (*insurance.Model, error)
	GetList(filter *insurance.Filter) ([]*insurance.Model, error)
	Create(insurance *insurance.Model) (*uuid.UUID, error)
	Update(insurance *insurance.Model) error
	MarkDelete(id *uuid.UUID) error
	UnMarkDelete(id *uuid.UUID) error
}

type InsuranceHandler struct {
	insuranceService InsuranceService
}

func NewInsuranceHandler(s InsuranceService) *InsuranceHandler {
	return &InsuranceHandler{
		insuranceService: s,
	}
}

func (h *InsuranceHandler) GetById(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	insurance, err := h.insuranceService.GetById(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, insurance)
}

func (h *InsuranceHandler) GetList(c *gin.Context) {
	var filter insurance.Filter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	insurances, err := h.insuranceService.GetList(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, insurances)
}

func (h *InsuranceHandler) Create(c *gin.Context) {
	var insurance insurance.Model

	if err := c.ShouldBindJSON(&insurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	uuid, err := h.insuranceService.Create(&insurance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, uuid)
}

func (h *InsuranceHandler) Update(c *gin.Context) {
	var insurance insurance.Model

	if err := c.ShouldBindJSON(&insurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	err := h.insuranceService.Update(&insurance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance updated successfully"})
}

func (h *InsuranceHandler) MarkDelete(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	err := h.insuranceService.MarkDelete(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance deleted successfully"})
}

func (h *InsuranceHandler) UnMarkDelete(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	err := h.insuranceService.UnMarkDelete(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance un deleted successfully"})
}
