package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	req "github.com/kannan112/token-application/pkg/api/handler/request"
	"github.com/kannan112/token-application/pkg/domain"
	interfaces "github.com/kannan112/token-application/pkg/repository/interface"
	services "github.com/kannan112/token-application/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) GetTicket(ctx context.Context, id uint) (domain.Ticket, error) {
	Ticket, err := c.userRepo.GetTicket(ctx, id)
	return Ticket, err
}

func (c *userUseCase) CreateClinic(ctx context.Context, data req.Clinic) (uint, error) {
	id, err := c.userRepo.CreateClinic(ctx, data)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (c *userUseCase) CreateTicket(ctx context.Context, id uint) (uint, error) {
	check, err := c.userRepo.CheckClinic(ctx, uint(id))
	if err != nil {
		return 0, err
	}
	if !check {
		return 0, errors.New("register clinic")
	}
	UHID := generateUniqueID(5)
	fmt.Println(UHID)
	uId, err := c.userRepo.CreateTicket(ctx, uint(id), UHID)
	if err != nil {
		return 0, err
	}
	return uId, nil
}

func generateUniqueID(digits int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(10000)
	max := int64(99999)
	randomID := min + rand.Int63n(max-min+1)
	return fmt.Sprintf("%0*d", digits, randomID)
}
