package oauth

type Handler struct {
	service *oauthService
}

func NewHandler(oauthSvc *oauthService) *Handler {
	return &Handler{
		service: oauthSvc,
	}
}
