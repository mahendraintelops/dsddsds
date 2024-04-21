package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/models"
	"github.com/mahendraintelops/dsddsds/dddds/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type DadadController struct {
	dadadService *services.DadadService
}

func NewDadadController() (*DadadController, error) {
	dadadService, err := services.NewDadadService()
	if err != nil {
		return nil, err
	}
	return &DadadController{
		dadadService: dadadService,
	}, nil
}

func (dadadController *DadadController) CreateDadad(context *gin.Context) {
	// validate input
	var input models.Dadad
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger dadad creation
	dadadCreated, err := dadadController.dadadService.CreateDadad(&input)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, dadadCreated)
}

func (dadadController *DadadController) FetchDadad(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger dadad fetching
	dadad, err := dadadController.dadadService.GetDadad(id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceName := os.Getenv("SERVICE_NAME")
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// get the current span by the request context
		currentSpan := trace.SpanFromContext(context.Request.Context())
		currentSpan.SetAttributes(attribute.String("dadad.id", strconv.FormatInt(dadad.Id, 10)))
	}

	context.JSON(http.StatusOK, dadad)
}
