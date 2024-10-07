package models

type AuthInfo struct {
	UserID string
	Email  string
	Role   string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string  `json:"token"`
	User  GetUser `json:"user"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type RegisterConfirm struct {
	Email string     `json:"email"`
	Otp   string     `json:"otp"`
	User  CreateUser `json:"user"`
}

type ChangePassword struct {
	Login       string `json:"login"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
