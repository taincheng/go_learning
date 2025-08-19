package sql_practice2

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
题目2：事务语句
	假设有两个表：
		accounts 表（包含字段 id 主键， balance 账户余额）
		transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，
	实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
	如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
    并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance float64
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64

	// 定义外键关联
	FromAccount Account `gorm:"foreignKey:FromAccountId"`
	ToAccount   Account `gorm:"foreignKey:ToAccountId"`
}

func transactionMoney1(db *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 检查 A 账户余额，并且加更新锁
		var formAccount Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", fromAccountId).
			First(&formAccount).Error; err != nil {
			return fmt.Errorf("查询转出账户余额失败: %w", err)
		}

		// 2. 检查A 账户
		if formAccount.Balance < amount {
			return fmt.Errorf("账户余额不足")
		}

		// 3. 改变 A 账户余额
		if err := tx.Model(&Account{}).
			Where("id = ?", fromAccountId).
			Update("balance", gorm.Expr("balance - ?", amount)).
			Error; err != nil {
			return fmt.Errorf("更新转出用户余额失败：%w", err)
		}

		// 4. 改变 B 账户余额
		if err := tx.Model(&Account{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", toAccountId).
			Update("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return fmt.Errorf("更新转入用户余额失败：%w", err)
		}

		// 5. 写入转账记录
		transaction := Transaction{
			FromAccountId: fromAccountId,
			ToAccountId:   toAccountId,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("创建交易记录失败：%w", err)
		}

		return nil
	})
}

func transactionMoney2(db *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) {
	// 转账事务, A 向 B 转账 100
	tx := db.Begin()
	var A, B Account
	tx.Model(&Account{}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&A, fromAccountId)
	tx.Model(&Account{}).Clauses(clause.Locking{Strength: "UPDATE"}).First(&B, toAccountId)

	if A.Balance >= amount {
		A.Balance -= amount
		B.Balance += amount
		updateA := tx.Model(&A).Update("balance", A.Balance)
		if updateA.Error != nil {
			tx.Rollback()
		}

		updateB := tx.Model(&B).Update("balance", B.Balance)
		if updateB.Error != nil {
			tx.Rollback()
		}

		transaction := Transaction{FromAccountId: A.ID, ToAccountId: B.ID, Amount: amount}
		createTransaction := tx.Create(&transaction)

		if createTransaction.Error != nil {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	tx.Commit()
}

func Run() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db.AutoMigrate(&Account{}, &Transaction{})
	accounts := []Account{
		{Balance: 100},
		{Balance: 200},
	}
	db.Create(&accounts)

	needMoney := 100.0

	//err := transactionMoney1(db, 1, 2, needMoney)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	transactionMoney2(db, 1, 2, needMoney)
}
