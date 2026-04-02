package app
import (
	doctorClient "appointment-service/internal/infrastructure/http"
	httpHandler "appointment-service/internal/transport/http"
	"appointment-service/internal/repository"
	"appointment-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type App struct { router *gin.Engine }

func NewApp() *App {
	router := gin.Default()
	repo := repository.NewMemoryAppointmentRepository()
	client := doctorClient.NewDoctorHTTPClient("http://localhost:8081")
	uc := usecase.NewAppointmentUsecase(repo, client)
	handler := httpHandler.NewHandler(uc)
	handler.Register(router)
	return &App{router}
}

func (a *App) Run() error { return a.router.Run(":8082") }