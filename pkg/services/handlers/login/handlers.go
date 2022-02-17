package userhandler

import (
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"kubegems.io/pkg/model/client"
	"kubegems.io/pkg/model/forms"
	"kubegems.io/pkg/services/auth"
	"kubegems.io/pkg/services/handlers"
	"kubegems.io/pkg/services/utils"
)

var (
	tags = []string{"login"}
)

type Handler struct {
	Path        string
	ModelClient client.ModelClientIface
}

func (h *Handler) Login(req *restful.Request, resp *restful.Response) {
	cred := &auth.Credential{}
	if err := utils.BindData(req, cred); err != nil {
		utils.BadRequest(resp, err)
		return
	}
	authModule := auth.NewAuthenticateModule(h.ModelClient)
	authenticator := authModule.GetAuthenticateModule(cred.Source)
	if authenticator == nil {
		utils.Unauthorized(resp, nil)
		return
	}
	uinfo, err := authenticator.GetUserInfo(cred)
	if err != nil {
		utils.Unauthorized(resp, err)
		return
	}
	uinternel := h.getOrCreateUser(uinfo)
	now := time.Now()
	uinternel.LastLoginAt = &now
	h.ModelClient.Update(uinternel.AsObject())
	user := &forms.UserCommon{
		Username: uinternel.Username,
		Email:    uinternel.Email,
		ID:       uinternel.ID,
		Role:     uinternel.Role,
	}
	jwt := &auth.JWT{}
	token, _, err := jwt.GenerateToken(user, time.Duration(time.Hour*24))
	if err != nil {
		utils.Unauthorized(resp, err)
	}
	utils.OK(resp, token)
}

func (h *Handler) getOrCreateUser(uinfo *auth.UserInfo) *forms.UserInternal {
	u := forms.UserInternal{}
	uobj := u.AsObject()
	if h.ModelClient.Exist(uobj, client.Where("username", client.Eq, uinfo.Username)) {
		h.ModelClient.Get(uobj, client.Where("username", client.Eq, uinfo.Username))
		return u.Data()
	}
	newUser := &forms.UserInternal{
		Username: uinfo.Username,
		Email:    uinfo.Email,
		Source:   uinfo.Source,
	}
	h.ModelClient.Create(newUser.AsObject())
	return newUser.Data()
}

func (h *Handler) Regist(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/" + h.Path)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/").
		To(h.Login).
		Doc("login, get token").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, handlers.MessageOK, nil))

	container.Add(ws)
}

type User struct {
	Username string
}