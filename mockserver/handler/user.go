package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/util/v8"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver/logic"
)

// HandleGetUsers is the handler of GET Users API.
func HandleGetUsers(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /users List all users
	users, sc, err := lgc.GetUsers()
	for i, u := range users {
		u.Password = ""
		users[i] = u
	}
	if err != nil {
		return nil, sc, err
	}
	// TODO authorization
	return &graylog.UsersBody{Users: users}, sc, nil
}

// HandleGetUser is the handler of GET User API.
func HandleGetUser(
	u *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /users/{username} Get user details
	name := ps.PathParam("username")
	// TODO authorization
	user, sc, err := lgc.GetUser(name)
	if user != nil {
		user.Password = ""
	}
	return user, sc, err
}

// HandleCreateUser is the handler of Create User API.
func HandleCreateUser(
	u *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// POST /users Create a new user account.
	if sc, err := lgc.Authorize(u, "users:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet("username", "email", "permissions", "full_name", "password"),
			Optional: set.NewStrSet("startpage", "timezone", "session_timeout_ms", "roles"),
		})
	if err != nil {
		return nil, sc, err
	}

	user := &graylog.User{}
	if err := util.MSDecode(body, user); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return nil, 400, err
	}

	sc, err = lgc.AddUser(user)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleUpdateUser is the handler of Update User API.
func HandleUpdateUser(
	u *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /users/{username} Modify user details.
	userName := ps.PathParam("username")
	if sc, err := lgc.Authorize(u, "users:edit", userName); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Optional:     set.NewStrSet("email", "permissions", "full_name", "password", "timezone", "session_timeout_ms", "start_page", "roles"),
			Ignored:      set.NewStrSet("id", "preferences", "external", "read_only", "session_active", "last_activity", "client_address"),
			ExtForbidden: false,
		})
	if err != nil {
		return nil, sc, err
	}

	prms := &graylog.UserUpdateParams{Username: userName}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as UserUpdateParams")
		return nil, 400, err
	}
	sc, err = lgc.UpdateUser(prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleDeleteUser is the handler of Delete User API.
func HandleDeleteUser(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /users/{username} Removes a user account
	name := ps.PathParam("username")
	// TODO authorization
	sc, err := lgc.DeleteUser(name)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
