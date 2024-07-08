package device

type CreateDeviceDto struct {
	Name             string `form:"name" validate:"required,min=5,max=20"`
	Description      string `form:"description" validate:"max=140"`
	ClientDeviceID   string `form:"client_device_id" validate:"required"`
	CustomAttributes string `form:"custom_attributes" validate:"json"`
}

type RequestQRDto struct {
	SessionID string `query:"session_id" validate:"required"`
}
