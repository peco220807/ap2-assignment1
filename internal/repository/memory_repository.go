package repository
import (
	"errors"
	"sync"
	"appointment-service/internal/model"
)

type MemoryAppointmentRepository struct {
	data map[string]model.Appointment
	mu   sync.RWMutex
}

func NewMemoryAppointmentRepository() *MemoryAppointmentRepository {
	return &MemoryAppointmentRepository{
		data: make(map[string]model.Appointment),
	}
}

func (r *MemoryAppointmentRepository) Create(a *model.Appointment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[a.ID] = *a
	return nil
}

func (r *MemoryAppointmentRepository) GetByID(id string) (*model.Appointment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	a, ok := r.data[id]
	if !ok {
		return nil, errors.New("appointment not found")
	}
	return &a, nil
}

func (r *MemoryAppointmentRepository) GetAll() ([]model.Appointment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := make([]model.Appointment, 0, len(r.data))
	for _, v := range r.data {
		res = append(res, v)
	}
	return res, nil
}

func (r *MemoryAppointmentRepository) Update(a *model.Appointment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[a.ID] = *a
	return nil
}