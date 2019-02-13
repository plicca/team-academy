package shift

import (
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type Shift struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	SubjectID   int
	ClassroomID int
	Type        string
	ShiftNum    int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Shift{}) {
		return false, db.CreateTable(Shift{}).Error
	}
	return true, nil
}

func CreateShift(db *gorm.DB, shift Shift) (err error) {
	return db.Save(&shift).Error
}

func UpdateShift(db *gorm.DB, shift Shift) (err error) {
	if shift.ID <= 0 {
		err = component.ErrMissingParameters
		return
	}

	_, err = GetShiftByID(db, shift.ID)
	if err != nil {
		return
	}

	return db.Model(&Shift{}).Update(shift).Error
}

func DeleteShift(db *gorm.DB, id int) (err error) {
	return db.Delete(&Shift{ID: id}).Error
}

func GetShiftByID(db *gorm.DB, id int) (shift Shift, err error) {
	err = db.First(&shift, &Shift{ID: id}).Error
	return
}