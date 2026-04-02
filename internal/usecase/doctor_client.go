package usecase

type DoctorClient interface {
	DoctorExists(id string) (bool, error)
}