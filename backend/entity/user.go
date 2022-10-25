package entity

import (
	"gorm.io/gorm"
)


type Student struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex"`
	S_ID          string `gorm:"uniqueIndex"`
	Phone         string


	Registration []Registration `gorm:"ForeignKey: StudentID"`
}


type Subject struct {
	gorm.Model
	CODE string
	Name         string
	// Credit       uint
	// Section      uint
	// Day          string
	// Take         uint
	Registration []Registration `gorm:"ForeignKey: SubjectID"`

	// TeacherID *uint
	// Teacher   Teacher

	// FacultyID *uint
	// Faculty   Faculty

	// OfficerID *uint
	// Officer   Officer

	// TimeID *uint
	// Time   Time
}

type State struct {
	gorm.Model
	Name         string  `gorm:"uniqueIndex"`
	Registration []Registration `gorm:"ForeignKey: StateID"`
}

type Registration struct {
	gorm.Model

	SubjectID *uint
	Subject   Subject `gorm:"references:id"`
	StudentID *uint
	Student   Student `gorm:"references:id"`
	StateID   *uint
	State     State `gorm:"references:id"`
}
