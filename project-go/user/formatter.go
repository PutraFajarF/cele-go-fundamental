package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NoHandphone string `json:"no_handphone"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		NoHandphone: user.NoHandphone,
		Email:       user.Email,
		Token:       token,
	}

	return formatter
}
