package helper

import (
	"fmt"
	"gorm.io/gorm"
)

func RollbackHelper(tx *gorm.DB) {
	if r := recover(); r != nil {
		if tx != nil && tx.Statement != nil && tx.Statement.ConnPool != nil {
			tx.Rollback()
			fmt.Println("Rollback transaction from panic")
		}
		panic(r)
	} else if tx != nil && tx.Statement != nil && tx.Statement.ConnPool != nil {
		tx.Rollback()
		fmt.Println("Rollback uncommitted transaction")
	}
}
