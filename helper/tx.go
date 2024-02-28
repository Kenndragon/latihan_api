package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	if err := recover(); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			PanicError(errorRollback)
		}
		panic(err)
	} else {
		if errorCommit := tx.Commit().Error; errorCommit != nil {
			PanicError(errorCommit)
		}
	}

}
