package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/util/v8"
	"github.com/suzuki-shunsuke/go-set"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver/logic"
)

// HandleGetRole is the handler of GET Role API.
func HandleGetRole(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /roles/{rolename} Retrieve permissions for a single role
	// get parameter
	// logging
	name := ps.PathParam("rolename")
	lgc.Logger().WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	// authorization
	if sc, err := lgc.Authorize(user, "roles:read", name); err != nil {
		return nil, sc, err
	}
	// call logic
	return lgc.GetRole(name)
}

// HandleGetRoles is the handler of GET Roles API.
func HandleGetRoles(
	user *graylog.User, lgc *logic.Logic, _ *http.Request, _ Params,
) (interface{}, int, error) {
	// GET /roles List all roles
	arr, total, sc, err := lgc.GetRoles()
	if err != nil {
		return arr, sc, err
	}
	return &graylog.RolesBody{Roles: arr, Total: total}, sc, nil
}

// HandleCreateRole is the handler of Create Role API.
func HandleCreateRole(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// POST /roles Create a new role
	if sc, err := lgc.Authorize(user, "roles:create"); err != nil {
		return nil, sc, err
	}
	// validate and filter request body
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("name", "permissions"),
			Optional:     set.NewStrSet("description"),
			Ignored:      set.NewStrSet("read_only"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	role := &graylog.Role{}
	if err := util.MSDecode(body, &role); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Warn("Failed to parse request body as Role")
		return nil, 400, err
	}

	if sc, err := lgc.AddRole(role); err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return role, sc, nil
}

// HandleUpdateRole is the handler of Update Role API.
func HandleUpdateRole(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /roles/{rolename} Update an existing role
	name := ps.PathParam("rolename")
	if sc, err := lgc.Authorize(user, "roles:edit", name); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("name", "permissions"),
			Optional:     set.NewStrSet("description"),
			Ignored:      set.NewStrSet("read_only"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	prms := &graylog.RoleUpdateParams{}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Role")
		return nil, 400, err
	}

	role, sc, err := lgc.UpdateRole(name, prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return role, sc, nil
}

// HandleDeleteRole is the handler of Delete Role API.
func HandleDeleteRole(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
	name := ps.PathParam("rolename")
	if sc, err := lgc.Authorize(user, "roles:delete", name); err != nil {
		return nil, sc, err
	}
	sc, err := lgc.DeleteRole(name)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
