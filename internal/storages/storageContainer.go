package storages

import (
	"github.com/jinzhu/gorm"
)

type StorageContainer struct {
	BankStorage          bankStorage
	CurrencyStorage      currencyStorage
	ExpenseStorage       expenseStorage
	IncomeStorage        incomeStorage
	PaymentPlanStorage   paymentPlanStorage
	PaymentStorage       paymentStorage
	UserStorage          userStorage
	AdvertiserStorage    AdvertiserStorage
	AdvertisementStorage AdvertisementStorage
	BannerStorage        BannerStorage
	BannerPlaceStorage   BannerPlaceStorage
}

func NewStorageContainer(db gorm.DB) StorageContainer {
	return StorageContainer{
		BankStorage:          bankStorage{DB: db},
		CurrencyStorage:      currencyStorage{DB: db},
		ExpenseStorage:       expenseStorage{DB: db},
		IncomeStorage:        incomeStorage{DB: db},
		PaymentPlanStorage:   paymentPlanStorage{DB: db},
		PaymentStorage:       paymentStorage{DB: db},
		UserStorage:          userStorage{DB: db},
		AdvertiserStorage:    AdvertiserStorage{db: db},
		AdvertisementStorage: AdvertisementStorage{db: db},
		BannerStorage:        BannerStorage{db: db},
		BannerPlaceStorage:   BannerPlaceStorage{db: db},
	}
}
