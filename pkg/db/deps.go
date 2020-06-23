package db

import (
	"github.com/google/wire"

	"github.com/skygeario/skygear-server/pkg/auth/config"
)

func ProvideSQLBuilder(c *config.DatabaseCredentials, id config.AppID) SQLBuilder {
	return NewSQLBuilder("auth", c.DatabaseSchema, string(id))
}

var DependencySet = wire.NewSet(
	NewContext,
	wire.Struct(new(SQLExecutor), "*"),
	ProvideSQLBuilder,
)