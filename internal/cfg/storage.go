package cfg

import (
	"money/internal/core"
	"money/internal/repo/postgres"
)

func GetStorages(p PostgresConfig, s SecurityConfig) (*core.Storage, error) {
	ps, err := postgres.GetPostgresStorage(p.ConnectionString, s.JWTKey)
	if err != nil {
		return nil, err
	}

	cs := &core.Storage{}
	cs.UserStorager = ps
	return cs, err
}
