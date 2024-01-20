package interfaces

import (
	"context"

	req "github.com/kannan112/token-application/pkg/api/handler/request"
	"github.com/kannan112/token-application/pkg/domain"
)

type UserUseCase interface {
	GetTicket(ctx context.Context, id uint) (domain.Ticket, error)
	CreateClinic(ctx context.Context, data req.Clinic) (uint, error)
	CreateTicket(ctx context.Context, id uint) (uint, error)
}
