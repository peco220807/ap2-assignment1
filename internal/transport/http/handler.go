package http
import (
	"net/http"
	"appointment-service/internal/model"
	"appointment-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct { uc usecase.AppointmentUsecase }

func NewHandler(u usecase.AppointmentUsecase) *Handler { return &Handler{u} }

type createRequest struct { Title, Description, DoctorID string }
type statusRequest struct { Status model.Status }

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/appointments", h.Create)
	r.GET("/appointments", h.GetAll)
	r.GET("/appointments/:id", h.GetByID)
	r.PATCH("/appointments/:id/status", h.UpdateStatus)
}

func (h *Handler) Create(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"bad request"})
		return
	}
	a, err := h.uc.Create(req.Title, req.Description, req.DoctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, a)
}

func (h *Handler) GetAll(c *gin.Context) {
	res,_ := h.uc.GetAll()
	c.JSON(http.StatusOK,res)
}

func (h *Handler) GetByID(c *gin.Context) {
	a, err := h.uc.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,a)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	var req statusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"bad request"})
		return
	}
	err := h.uc.UpdateStatus(c.Param("id"), req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.Status(http.StatusOK)
}