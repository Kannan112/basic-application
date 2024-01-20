package interfaces

import (
	"context"

	req "github.com/kannan112/token-application/pkg/api/handler/request"
	"github.com/kannan112/token-application/pkg/domain"
)

type UserRepository interface {
	GetTicket(ctx context.Context, id uint) (domain.Ticket, error)              // end result
	CreateClinic(ctx context.Context, data req.Clinic) (uint, error)            // your client id
	CreateTicket(ctx context.Context, ClientID uint, UHID string) (uint, error) // get your ticket id
	CheckClinic(ctx context.Context, clinicID uint) (bool, error)               // check clinic id
}
