package users

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type Users struct {
	repository NewUsersRepository
}

func NewUsers(repository NewUsersRepository) *Users {
	return &Users{repository}
}

func (u *Users) Registration(c echo.Context) error {
	var payload Payload
	resp := map[string]string{
		"message": "sucess registration",
	}

	err := json.NewDecoder(c.Request().Body).Decode(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	v := u.Validation(payload)

	if len(v) > 0 {
		return c.JSON(http.StatusBadRequest, v)
	}

	regErr := u.repository.Registration(payload)

	if regErr != nil {
		return c.JSON(http.StatusInternalServerError, regErr.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (u *Users) Update(c echo.Context) error {
	var payload Payload
	resp := map[string]string{
		"message": "success update user",
	}

	err := json.NewDecoder(c.Request().Body).Decode(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	errRe := u.repository.Update(payload)

	if errRe != nil {
		return c.JSON(http.StatusInternalServerError, errRe.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (u *Users) Profile(c echo.Context) error {
	id := c.Param("id")

	resp, err := u.repository.Profile(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (u *Users) Login(c echo.Context) error {
	var p Payload
	var r LoginResponse
	if err := json.NewDecoder(c.Request().Body).Decode(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := u.repository.Login(p.PhoneNumber)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if user.Password == p.Password {
		// Generate JWT RS256
	}

	return c.JSON(http.StatusOK, r)
}

func (u *Users) Validation(payload Payload) []map[string]string {
	var err []map[string]string
	re := regexp.MustCompile(`^(62)8[1-9][0-9]{6,9}$`)

	if len(payload.PhoneNumber) < 10 || len(payload.PhoneNumber) > 13 {
		err = append(err, map[string]string{
			"field": "phoneNumber",
			"error": "must be a minimum of 10 characters and a maximum of 13 characters",
		})
	}

	if !re.MatchString(payload.PhoneNumber) {
		err = append(err, map[string]string{
			"field": "phoneNumber",
			"error": "must be +62",
		})
	}

	if len(payload.FullName) < 3 || len(payload.FullName) > 60 {
		err = append(err, map[string]string{
			"field": "fullname",
			"error": "must be a minimum of 3 characters and a maximum of 60 characters",
		})
	}

	if len(payload.Password) < 6 || len(payload.Password) > 64 {
		err = append(err, map[string]string{
			"field": "password",
			"error": "must be a minimum of 6 characters and a maximum of 64 characters",
		})
	}

	return err
}
