package handler

import (
	brantect_local "brantect-api-tm/brantect"
	"net/http"
	"strconv"

	"github.com/BrightsHoldings/gmobrs-go-lib/errors"
	"github.com/BrightsHoldings/gmobrs-go-lib/log"
	"github.com/BrightsHoldings/gmobrs-go-lib/where"
	"github.com/labstack/echo/v4"
)

// Search implements interfaces.TmRightsSystemHandlerInterface.
func (s *TmRightsSystemHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Debugf("TmRightsSystemHandler > Search processing... \n")

		// init variables
		var (
			searchAttorney string
			conditions     []where.ConditionInterface
			conditionOr    where.ConditionInterface
			resp           *brantect_local.TmRightsSystemSearchResponse
		)

		// init response
		resp = brantect_local.NewTmRightsSystemSearchResponse(c.Path())
		resp.Data = nil
		//get filter params
		filterParams := new(brantect_local.TmRightsSystemSearchFilterParam)
		offSetStr := c.QueryParam("offset")
		offSet, errOffset := strconv.Atoi(offSetStr)
		if errOffset != nil {
			offSet = 0
		}
		filterParams.Offset = offSet

		limitStr := c.QueryParam("limit")
		limit, errLimit := strconv.Atoi(limitStr)
		if errLimit != nil {
			limit = 0 // default limit
		}
		filterParams.Limit = limit
		filterParams.ClientCd = c.QueryParam("client_cd")
		filterParams.TmName = ReplaceSpecialChar(c.QueryParam("tm_name"))
		filterParams.AppNo = c.QueryParam("app_no")
		filterParams.AppDate = c.QueryParam("app_date")

		// handle search attorney
		searchAttorney = c.QueryParam("search_attorney")
		log.Debugf("TmRightsSystemHandler > Search searchAttorney: %s \n", searchAttorney)

		// validate searchAttorney
		if searchAttorney == "" && c.QueryParams().Has("search_attorney") {
			resp.Status.Code = http.StatusBadRequest
			resp.Status.Type = http.StatusText(http.StatusBadRequest)
			resp.Errors = append(resp.Errors, errors.New(
				2030010003001,
				"search_attorney cannot be empty",
				searchAttorney,
			))
			return c.JSON(http.StatusBadRequest, resp)
		}

		if searchAttorney != "" {
			validateClientCd := validateClientCd(searchAttorney)
			if !validateClientCd {
				resp.Status.Code = http.StatusBadRequest
				resp.Status.Type = http.StatusText(http.StatusBadRequest)
				resp.Errors = append(resp.Errors, errors.New(
					2030010003002,
					"search_attorney is invalid (not empty, 11 characters and numeric)",
					searchAttorney,
				))
				return c.JSON(http.StatusBadRequest, resp)
			}

			// make condition for searchAttorney
			attorneyCondition := where.Or(
				&where.Condition{
					Target:    "dom_attorney_client_cd",
					Condition: "=",
					Value:     searchAttorney,
				},
				&where.Condition{
					Target:    "int_attorney_client_cd",
					Condition: "=",
					Value:     searchAttorney,
				},
			)
			conditions = append(conditions, attorneyCondition)
		}

		if len(conditions) > 0 {
			conditionOr = where.Or(conditions...)
		}

		tx, errBeginTransaction := s.tmRightsSystemService.BrantectRepository().Begin()
		if errBeginTransaction != nil {
			log.Debugf("TmRightsSystemHandler > Search errBeginTransaction: %v. \n", errBeginTransaction)

			resp.Status.Code = http.StatusInternalServerError
			resp.Status.Type = http.StatusText(http.StatusInternalServerError)
			resp.Errors = append(resp.Errors, errors.New(
				2030010003004,
				"Error init transaction",
				nil,
			))
			return c.JSON(http.StatusInternalServerError, resp)
		}

		defer s.tmRightsSystemService.BrantectRepository().Rollback(tx)

		// call service to search tmRights
		tmRights, totalCount, errsSearch := s.tmRightsSystemService.Search(tx, conditionOr, *filterParams)
		if errsSearch != nil {
			log.Debugf("TmImagesUserHandler > Search error. \n")
			resp.Status.Code = http.StatusInternalServerError
			resp.Status.Type = http.StatusText(http.StatusInternalServerError)
			resp.Errors = append(resp.Errors, errsSearch...)
			return c.JSON(http.StatusInternalServerError, resp)
		}

		// set response data
		resp.Data = tmRights
		resp.ContentRange = calculateContentRange(filterParams.Offset, filterParams.Limit, totalCount)

		// tx commit
		if errCommit := s.tmRightsSystemService.BrantectRepository().Commit(tx); errCommit != nil {
			log.Debugf("TmRightsSystemHandler > Search errCommit: %v. \n", errCommit)
			resp.Status.Code = http.StatusInternalServerError
			resp.Status.Type = http.StatusText(http.StatusInternalServerError)
			resp.Errors = append(resp.Errors, errors.New(
				2030010003005,
				"Error commit transaction",
				nil,
			))
			return c.JSON(http.StatusInternalServerError, resp)
		}

		return c.JSON(http.StatusOK, resp)
	}
}
