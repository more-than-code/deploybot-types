package types

type AuthenticationInput struct {
	Email    string
	Password string
}

type AuthenticationSsoInput struct {
	IdToken string
}

type AuthenticationOutput struct {
	UserId       ObjectId `json:"userId"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
}

type User struct {
	Id           ObjectId `json:"id" bson:"_id"`
	Subject      string   `json:"subject"`
	Email        string   `json:"email"`
	ContactEmail string   `json:"contactEmail"`
	Password     string   `json:"password"`
	Name         string   `json:"name"`
	AvatarUrl    string   `json:"avatarUrl"`
	CreatedAt    Datetime `json:"createdAt"`
}

type CreateUserInput struct {
	Name             string
	Email            string
	Password         string
	VerificationCode string
}

type GetUsersInput struct {
	UserIds []ObjectId
}

type GetUsersOutput struct {
	Items      []User `json:"items"`
	TotalCount int    `json:"totalCount"`
}

type Claims struct {
	Iss           string
	Nbf           int64
	Aud           string
	Sub           string
	Email         string
	EmailVerified bool `json:"email_verified"`
	Azp           string
	Name          string
	Picture       string
	GivenName     string `json:"given_name"`
	Iat           int64
	Exp           int64
	Jti           string
}
