package transcation

type Service interface {
	GetAllTranscation() ([]Transcation, error)

	GetTranscationById(ID int) (*Transcation, error)

	CreateTranscation(userID int, transcation *TranscationSpec) error

	UpdateTranscation(ID int, transcation *TranscationSpec) error
}

type Repository interface {
	GetAllTranscation() ([]Transcation, error)

	GetTranscationById(ID int) (*Transcation, error)

	CreateTranscation(transcation Transcation) error

	UpdateTranscation(ID int, transcation Transcation) error
}
