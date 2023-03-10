package postgres

import (
	"github.com/mamochiro/beef/internals/utils"

	"gorm.io/gorm"
)

//go:generate mockery --name=Interface
type Interface interface {
	Raw(ent interface{}, sql string, value ...interface{}) (err error)
	Create(ent interface{}) (err error)
	Find(filters map[string]interface{}, ent interface{}) error
	First(filters map[string]interface{}, ent interface{}) error
	FirstOrCreate(filters map[string]interface{}, ent interface{}) error
	Last(filters map[string]interface{}, ent interface{}) error
	List(tableName string, offset, limit int64, filters, order, ent interface{}) (*utils.Pagination, error)
	Update(filters map[string]interface{}, ent interface{}) error
	UpdateByMapInterface(model interface{}, filters map[string]interface{}, fields map[string]interface{}) error
	Begin() *gorm.DB
	Commit(db *gorm.DB) *gorm.DB
	Rollback(db *gorm.DB) *gorm.DB
	SavePoint(db *gorm.DB, input string) *gorm.DB
	RollbackTo(db *gorm.DB, input string) *gorm.DB
	CreateTransaction(db *gorm.DB, ent interface{}) error
	UpdateTransaction(db *gorm.DB, filters map[string]interface{}, ent interface{}) error
	IsErrorRecordNotFound(err error) bool
	Upsert(tableName string, filters map[string]interface{}, ent interface{}) error
}
