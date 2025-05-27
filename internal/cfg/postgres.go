package cfg

type PostgresConfig struct {
	ConnectionString     string // Адрес БД со справочными данными и данными для аутентификации
	DataConnectionString string // Адрес pg-sharding/SPQR или connection pooler на данные бизнесовые данные
}
