package service

import (
	"example.com/domain"
	"golang.org/x/crypto/bcrypt"
)

type DriverService struct {
    repo domain.DriverRepository
}

func NewDriverService(repo domain.DriverRepository) *DriverService {
    return &DriverService{repo}
}

func (s *DriverService) RegisterDriver(d *domain.Driver) error {
    // Generate a hashed password with a default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    d.Password = string(hashedPassword)
    
    return s.repo.Create(d)
}

func (s *DriverService) GetDriver(id uint) (*domain.Driver, error) { return s.repo.GetByID(id) }
func (s *DriverService) RemoveDriver(id uint) error              { return s.repo.Delete(id) }

func (s *DriverService) GetAllDrivers() ([]domain.Driver, error) {
    return s.repo.List()
}

func (s *DriverService) UpdateDriver(id uint, data map[string]interface{}) error {
    // If the update includes a password, hash it here!
    if pwd, ok := data["password"].(string); ok && pwd != "" {
        hashed, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
        data["password"] = string(hashed)
    }
    return s.repo.Update(id, data)
}

