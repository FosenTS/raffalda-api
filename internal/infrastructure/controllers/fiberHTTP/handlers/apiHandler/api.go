package apihandler

import (
	"raffalda-api/internal/domain/entity"
	"raffalda-api/internal/domain/storage/dto"
	"raffalda-api/pkg/advancedlog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerApi) RegisterGroup(g fiber.Router) {
	g.Get("/generalAnalyze", h.GeneralAnalyze)

	g.Post("/test", h.Test)
	g.Post("/parseData", h.ParseData)

	g.Get("/getWarehouses", h.GetAllWarehouse)
	g.Post("/updateWarehouse", h.UpdateWarehouse)
	g.Get("/getWarehouseById", h.GetWarehouseById)
	g.Post("/storeWarehouse", h.StoreWarehouse)

	g.Post("/storeWarehouseMerchandise", h.StoreWarehouseMerchandise)
	g.Post("/updateWarehouseMerchandise", h.UpdateWarehouseMerchandise)
	g.Get("/getAllMerchandiseMoreInfo", h.GetAllMerchandiseMoreInfo)

	g.Post("/storeSoldPoint", h.StoreSoldPoint)
	g.Get("/getAllSoldPoint", h.GetAllSoldPoint)
	g.Get("/getSoldPointById", h.GetSoldPointById)

	g.Post("/storeTransaction", h.StoreTransaction)
	g.Get("/getAllTransactions", h.GetAllTransaction)
	g.Get("/getTransactionBy", h.GetTransactionBy)
}

func (h *handlerApi) GeneralAnalyze(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	notifys, err := h.analyzeService.GeneralWarehouseAnalyze(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//bodyMessage, err := json.Marshal(notifys)
	//if err != nil {
	//	logF.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(notifys)
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
func (h *handlerApi) StoreSoldPoint(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	sP := new(dto.SoldPointCreate)
	if err := ctx.BodyParser(sP); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.soldPointService.StoreSoldPoint(ctx.Context(), sP)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handlerApi) GetAllSoldPoint(ctx *fiber.Ctx) error {

	sP, err := h.soldPointService.GetAllSoldPoints(ctx.Context())
	if err != nil {
		h.log.Errorln(err)
		return err
	}

	//bodyMessage, err := json.Marshal(sP)
	//if err != nil {
	//	h.log.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(sP)
}

func (h *handlerApi) GetSoldPointById(ctx *fiber.Ctx) error {
	idC := ctx.Query("id")
	id, err := strconv.Atoi(idC)
	if err != nil {
		h.log.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	sP, err := h.soldPointService.GetSoldPointById(ctx.Context(), uint(id))
	if err != nil {
		h.log.Errorln(err)
		return err
	}

	//bodyMessage, err := json.Marshal(sP)
	//if err != nil {
	//	h.log.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(sP)
}

func (h *handlerApi) GetAllMerchandiseMoreInfo(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")

	ms, err := h.warehousService.GetAllMerchandiseMoreInfo(ctx.Context(), uint(page))
	if err != nil {
		h.log.Errorln(err)
		return err
	}
	//
	//bodyMessage, err := json.Marshal(ms)
	//if err != nil {
	//	h.log.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(ms)
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

func (h *handlerApi) UpdateWarehouseMerchandise(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	warehouseMerchandiseUpdate := new(dto.WarehouseMerchandise)
	if err := ctx.BodyParser(warehouseMerchandiseUpdate); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.warehousService.UpdateWarehouseMerchandise(ctx.Context(), warehouseMerchandiseUpdate)
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
		//
		//bodyResponse, err := json.Marshal(warehouses)
		//if err != nil {
		//	logF.Errorln(err)
		//	return ctx.SendStatus(fiber.StatusInternalServerError)
		//}

		return ctx.Status(fiber.StatusOK).JSON(warehouses)
	}

	warehouses, err := h.warehousService.GetAll(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	//
	//bodyResponse, err := json.Marshal(warehouses)
	//if err != nil {
	//	logF.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(warehouses)
}

func (h *handlerApi) GetExpireStats(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	queries := ctx.Queries()

	var stats *entity.WarehouseStatistic
	if q, ok := queries["id"]; ok {
		id, err := strconv.Atoi(q)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		stats, err = h.warehousService.GetExpireStatsByWarehouseId(ctx.Context(), uint(id))
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}
	stats, err := h.warehousService.GetAllExpireStats(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(stats)
}

func (h *handlerApi) UpdateWarehouse(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	warehouseUpdate := new(dto.Warehouse)
	if err := ctx.BodyParser(warehouseUpdate); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.warehousService.UpdateWarehouse(ctx.Context(), warehouseUpdate)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
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
	//
	//bodyResponce, err := json.Marshal(warehouse)
	//if err != nil {
	//	logF.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(warehouse)
}

func (h *handlerApi) StoreTransaction(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	transactionCreate := new(dto.TransactionCreate)
	if err := ctx.BodyParser(transactionCreate); err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := h.transactionService.InsertTransaction(ctx.Context(), transactionCreate)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handlerApi) GetTransactionBy(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	var transaction *entity.TransactionInfo
	var err error
	id := ctx.QueryInt("id")
	if id != 0 {
		transaction, err = h.transactionService.GetTransactionById(ctx.Context(), uint(id))
		if err != nil {
			logF.Errorln(err)
			return err
		}
	} else if ctx.QueryInt("warehouseId") != 0 {
		transaction, err = h.transactionService.GetTransactionByWarehousesId(ctx.Context(), uint(ctx.QueryInt("warehouseId")))
		if err != nil {
			logF.Errorln(err)
			return err
		} else if ctx.QueryInt("soldPointId") != 0 {
			transaction, err = h.transactionService.GetTransactionBySoldPointId(ctx.Context(), uint(ctx.QueryInt("soldPointId")))
			if err != nil {
				logF.Errorln(err)
				return err
			} else if ctx.QueryInt("merchandiseId") != 0 {
				transaction, err = h.transactionService.GetTransactionByMerchandiseId(ctx.Context(), uint(ctx.QueryInt("merchandiseId")))
				if err != nil {
					logF.Errorln(err)
					return err
				}
			}
		}
	}
	//
	//bodyResponse, err := json.Marshal(transaction)
	//if err != nil {
	//	logF.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(transaction)
}

func (h *handlerApi) GetAllTransaction(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	transactions, err := h.transactionService.GetAllTransactions(ctx.Context())
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//bodyResponse, err := json.Marshal(transactions)
	//if err != nil {
	//	logF.Errorln(err)
	//	return ctx.SendStatus(fiber.StatusInternalServerError)
	//}

	return ctx.Status(fiber.StatusOK).JSON(transactions)
}

func (h *handlerApi) DeleteTransaction(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	idC := ctx.Query("id")
	id, err := strconv.Atoi(idC)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	err = h.transactionService.DeleteTransaction(ctx.Context(), uint(id))
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
