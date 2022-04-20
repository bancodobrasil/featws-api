package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bancodobrasil/featws-api/models"
	payloads "github.com/bancodobrasil/featws-api/payloads/v1"
	responses "github.com/bancodobrasil/featws-api/responses/v1"
	"github.com/bancodobrasil/featws-api/services"
	"github.com/gin-gonic/gin"
)

// CreateRulesheet godoc
// @Summary 		Create Rulesheet
// @Description Create Rulesheet description
// @Tags 				rulesheet
// @Accept  		json
// @Produce  		json
// @Param 			rulesheet body payloads.Rulesheet true "Rulesheet body"
// @Success 		200 {object} payloads.Rulesheet
// @Header 			200 {string} Authorization "token access"
// @Failure 		400,404 {object} responses.Error
// @Failure 		500 {object} responses.Error
// @Failure 		default {object} responses.Error
// @Security 		ApiKeyAuth
// @Router 			/rulesheets [post]
func CreateRulesheet() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var payload payloads.Rulesheet
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: err.Error(),
			})
			return
		}

		// use the validator libraty to validate required fields
		if validationErr := validate.Struct(&payload); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: validationErr.Error(),
			})
			return
		}

		entity, err := models.NewRulesheetV1(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		err = services.CreateRulesheet(ctx, &entity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		var response = responses.NewRulesheet(&entity)
		c.JSON(http.StatusCreated, response)
	}
}

// GetRulesheet godoc
// @Summary 		List Rulesheets
// @Description List Rulesheet description
// @Tags 				rulesheet
// @Accept  		json
// @Produce  		json
// @Success 		200 {array} payloads.Rulesheet
// @Header 			200 {string} Authorization "token access"
// @Failure 		400,404 {object} responses.Error
// @Failure 		500 {object} responses.Error
// @Failure 		default {object} responses.Error
// @Security 		ApiKeyAuth
// @Router 			/rulesheets [get]
func GetRulesheets() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		query := c.Request.URL.Query()
		filter := make(map[string]interface{})
		for param, value := range query {
			if len(value) == 1 {
				filter[param] = value[0]
				continue
			}
			filter[param] = value
		}

		entities, err := services.FetchRulesheets(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		var response = make([]responses.Rulesheet, len(entities))

		for index, entity := range entities {
			response[index] = responses.NewRulesheet(entity)
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetRulesheet godoc
// @Summary 		Get Rulesheet by ID
// @Description Get Rulesheet by ID description
// @Tags 				rulesheet
// @Accept  		json
// @Produce  		json
// @Param				id path string true "Rulesheet ID"
// @Success 		200 {array} payloads.Rulesheet
// @Header 			200 {string} Authorization "token access"
// @Failure 		400,404 {object} responses.Error
// @Failure 		500 {object} responses.Error
// @Failure 		default {object} responses.Error
// @Security 		ApiKeyAuth
// @Router 			/rulesheets/{id} [get]
func GetRulesheet() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id, exists := c.Params.Get("id")

		if !exists {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: "Required param 'id'",
			})
			return
		}

		entity, err := services.FetchRulesheet(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		if entity != nil {
			var response = responses.NewRulesheet(entity)

			c.JSON(http.StatusOK, response)
			return
		}

		c.String(http.StatusNotFound, "")
	}
}

// UpdateRulesheet godoc
// @Summary 		Update Rulesheet by ID
// @Description Update Rulesheet by ID description
// @Tags 				rulesheet
// @Accept  		json
// @Produce  		json
// @Param				id path string true "Rulesheet ID"
// @Param 			rulesheet body payloads.Rulesheet true "Rulesheet body"
// @Success 		200 {array} payloads.Rulesheet
// @Header 			200 {string} Authorization "token access"
// @Failure 		400,404 {object} responses.Error
// @Failure 		500 {object} responses.Error
// @Failure 		default {object} responses.Error
// @Security 		ApiKeyAuth
// @Router 			/rulesheets/{id} [put]
func UpdateRulesheet() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id, exists := c.Params.Get("id")

		if !exists {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: "Required param 'id'",
			})
			return
		}

		var payload payloads.Rulesheet
		// validate the request body
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: err.Error(),
			})
		}

		// use the validator libraty to validate required fields
		if validationErr := validate.Struct(&payload); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: validationErr.Error(),
			})
			return
		}

		payload.ID = id

		entity, err := models.NewRulesheetV1(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		updatedEntity, err := services.UpdateRulesheet(ctx, entity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		if updatedEntity != nil {
			var response = responses.NewRulesheet(updatedEntity)

			c.JSON(http.StatusOK, response)
			return
		}

		c.String(http.StatusNotFound, "")
	}
}

// DeleteRulesheet godoc
// @Summary 		Delete Rulesheet by ID
// @Description Delete Rulesheet by ID description
// @Tags 				rulesheet
// @Accept  		json
// @Produce  		json
// @Param				id path string true "Rulesheet ID"
// @Success 		200 {string} string ""
// @Header 			200 {string} Authorization "token access"
// @Failure 		400,404 {object} responses.Error
// @Failure 		500 {object} responses.Error
// @Failure 		default {object} responses.Error
// @Security 		ApiKeyAuth
// @Router 			/rulesheets/{id} [delete]
func DeleteRulesheet() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id, exists := c.Params.Get("id")

		if !exists {
			c.JSON(http.StatusBadRequest, responses.Error{
				Error: "Required param 'id'",
			})
			return
		}

		deleted, err := services.DeleteRulesheet(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{
				Error: err.Error(),
			})
			return
		}

		if !deleted {
			c.String(http.StatusNotFound, "")
			return
		}

		c.String(http.StatusNoContent, "")
	}
}
