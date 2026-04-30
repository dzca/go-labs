package domain

type Driver struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
    Email     string `json:"email" gorm:"unique;not null" binding:"required,email"`
    Password  string `json:"password,omitempty" gorm:"not null" binding:"required"`
}

type DriverRepository interface {
    Create(driver *Driver) error
    GetByID(id uint) (*Driver, error)
    Update(id uint, data map[string]interface{}) error
    Delete(id uint) error
    List() ([]Driver, error)
}
