package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"main.go/auth"
	"main.go/model"
)

type WebId struct {
	Id uint `json:"id"`
}
type WebsiteResponse struct {
	ID          uint             `json:"id"`
	Created_at  time.Time        `json:"created_at"`
	Updated_at  time.Time        `json:"updated_at"`
	User_id     uint             `json:"user_id"`
	Site_key    string           `json:"site_key"`
	Secret_key  string           `json:"secret_key"`
	Label       string           `json:"label"`
	Alert       bool             `json:"alert"`
	Subdomain   bool             `json:"subdomain"`
	Version     uint             `json:"version"`
	Alert_limit uint             `json:"alert_limit"`
	Website_v1  model.Website_v1 `json:"web_v1"`
	Token       string           `json:"token"`
}

func NewWebResponse(web *model.Website) *WebsiteResponse {
	token, _ := auth.GenerateJWT(web.ID, web.Label)
	wr := &WebsiteResponse{
		ID:          web.ID,
		Created_at:  web.Created_at,
		Updated_at:  web.Updated_at,
		User_id:     web.User_id,
		Site_key:    web.Site_key,
		Secret_key:  web.Secret_key,
		Label:       web.Label,
		Alert:       web.Alert,
		Subdomain:   web.Subdomain,
		Version:     web.Version,
		Alert_limit: web.Alert_limit,
		Website_v1:  web.Website_v1,
		Token:       token,
	}
	return wr
}

func (u *Users) getWebsite(c echo.Context) error {
	var input WebId
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	web, err := u.Store.GetWebsite(input.Id)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, NewWebResponse(web))

}
func (u *Users) updateWebsite(c echo.Context) error {
	var input WebsiteResponse
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	newWeb := &model.Website{
		ID:          input.ID,
		Site_key:    input.Site_key,
		Secret_key:  input.Secret_key,
		Label:       input.Label,
		Alert:       input.Alert,
		Subdomain:   input.Subdomain,
		Version:     input.Version,
		Alert_limit: input.Alert_limit,
		Updated_at:  time.Now(),
	}
	err := u.Store.UpdateWebsite(newWeb)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not update website in the database", err)
	}
	return c.JSON(http.StatusOK, NewWebResponse(newWeb))
}
