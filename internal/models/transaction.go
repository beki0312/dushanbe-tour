package models

import "time"

type GetCourse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      string    `json:"amount"`
	ImageName   string    `json:"image_name" gorm:"column:image_name"`
	Image       string    `json:"image" gorm:"column:image"`
	DateBegin   time.Time `json:"date_begin"`
	DateEnd     time.Time `json:"date_end"`
	Created_At  time.Time `json:"created_At"`
}

type GetLesson struct {
	Id          int64  `json:"id"`
	CourseID    int64  `json:"course_id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Time        string `json:"time"`
	CreatedAt   string `json:"created_at"`
}

type Language struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Infrastructure struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type InfrastructureAll struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Link       string `json:"link"`
	Mail       string `json:"email"`
	Descriptin string `json:"descriptin"`
	Image      string `json:"image"`
	Active     bool   `json:"active"`
	Created_at string `json:"created_at"`
}

type GetCourseNew struct {
	Id                    int64     `json:"id"`
	CourseName            string    `json:"courseName" gorm:"column:name"`
	Title                 string    `json:"title"`
	CourseDescription     string    `json:"courseDescription" gorm:"column:description"`
	Amount                string    `json:"amount"`
	ImageName             string    `json:"image_name" gorm:"column:image_name"`
	Image                 string    `json:"image" gorm:"column:image"`
	Category              string    `json:"category"`
	Status                string    `json:"status"`
	InstructorName        string    `json:"instructorName" gorm:"column:name"`
	InsrtuctorUnvon       string    `json:"insrtuctorUnvon" gorm:"column:unvon"`
	InstructorDescription string    `json:"instructorDescription" gorm:"column:description"`
	Lectures              int       `json:"lectures"`
	Quizzes               int       `json:"quizzes"`
	Duration              string    `json:"duration"`
	SkillLevel            string    `json:"skill_Level"`
	Language              string    `json:"language"`
	Students              int       `json:"students"`
	Assessments           string    `json:"assessments"`
	DateBegin             time.Time `json:"date_begin"`
	DateEnd               time.Time `json:"date_end"`
	CreatedAt             time.Time `json:"created_At"`
}

type Student struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Active    bool   `json:"active"`
	Created   string `json:"created"`
}

type Registration struct {
	FirstName string `json:"firs_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type Auth struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Log struct {
	Id   int64  `json:"id"`
	Hash string `json:"password" gorm:"column:password"`
}
type SignUp struct {
	CourseId  int64 `json:"course_id"`
	StudentId int64 `json:"student_id"`
}
