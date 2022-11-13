package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Paginate(app *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(app.DefaultQuery("page", "0"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(app.DefaultQuery("page_size", "100"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
