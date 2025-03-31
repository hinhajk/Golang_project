package _case

import (
	"fmt"
	"gorm.io/gorm"
)

func (t *Teacher) BeforeSave(*gorm.DB) error {
	fmt.Println("hook BeforeSave")
	return nil
}

func (t *Teacher) AfterSave(*gorm.DB) error {
	fmt.Println("hook AfterSave")
	return nil
}

func (t *Teacher) BeforeUpdate(*gorm.DB) error {
	fmt.Println("hook BeforeUpdate")
	return nil
}

func (t *Teacher) AfterUpdate(*gorm.DB) error {
	fmt.Println("hook AfterUpdate")
	return nil
}

func (t *Teacher) BeforeDelete(*gorm.DB) error {
	fmt.Println("hook BeforeDelete")
	return nil
}

func (t *Teacher) AfterDelete(*gorm.DB) error {
	fmt.Println("hook AfterDelete")
	return nil
}

func (t *Teacher) AfterFind(*gorm.DB) error {
	fmt.Println("hook AfterFind")
	return nil
}
