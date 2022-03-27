package params

type Email struct {
	SMTPHost           string
	Port               int
	Username           string
	Password           string
	Sender             string
	Receivers          []string
	Subject            string
	Body               string
	AttachmentFilePath []string
}
