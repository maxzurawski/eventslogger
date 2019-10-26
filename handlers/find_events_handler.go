package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/maxzurawski/eventslogger/publishers"

	"github.com/labstack/echo"
	"github.com/maxzurawski/eventslogger/dto"
	"github.com/maxzurawski/eventslogger/service"
	"github.com/maxzurawski/utilities/stringutils"
)

func FindEventsHandler(c echo.Context) error {
	searchDto := &dto.SearchDTO{}

	processId := c.QueryParam("processId")
	if !stringutils.IsZero(processId) {
		searchDto.ProcessId = processId
	}

	topic := c.QueryParam("topic")
	if !stringutils.IsZero(topic) {
		searchDto.Topic = topic
	}

	routingKey := c.QueryParam("routingKey")
	if !stringutils.IsZero(routingKey) {
		searchDto.RoutingKey = routingKey
	}

	errorMsg := c.QueryParam("errorMsg")
	if !stringutils.IsZero(errorMsg) {
		searchDto.ErrorMsg = errorMsg
	}

	sensorsUuid := c.QueryParam("sensorsUuid")
	if !stringutils.IsZero(sensorsUuid) {
		searchDto.SensorsUuid = sensorsUuid
	}

	serviceParam := c.QueryParam("service")
	if !stringutils.IsZero(serviceParam) {
		searchDto.Service = serviceParam
	}

	publishedOnFrom := c.QueryParam("publishedOnFrom")
	if !stringutils.IsZero(publishedOnFrom) {
		if parse, err := time.Parse(time.RFC3339, publishedOnFrom); err != nil {
			publishers.Logger().Warn(uuid.New().String(), "", "could not parse [publishedOnFrom]")
		} else {
			searchDto.PublishedOnFrom = &parse
		}
	}

	publishedOnTo := c.QueryParam("publishedOnTo")
	if !stringutils.IsZero(publishedOnTo) {
		if parse, err := time.Parse(time.RFC3339, publishedOnTo); err != nil {
			publishers.Logger().Warn(uuid.New().String(), "", "could not parse [publishedOnTo]")
		} else {
			searchDto.PublishedOnTo = &parse
		}
	}

	result, err := service.Service.FindBy(*searchDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(result) == 0 {
		return c.JSON(http.StatusNoContent, result)
	}

	return c.JSON(http.StatusOK, result)

}
