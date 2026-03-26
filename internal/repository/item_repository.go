package repository

import (
	"github.com/woulongplum/Box-watcher/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db:db}
} 

func (r *ItemRepository) Upsert(item *model.Item) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "url"}},
		
		DoUpdates: clause.AssignmentColumns([]string{"name", "price", "in_stock", "status", "source", "updated_at"}),
	}).Create(item).Error
}
 