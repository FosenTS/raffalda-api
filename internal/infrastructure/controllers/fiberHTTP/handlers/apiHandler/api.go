package apihandler

import (
	"encoding/json"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerApi) RegisterGroup(g fiber.Router) {
	g.Post("/test", h.Test)
	g.Post("/parseData", h.ParseData)

	g.Get("/getWarehouses", h.GetAllWarehouse)
	g.Get("/getWarehouseById", h.GetWarehouseById)
	g.Post("/storeWarehouse", h.StoreWarehouse)

	g.Post("/storeWarehouseMerchandise", h.StoreWarehouseMerchandise)
	g.Get("/getAllMerchandiseMoreInfo", h.GetAllMerchandiseMoreInfo)
}

func (h *handlerApi) Test(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("test")
}

func (h *handlerApi) ParseData(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	datas, err := h.merchandiseParser.MerchandiseFileParse(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = h.merchandiseService.StoreBulkParse(ctx.Context(), datas)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handlerApi) GetAllMerchandiseMoreInfo(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")

	ms, err := h.warehousService.GetAllMerchandiseMoreInfo(ctx.Context(), uint(page))
	if err != nil {
		h.log.Errorln(err)
		return err
	}

	bodyMessage, err := json.Marshal(ms)
	if err != nil {
		h.log.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).Send(bodyMessage)
}

func (h *handlerApi) StoreWarehouseMerchandise(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	warehouseMerchandiseCreate := new(dto.WarehouseMerchandiseCreate)
	if err := ctx.BodyParser(warehouseMerchandiseCreate); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.warehousService.StoreWarehouseMerchandise(ctx.Context(), warehouseMerchandiseCreate)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handlerApi) StoreWarehouse(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	warehouseCreate := new(dto.WarehouseCreate)
	if err := ctx.BodyParser(warehouseCreate); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.warehousService.StoreNewWarehouse(ctx.Context(), warehouseCreate)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handlerApi) GetAllWarehouse(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	isMoreInfo := ctx.QueryBool("isMore")
	if isMoreInfo {
		warehouses, err := h.warehousService.GetAllAndMoreInfo(ctx.Context())
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		bodyResponse, err := json.Marshal(warehouses)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.Status(fiber.StatusOK).Send(bodyResponse)
	}

	warehouses, err := h.warehousService.GetAll(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	bodyResponse, err := json.Marshal(warehouses)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).Send(bodyResponse)
}

func (h *handlerApi) GetWarehouseById(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	idC := ctx.Query("id")
	id, err := strconv.Atoi(idC)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	warehouse, err := h.warehousService.GetById(ctx.Context(), uint(id))
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	bodyResponce, err := json.Marshal(warehouse)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).Send(bodyResponce)
}
