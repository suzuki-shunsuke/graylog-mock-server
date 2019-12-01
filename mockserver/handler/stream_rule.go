package handler

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/v8/util"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver/logic"
)

// HandleGetStreamRules is the handler of Get Stream Rules API.
func HandleGetStreamRules(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	streamID := ps.PathParam("streamID")
	arr, total, sc, err := lgc.GetStreamRules(streamID)
	if err != nil {
		return nil, sc, err
	}
	return &graylog.StreamRulesBody{StreamRules: arr, Total: total}, sc, nil
}

// HandleGetStreamRule is the handler of Get a Stream Rule API.
func HandleGetStreamRule(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules/{streamRuleId} Get a single stream rules
	// TODO authorization
	return lgc.GetStreamRule(
		ps.PathParam("streamID"), ps.PathParam("streamRuleID"))
}

// HandleCreateStreamRule is the handler of Create a Stream Rule API.
func HandleCreateStreamRule(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	streamID := ps.PathParam("streamID")
	ok, err := lgc.HasStream(streamID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("stream <%s> not found", streamID)
	}

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("value", "field"),
			Optional:     set.NewStrSet("type", "description", "inverted"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	rule := &graylog.StreamRule{}
	if err := util.MSDecode(body, rule); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return nil, 400, err
	}
	lgc.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	sc, err = lgc.AddStreamRule(rule)
	if err != nil {
		logic.LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "rule": rule, "status_code": sc,
		}), "Faield to add rule to mock server")
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": rule.ID}, sc, nil
}

// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}

// HandleUpdateStreamRule is the handler of Update a Stream Rule API.
func HandleUpdateStreamRule(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	streamID := ps.PathParam("streamID")
	ruleID := ps.PathParam("streamRuleID")

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("value", "field"),
			Optional:     set.NewStrSet("type", "description", "inverted"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	prms := &graylog.StreamRuleUpdateParams{}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRuleUpdateParams")
		return nil, 400, err
	}
	lgc.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": prms,
	}).Debug("request body")

	prms.StreamID = streamID
	prms.ID = ruleID
	sc, err = lgc.UpdateStreamRule(prms)
	if err != nil {
		logic.LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "rule": &prms, "status_code": sc,
		}), "faield to update stream rule")
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": prms.ID}, sc, nil
}

// HandleDeleteStreamRule is the handler of Delete a Stream Rule API.
func HandleDeleteStreamRule(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleId} Delete a stream rule
	streamID := ps.PathParam("streamID")
	id := ps.PathParam("streamRuleID")
	// TODO authorization
	sc, err := lgc.DeleteStreamRule(streamID, id)
	return nil, sc, err
}
