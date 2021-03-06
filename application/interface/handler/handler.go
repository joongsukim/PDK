package handler

import (
	"net/http"

	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/KumKeeHyun/PDK/application/interface/presenter"
	"github.com/KumKeeHyun/PDK/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	nu usecase.NodeUsecase
	su usecase.SensorUsecase
}

func NewHandler(nu usecase.NodeUsecase, su usecase.SensorUsecase) *Handler {
	return &Handler{
		nu: nu,
		su: su,
	}
}

func (h *Handler) GetAllInfo(c *gin.Context) {
	nodes, err := h.nu.GetAllNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *Handler) RegisterNode(c *gin.Context) {
	var node presenter.Node

	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new, err := h.nu.RegisterNode(&node)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *new)
}

func (h *Handler) DeleteNode(c *gin.Context) {
	var node presenter.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dn, err := h.nu.DeleteNode(&node)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *dn)
}

func (h *Handler) GetSensorsInfo(c *gin.Context) {
	sensors, err := h.su.GetAllSensors()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensors)
}

func (h *Handler) RegisterSensor(c *gin.Context) {
	var sensor model.Sensor

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new, err := h.su.RegisterSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *new)
}

func (h *Handler) DeleteSensor(c *gin.Context) {
	var sensor model.Sensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ds, err := h.su.DeleteSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *ds)
}

func (h *Handler) RegisterInfo(c *gin.Context) {
	nodeInfo, _ := h.nu.GetRegister()
	sensorInfo, _ := h.su.GetRegister()
	msg := map[string]interface{}{
		"node_info":   nodeInfo,
		"sensor_info": sensorInfo,
	}
	c.JSON(http.StatusOK, msg)
}
