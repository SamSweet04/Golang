package main

type EmailService struct {
	db *gorm. DB
	smtpHost string
	smtpPassword string
	smtpPort int
}
func NewEmailService(db *gorm.DB, smtpHost string, smtpPassword string, smtpPort int) *EmailService  { 
	return &EmailService{
		db:  db, 
		smtpHost: smtpHost,
		smtpPassword: smtpPassword,
		smtpPort: smtpPort,
   }

}

func (s *EmailService) Send(from string, to string, subject string, message string) error
 { email := EmailGorm{
	 From: from,
	To: to,
	Subject: subject,
	Message: message,
 }
	err := s.db.Create(&email).Error
	if err != nil {
		log. Println(err)
		return err
	}
	auth := smtp.PlainAuth( identity: "", from, s.smtpPassword, s.smtpHost) server := fmt.Sprintf( format "%s:%d", s.smtpHost, s.smtpPort) err = smtp.SendMail(server, auth, from, []string{to}, []byte(message)) 
	if err != nil {
	
	}
	return nil
}
// Let us examine the code block from above. There we have one struct, EmailService, with only one method, Send. We use this service for sending emails. Even if it looks fine, we realize that this code breaks every aspect of SRP when we scratch the surface.
// The responsibility of EmailService is not just to send emails but to store an email message into DB AND send it via SMTP protocol.
// Take a closer look at the sentence above. The word "and" is bold with purpose. Using such an expression does not look like the case where we describe a single responsibility.

type EmailGorm struct {
	gorm. Model
	From string 
	To string
	Subject string
	Message string
}
type EmailRepository interface {
Save(from string, to string, subject string, message string) error
}
type EmailDBRepository struct {
db *gorm. DB
}
func NewEmailRepository(db #gorm.DB) EmailRepository { return &EmailDBRepository{
db: db,
}
}
func (r *EmailDBRepository) Save(from string, to string, subject string, message string) error {
	email := EmailGorm{
	From: from, To:
	to,
	Subject: subject,
	Message: message,
	}
	err := r.db.Create(&email).Error
	if err != nil {
	log. Println(err)
	return err
	}
	return nil
	}
	type EmailSender interface {
		Send(from string, to string, subject string, message string) error
	}
	type EmailSMTPSender struct { 
		smtpHost string
		smtpPassword string 
		smtpPort int
	}
	func NewEmailSender(smtpHost string, smtpPassword string, smtpPort int) EmailSender { 
		return &EmailSMTPSender{
			smtpHost: smtpHost,
			smtpPassword: smtpPassword,
			smtpPort:
			smtpPort,
		}
	}
	func NewEmailService(repository EmailRepository, sender EmailSender) *EmailService {
		return &EmailService{
			 repository: repository,
		     sender: sender,
		}
	}
func (s *EmailService) Send(from string, to string, subject string, message string) error {
	err := s.repository.Save(from, to, subject, message) 
	if err != nil {
	   return err
}
	return s.sender.Send(from, to, subject, message)
}
// Here we provide two new structs. The first one is EmailDBRepository as an implementation for the EmailRepository interface. It includes support for persisting data in the underlying database.
// The second structure is EmailSMTPSender that implements the EmailSender interface. This struct is responsible for only email sending over SMPT protocol.
// Finally, the new EmailService contains interfaces from above and delegates the request for email sending.


