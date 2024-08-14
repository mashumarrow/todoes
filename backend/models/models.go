package models

import (
    //"gorm.io/gorm"
)

type Subject struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null"`
}

type Classroom struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null"`
}

type Schedule struct {
    ID         uint   `gorm:"primaryKey"`
    SubjectID  uint   `gorm:"not null"`
    ClassroomID uint  `gorm:"not null"`
    DayOfWeek  string `gorm:"type:enum('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday');not null"`
    Period     int    `gorm:"not null"`
    Subject    Subject
    Classroom  Classroom
}

type Todo struct {
    ID        uint   `gorm:"primaryKey"`
    SubjectID uint   `gorm:"not null"`
    Title     string `gorm:"not null"`
    Completed bool   `gorm:"default:false"`
    Subject   Subject
}
