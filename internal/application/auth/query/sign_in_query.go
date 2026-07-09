package query

type SignInQuery struct {
	Login    string
	Password string
}

type SignInQueryResult struct {
	AccessToken string
}
