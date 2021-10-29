package model
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-blog-service/pkg/setting"
	"github.com/gin-blog-service/global"
	"fmt"
)
type Model struct {
	ID	uint32 `gorm:"primary_key" json:"id"`
	CreatedOn string `json:"created_on"`
	CreatedBy string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy uint32 `json:"modified_by"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8	`json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t", databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}