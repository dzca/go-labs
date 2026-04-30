package web

import (
	"net/http"
	"strconv"

	"example.com/domain"
	"example.com/service"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
    svc *service.DriverService
}

func NewDriverHandler(svc *service.DriverService) *DriverHandler {
    return &DriverHandler{svc}
}

func (h *DriverHandler) Create(c *gin.Context) {
    var driver domain.Driver
    if err := c.ShouldBindJSON(&driver); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.svc.RegisterDriver(&driver); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, driver)
}

func (h *DriverHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    driver, err := h.svc.GetDriver(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Driver not found"})
        return
    }
    c.JSON(http.StatusOK, driver)
}

// GET /drivers
func (h *DriverHandler) GetAll(c *gin.Context) {
    drivers, err := h.svc.GetAllDrivers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch drivers"})
        return
    }
    c.JSON(http.StatusOK, drivers)
}

// PATCH /drivers/:id
func (h *DriverHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    // We use a map to allow partial updates (e.g., only updating last_name)
    var updateData map[string]interface{}
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if err := h.svc.UpdateDriver(uint(id), updateData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Driver updated successfully"})
}

// DELETE /drivers/:id
func (h *DriverHandler) Delete(c *gin.Context) {
    // 1. Get ID from URL (/drivers/123)
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    // 2. Call Service
    if err := h.svc.RemoveDriver(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete driver"})
        return
    }

    // 3. Return Success (204 No Content is standard for DELETE)
    c.Status(http.StatusNoContent)
}
