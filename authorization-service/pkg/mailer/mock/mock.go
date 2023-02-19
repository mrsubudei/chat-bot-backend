package mock_mailer

type MockMailer struct {
}

func NewMockMailer() *MockMailer {
	return &MockMailer{}
}

func (v MockMailer) DialAndSend(userEmail string, userId int32,
	verificationToken string) error {

	return nil
}
