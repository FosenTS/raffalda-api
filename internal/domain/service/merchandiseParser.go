package service

import (
	"context"
	"os"
	"raffalda-api/internal/application/config"
	"raffalda-api/internal/domain/service/sto"

	"github.com/gocarina/gocsv"
)

type MerchandiseParser interface {
	MerchandiseFileParse(ctx context.Context) ([]*sto.MerchandiseParse, error)
}

type merchandiseParser struct {
	config config.MerchandiseParserConfig
}

func NewMerchandiseParser(config config.MerchandiseParserConfig) MerchandiseParser {
	return &merchandiseParser{
		config: config,
	}
}

func (mP *merchandiseParser) MerchandiseFileParse(ctx context.Context) ([]*sto.MerchandiseParse, error) {
	merchandiseFile, err := os.Open(mP.config.DataFileAbsPath)
	if err != nil {
		return nil, err
	}

	merchindiseData := make([]*sto.MerchandiseParse, 0)

	if err := gocsv.UnmarshalFile(merchandiseFile, &merchindiseData); err != nil {
		return nil, err
	}

	return merchindiseData, nil
}
