package middleware

import (
	"encoding/json"
	"errors"
	"raffalda-api/internal/infrastructure/controllers/safeobject"
	"raffalda-api/pkg/advancedlog"

	"github.com/gofiber/fiber/v2"
)

var invalidPolicy = errors.New("invalid policy")

func (m *middleware) GetPolicy(ctx *fiber.Ctx) (*safeobject.Policy, error) {
	logF := advancedlog.FunctionLog(m.log)

	policyH := ctx.GetRespHeader("policy")

	var policy *safeobject.Policy
	err := json.Unmarshal([]byte(policyH), &policy)
	if err != nil {
		logF.Errorln(err)
		return nil, invalidPolicy
	}

	return policy, nil
}

func (m *middleware) CreatePolicyFunc() func(*fiber.Ctx) error {
	logF := advancedlog.FunctionLog(m.log)
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		policy, err := m.auth.Policy(ctx.Context(), token)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		policyData, err := json.Marshal(policy)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		ctx.Set("policy", string(policyData))

		return ctx.Next()
	}
}
