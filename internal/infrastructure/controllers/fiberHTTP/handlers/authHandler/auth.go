package authhandler

import (
	"errors"
	"fmt"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"

	"github.com/gofiber/fiber/v2"
)

var invalidAccessToken = errors.New("invalid access token")

func (h *handlerAuth) RegisterGroup(g fiber.Router) {

	policyChecker := h.middleware.CreatePolicyFunc()

	g.Post("/login", h.Login)
	g.Post("/register", h.Register)
	g.Get("/check", h.Check)
	g.Post("/refresh", h.Refresh)

	g.Post("/fastLoginCreate", policyChecker, h.FastLoginCreate)
	g.Post("/fastLogin", h.FastLogin)
}

func (h *handlerAuth) FastLogin(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	token := ctx.Query("a")
	fmt.Println(token)
	if token == "" {
		logF.Errorln(invalidAccessToken)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	pair, err := h.authService.FastLogin(ctx.Context(), token)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(pair)

}

func (h *handlerAuth) FastLoginCreate(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	policy, err := h.middleware.GetPolicy(ctx)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	qr, err := h.authService.CreateFastLogin(ctx.Context(), policy)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Set("Content-Type", "image/png")
	return ctx.Status(fiber.StatusOK).Send(qr)
}

func (h *handlerAuth) Login(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	login := new(dto.Login)
	if err := ctx.BodyParser(login); err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid login format")
	}

	tokens, err := h.authService.Login(ctx.Context(), login)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Status(fiber.StatusOK).JSON(tokens)
}

func (h *handlerAuth) Refresh(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	token := ctx.Query("refresh")

	pairTokens, err := h.authService.Refresh(ctx.Context(), token)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Status(fiber.StatusOK).JSON(pairTokens)
}

func (h *handlerAuth) Check(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	token := ctx.Get("Authorization")

	user, err := h.authService.Check(ctx.Context(), token)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *handlerAuth) Register(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	register := new(dto.UserCreate)
	if err := ctx.BodyParser(register); err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid register format")
	}

	err := h.authService.Register(ctx.Context(), register)
	if err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("user not created")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
