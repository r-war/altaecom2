package constants

// Date format constants
const (
	Authorization     = "Authorization"
	DateTimeFormatStd = "2006-01-02 15:04:05"
	DateFormatStd     = "2006-01-02"
	TimeFormatStd     = "15:04:05"
	TimeFormatShort   = "15:04"
)

// Gender enum
const (
	GenderMale   = "pria"
	GenderFemale = "wanita"
)

// Mail templates
const (
	MailMainTemplate          = "templates/layout/Main.html"
	MailLogoTemplate          = "templates/layout/Logo.html"
	MailFooterTemplate        = "templates/layout/Footer.html"
	MailResetPasswordTemplate = "templates/page/ResetPassword.html"
	DESILogoURL               = "https://desi-test.s3.ap-southeast-1.amazonaws.com/logo/608f1a20a23878e1368c563e.png"
)

// OTP purposes
const (
	OTPPurposeResetPasswordUser  = "reset-password-user"
	OTPPurposeResetPasswordAdmin = "reset-password-admin"
)

const (
	DefaultEntity           = "DESI"
	DefaultEmploymentStatus = "DEVELOPER DESI"
	DefaultGrade            = "L14"
	DefaultAdminID          = 1
)

// ValidGender return true if input is a valid gender enum
func ValidGender(gender string) bool {
	return gender == GenderMale || gender == GenderFemale
}

// NormalizeGender return gender value if valid, otherwise return empty string
func NormalizeGender(gender string) string {
	if gender == "" {
		return ""
	}

	if gender == "male" {
		gender = GenderMale
	} else if gender == "female" {
		gender = GenderFemale
	}

	if ValidGender(gender) {
		return gender
	}

	return ""
}
