package project

type CreateProjectDto struct {
	Name        string `form:"name" validate:"required,min=5,max=50"`
	Alias       string `form:"alias" validate:"required,min=5,max=20"`
	WebhookURL  string `form:"webhook_url" validate:"max=255"`
	Description string `form:"description" validate:"max=140"`
}
