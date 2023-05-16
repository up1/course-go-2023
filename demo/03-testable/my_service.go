package main

// type MyService struct {
// }

// func (s *MyService) Process() error {
// 	// 1. Create database connection
// 	// 2. Insert data into database
// 	// 3. Create email service
// 	// 4. Send email
// }

type MyService struct {
	Db    IDatabase
	Email IEmailService
}

func (s *MyService) Process() error {
	// 1. Create database connection
	s.Db.connect()
	// 2. Insert data into database
	s.Db.insertData()
	// 3. Create email service
	s.Email.connect()
	// 4. Send email
	s.Email.send()

	return nil
}

func NewMyService(db IDatabase, email IEmailService) *MyService {
	return &MyService{
		Db:    db,
		Email: email,
	}
}

type IDatabase interface {
	connect() error
	insertData() error
}

type Database struct {
}

func (d *Database) connect() error {
	return nil
}
func (d *Database) insertData() error {
	return nil
}

type IEmailService interface {
	connect() error
	send() error
}

type EmailService struct {
}

func (e *EmailService) connect() error {
	return nil
}

func (e *EmailService) send() error {
	return nil
}
