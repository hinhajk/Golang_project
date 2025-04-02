package _case

import (
	"errors"
	"gorm.io/gorm"
)

// 事务
func Transaction() {
	t := teacherTemp
	t1 := teacherTemp
	t2 := teacherTemp
	// 创建事务
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		if err := tx.Create(&t1).Error; err != nil {
			return err
		}

		//回滚子事务，不影响大事务最终结果
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(t2)
			return errors.New("rollback t1")
		})

		return nil
	})
}
