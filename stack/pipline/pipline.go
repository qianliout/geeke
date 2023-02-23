package pipline

import (
	"time"

	"outback/stack/items"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Create interface {
	CreateProfile(ctx context.Context, data items.Profile) error
	CreateCode(ctx context.Context, data items.NameCode) error
	CreateBalance(ctx context.Context, data items.Balance) error
	CreateCodeCashFlow(ctx context.Context, data items.CashFlow) error
	SearchNameCode(ctx context.Context) ([]items.NameCode, error)
	UpdateNameCode(ctx context.Context, code string, update map[string]interface{}) error
}

type Dao struct {
	db *gorm.DB
}

func (d *Dao) UpdateNameCode(ctx context.Context, code string, update map[string]interface{}) error {
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	err := d.db.WithContext(tctx).Model(new(items.NameCode)).Where("code = ?", code).Updates(update).Error
	return err
}

func (d *Dao) SearchNameCode(ctx context.Context) ([]items.NameCode, error) {
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	res := make([]items.NameCode, 0)
	err := d.db.WithContext(tctx).Find(&res).Error
	return res, err
}

func (d *Dao) CreateBalance(ctx context.Context, data items.Balance) error {
	if err := d.db.AutoMigrate(&data); err != nil {
		log.Error().Err(err)
		return err
	}
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	return d.db.WithContext(tctx).Create(&data).Error
}

func (d *Dao) CreateCodeCashFlow(ctx context.Context, data items.CashFlow) error {
	// if err := d.db.AutoMigrate(&data); err != nil {
	// 	log.Error().Err(err)
	// 	return err
	// }
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	return d.db.WithContext(tctx).Create(&data).Error
}

func (d *Dao) CreateProfile(ctx context.Context, data items.Profile) error {
	// if err := d.db.AutoMigrate(&data); err != nil {
	// 	log.Error().Err(err)
	// 	return err
	// }
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	return d.db.WithContext(tctx).Create(&data).Error
}

func (d *Dao) CreateCode(ctx context.Context, data items.NameCode) error {
	// if err := d.db.AutoMigrate(&data); err != nil {
	// 	log.Error().Err(err)
	// 	return err
	// }
	tctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	return d.db.WithContext(tctx).Create(&data).Error
}

func NewCreate(db *gorm.DB) *Dao {
	return &Dao{db: db}
}
