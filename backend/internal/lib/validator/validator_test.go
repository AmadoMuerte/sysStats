package validator

import "testing"

func TestValidatePassword(t *testing.T) {
	validPasswords := []string{
		"A12345678",
		"L!@#$%^&*",
		"Password123",
		"L1234567890",
		"V!@#$%^&*()",
		"Password123!@#$%^&*",
	}

	for _, pass := range validPasswords {
		if err := ValidatePassword(pass); err != nil {
			t.Fatalf("Expected password %s to be valid", pass)
		}
	}

	invalidPasswords := []string{
		"1234567",
		"Password",
		"Password123!@#$%^&*()_",
	}

	for _, pass := range invalidPasswords {
		if err := ValidatePassword(pass); err == nil {
			t.Fatalf("Expected password %s to be invalid", pass)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	validEmails := []string{
		"test@mail.ru",
		"test@mail.com",
		"test@google.com",
		"test@yahoo.com",
		"Test@mail.ru",
		"BigTestmailcheck@ya.ru",
	}

	for _, email := range validEmails {
		if !ValidateEmail(email) {
			t.Fatalf("Expected email %s to be valid", email)
		}
	}

	invalidEmails := []string{
		"testmail.ru",
		"test.com",
		"testgoogle.com",
		"testyahoo.com",
		"Testmail.ru",
		"BigTestmailcheck@ya",
		"test@.ru",
		"test@.com",
		"test@mail.com.",
	}

	for _, email := range invalidEmails {
		if ValidateEmail(email) {
			t.Fatalf("Expected email %s to be invalid", email)
		}
	}
}
