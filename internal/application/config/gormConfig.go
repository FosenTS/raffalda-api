package config

import (
	"raffalda-api/pkg/mysync"
)

const gormConfigFilename = "gorm.config.yaml"

type GormConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	UseCA        bool
	CaPath       string
	TimeZone     string

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool `yaml:"skipDefaultTransaction" env-required:"true"`

	// FullSaveAssociations full save associations
	FullSaveAssociations bool `yaml:"fullSaveAssociations" env-required:"true"`

	// DryRun generate sql without execute
	DryRun bool `yaml:"dryRun" env-required:"true"`
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool `yaml:"prepareStmt" env-required:"true"`
	// DisableAutomaticPing
	DisableAutomaticPing bool `yaml:"disableAutomaticPing" env-required:"true"`
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool `yaml:"disableForeignKeyConstraintWhenMigrating" env-required:"true"`
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool `yaml:"ignoreRelationshipsWhenMigrating" env-required:"true"`
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool `yaml:"disableNestedTransaction" env-required:"true"`
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool `yaml:"allowGlobalUpdate" env-required:"true"`
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool `yaml:"queryFields" env-required:"true"`
	// CreateBatchSize default create batch size
	CreateBatchSize int `yaml:"createBatchSize" env-required:"true"`
	// TranslateError enabling error translation
	TranslateError bool `yaml:"translateError" env-required:"true"`
}

var (
	gormConfigInst     = &GormConfig{}
	loadGormConfigOnce = mysync.NewOnce()
)

func Gorm() GormConfig {
	loadGormConfigOnce.Do(func() {
		env := Env()
		gormConfigInst.Host = env.PostgresHost
		gormConfigInst.Port = env.PostgresPort
		gormConfigInst.Username = env.PostgresUsername
		gormConfigInst.Password = env.PostgresPassword
		gormConfigInst.DatabaseName = env.PostgresDatabaseName
		gormConfigInst.UseCA = env.PostgresUseCA
		gormConfigInst.CaPath = env.PostgresCaPath
		gormConfigInst.TimeZone = env.PostgresTimeZone

	})

	return *gormConfigInst
}
