package user

type UserFornatter struct {
	IDUser string `json:"id_user"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFornatter {
	formatter := UserFornatter {
		IDUser: user.IDUser.String(),
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		Role: user.Role,
		Token: token,
	}
	return formatter
}