package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thomaszub/go-url-shortener/api"
)

type Controller struct {
	urls map[string]*api.Url
}

func NewController() *Controller {
	return &Controller{
		urls: make(map[string]*api.Url),
	}
}

func (s *Controller) AddUrl(ctx echo.Context) error {
	url := new(api.Url)
	if err := ctx.Bind(url); err != nil {
		return err
	}
	h := md5.Sum([]byte(url.Url))
	id := hex.EncodeToString(h[:])
	url.Id = id
	s.urls[id] = url
	return ctx.JSON(http.StatusCreated, url)
}

func (s *Controller) DeleteUrl(ctx echo.Context, id string) error {
	delete(s.urls, id)
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Controller) FindURLById(ctx echo.Context, id string) error {
	url, found := s.urls[id]
	if !found {
		return ctx.JSON(http.StatusNotFound, api.Error{
			Code:    404,
			Message: fmt.Sprintf("URL with id %s not found", id),
		})
	}
	return ctx.JSON(http.StatusOK, url)
}
