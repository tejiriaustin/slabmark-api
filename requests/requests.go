package requests

type (
	CreateUserRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	AddAccountRequest struct {
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		Email      string `json:"email"`
		Department string `json:"department"`
		Phone      string `json:"phone"`
	}

	EditAccountRequest struct {
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		Email      string `json:"email"`
		Department string `json:"department"`
		Phone      string `json:"phone"`
	}

	LoginUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ForgotPasswordRequest struct {
		EmailAddress string
	}

	ResetPasswordRequest struct {
		ResetCode   string
		NewPassword string
	}
)
