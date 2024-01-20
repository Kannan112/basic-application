package repository

import (
	"context"
	"time"

	req "github.com/kannan112/token-application/pkg/api/handler/request"
	"github.com/kannan112/token-application/pkg/domain"
	interfaces "github.com/kannan112/token-application/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CreateClinic(ctx context.Context, data req.Clinic) (uint, error) {
	newClinic := domain.Clinic{Name: data.Name, Location: data.Location}
	err := c.DB.Create(&newClinic).Error
	if err != nil {
		return 0, err
	}
	return newClinic.ID, nil
}

func (c *userDatabase) GetTicket(ctx context.Context, id uint) (domain.Ticket, error) {
	var data domain.Ticket
	query := `SELECT t.uhcd FROM Tickets AS t where id=$1`
	err := c.DB.Raw(query, id).Scan(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *userDatabase) CreateTicket(ctx context.Context, ClientID uint, UHID string) (uint, error) {
	ticket := domain.Ticket{
		ClinicID: ClientID,
		UHCD:     UHID,
		TimeUTC:  time.Now().UTC().Format(time.RFC3339), // Format the time as needed
	}

	result := c.DB.Create(&ticket)
	if result.Error != nil {
		return 0, result.Error
	}

	return ticket.ID, nil
}

func (c *userDatabase) CheckClinic(ctx context.Context, clinicID uint) (bool, error) {
	var count int64
	err := c.DB.Model(&domain.Clinic{}).Where("id = ?", clinicID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
