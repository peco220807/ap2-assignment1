package usecase

import (
	"errors"
	"time"
	"appointment-service/internal/model"
	"appointment-service/internal/repository"
	"github.com/google/uuid"
)

type AppointmentUsecase interface {
	Create(title, desc, doctorID string) (*model.Appointment, error)
	GetByID(id string) (*model.Appointment, error)
	GetAll() ([]model.Appointment, error)
	UpdateStatus(id string, status model.Status) error
}

type appointmentUsecase struct {
	repo         repository.AppointmentRepository
	doctorClient DoctorClient
}

func NewAppointmentUsecase(r repository.AppointmentRepository, d DoctorClient) AppointmentUsecase {
	return &appointmentUsecase{r, d}
}

func (u *appointmentUsecase) Create(title, desc, doctorID string) (*model.Appointment, error) {
	if title == "" {
		return nil, errors.New("title required")
	}
	if doctorID == "" {
		return nil, errors.New("doctor_id required")
	}

	exists, err := u.doctorClient.DoctorExists(doctorID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("doctor does not exist")
	}

	now := time.Now()
	a := &model.Appointment{
		ID: uuid.NewString(), Title: title, Description: desc,
		DoctorID: doctorID, Status: model.StatusNew,
		CreatedAt: now, UpdatedAt: now,
	}
	err = u.repo.Create(a)
	return a, err
}

func (u *appointmentUsecase) GetByID(id string) (*model.Appointment, error) {
	return u.repo.GetByID(id)
}

func (u *appointmentUsecase) GetAll() ([]model.Appointment, error) {
	return u.repo.GetAll()
}

func (u *appointmentUsecase) UpdateStatus(id string, status model.Status) error {
	a, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	if a.Status == model.StatusDone && status == model.StatusNew {
		return errors.New("cannot move done back to new")
	}
	a.Status = status
	a.UpdatedAt = time.Now()
	return u.repo.Update(a)
}