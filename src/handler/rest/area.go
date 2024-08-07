package restserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/gin-gonic/gin"
	x "github.com/pkg/errors"
)

func (e *rest) GetAreaByID(c *gin.Context) {

	id := c.Param("id")

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error Parsing Area ID"))
		return
	}

	result, err := e.uc.GetAreaByID(c, aid)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, err)
		return
	}

	e.HttpSuccess(c, http.StatusOK, result, nil)
}

func (e *rest) CreateArea(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error reading body payload"))
		return
	}

	var params entity.AreaInput

	err = json.Unmarshal(data, &params)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error parsing body payload"))
		return
	}

	result, err := e.uc.CreateArea(c, params)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, err)
		return
	}

	e.HttpSuccess(c, http.StatusCreated, result, nil)
}

func (e *rest) GetArea(c *gin.Context) {

	params := entity.AreaParams{}

	err := c.Bind(&params)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, x.Wrap(err, "Error reading params payload"))
		return
	}

	result, p, err := e.uc.GetArea(c, params)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, err)
		return
	}

	e.HttpSuccess(c, http.StatusOK, result, &p)
}

func (e *rest) UpdateArea(c *gin.Context) {

	id := c.Param("id")

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error reading Area ID"))
		return
	}

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error reading body payload"))
		return
	}

	var params entity.UpdateAreaInput

	err = json.Unmarshal(data, &params)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, x.Wrap(err, "Error parsing body payload"))
		return
	}

	params.ID = aid

	result, err := e.uc.UpdateArea(c, params)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, err)
		return
	}

	e.HttpSuccess(c, http.StatusCreated, result, nil)
}

func (e *rest) DeleteArea(c *gin.Context) {

	id := c.Param("id")

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		e.HttpError(c, http.StatusBadRequest, err)
		return
	}

	err = e.uc.DeleteArea(c, aid)
	if err != nil {
		e.HttpError(c, http.StatusInternalServerError, err)
		return
	}

	e.HttpSuccess(c, http.StatusCreated, nil, nil)
}
