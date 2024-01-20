package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	req "github.com/kannan112/token-application/pkg/api/handler/request"
	res "github.com/kannan112/token-application/pkg/api/handler/response"
	services "github.com/kannan112/token-application/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary Show details of a specific ticket
// @Description Display details of a specific ticket identified by its ID
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} res.Response "Ticket details"
// @Failure 402 {object} res.Response "Error response"
// @Router /api/get-ticket/{id} [get]
func (cr *UserHandler) ShowTicket(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		res.ErrrorResponse(ctx, 402, "id params not found", err.Error())
		return
	}
	Ticket, err := cr.userUseCase.GetTicket(ctx, uint(id))
	if err != nil {
		res.ErrrorResponse(ctx, 402, "Ticket not found", err.Error())
		return
	}
	res.SuccessResponse(ctx, 200, "Your Ticket No", Ticket.UHCD)

}

// @Summary Create a new clinic
// @Description Create a new clinic with the specified attributes
// @Accept json
// @Produce json
// @Param clinic body req.Clinic true "Clinic details"
// @Success 200 {object} res.Response  "Success response"
// @Failure 400 {object} res.Response "Error response"
// @Router /api/clinic-create [post]
func (cr *UserHandler) CreateClinic(c *gin.Context) {
	var CreateClinic req.Clinic
	if err := c.Bind(&CreateClinic); err != nil {
		res.ErrrorResponse(c, 402, "failed to bind", err.Error())
		return
	}
	id, err := cr.userUseCase.CreateClinic(c, CreateClinic)
	if err != nil {
		res.ErrrorResponse(c, 400, "failed to reg clinic", err.Error())
		return
	}
	res.SuccessResponse(c, 200, "Create You Token Now", id)
}

// @Summary Create a new ticket
// @Description Create a new ticket for a specific clinic
// @Accept json
// @Produce json
// @Param id path int true "Clinic ID"
// @Success 200 {object} res.Response "Token Number"
// @Failure 400 {object} res.Response "Error response"
// @Router /api/create-ticket/{id} [post]
func (cr *UserHandler) CreateTicket(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		res.ErrrorResponse(ctx, 402, "id params not found", err.Error())
		return
	}
	Token, err := cr.userUseCase.CreateTicket(ctx, uint(id))
	if err != nil {
		res.ErrrorResponse(ctx, 400, "failed to create Token", err.Error())
		return
	}
	res.SuccessResponse(ctx, 200, "Token Number", Token)

}
