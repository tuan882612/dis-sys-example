package auth

const (
	QueryGetUserCreds = `
		SELECT user_id, password, status, role
		FROM public.users
		WHERE email = $1`

	QueryRegisterUser = `
		INSERT INTO public.users (user_id, email, password, created, status, role)
		VALUES ($1, $2, $3, $4, $5, $6)`
)
