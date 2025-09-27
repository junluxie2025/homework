package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type account struct {
	ID      string
	Balance float64
}

type transaction struct {
	ID            int64
	FromAccountId string
	ToAacountId   string
	Amount        float64
}

func RunTransaction(db *gorm.DB) {

	db.AutoMigrate(&account{}, &transaction{})
	accounts := []account{
		{ID: "A", Balance: 101},
		{ID: "B", Balance: 0},
	}
	db.Create(&accounts)

	err := db.Transaction(func(tx *gorm.DB) error {
		//A->B 转账 100
		accountA := account{ID: "A"}
		tx.Find(&accountA)

		if accountA.Balance < 100 {
			return errors.New("转账失败")
		} else {
			tx.Model(&account{}).Where("id = ?", "A").UpdateColumn("balance", gorm.Expr("balance - ?", 100))
			tx.Model(&account{}).Where("id = ?", "B").UpdateColumn("balance", gorm.Expr("balance + ?", 100))
			err := db.Create(&transaction{
				FromAccountId: "A",
				ToAacountId:   "B",
				Amount:        100,
			}).Error
			if err != nil {
				return err
			}
			return nil
		}

	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("转账成功")
	}
}
