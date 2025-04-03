package cash

type CashStorage struct {
	RoleCash
	LevelCash
}

func NewCashStorage() *CashStorage {
	return &CashStorage{RoleCash: RoleCash{}, LevelCash: LevelCash{}}
}
