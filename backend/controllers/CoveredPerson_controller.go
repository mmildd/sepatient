package controllers
 
import (
	"context"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mmildd_s/app/ent"
	"github.com/mmildd_s/app/ent/patient"
	"github.com/mmildd_s/app/ent/schemetype"
	"github.com/mmildd_s/app/ent/fund"
	"github.com/mmildd_s/app/ent/certificate"
	"github.com/mmildd_s/app/ent/coveredperson"
)
 // CoveredPersonController defines the struct for the CoveredPerson controller
type CoveredPersonController struct {
	client *ent.Client
	router gin.IRouter
}
//CoveredPerson struct
type CoveredPerson struct {
	Patient   		int
	SchemeType  	int
	Fund 			int
	Certificate     int
}
// CreateCoveredPerson handles POST requests for adding CoveredPerson entities
// @Summary Create coveredperson
// @Description Create coveredperson
// @ID create-coveredperson
// @Accept   json
// @Produce  json
// @Param coveredperson body CoveredPerson true "CoveredPerson entity"
// @Success 200 {object} ent.CoveredPerson
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coveredpersons [post]
func (ctl *CoveredPersonController) CreateCoveredPerson(c *gin.Context) {
	obj := CoveredPerson{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "CoveredPerson binding failed",
		})
		return
	}
	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}
	st, err := ctl.client.SchemeType.
		Query().
		Where(schemetype.IDEQ(int(obj.SchemeType))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "SchemeType not found",
		})
		return
	}
	f, err := ctl.client.Fund.
		Query().
		Where(fund.IDEQ(int(obj.Fund))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Fund not found",
		})
		return
	}
	ce, err := ctl.client.Certificate.
		Query().
		Where(certificate.IDEQ(int(obj.Certificate))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Certificate not found",
		})
		return
	}
	
	cp, err := ctl.client.CoveredPerson.
		Create().
		SetPatient(p).
		SetSchemeType(st).
		SetFund(f).
		SetCertificate(ce).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return 
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   cp,
	})
}

// GetCoveredPerson handles GET requests to retrieve a CoveredPerson entity
// @Summary Get a coveredperson entity by ID
// @Description get coveredperson by ID
// @ID get-coveredperson
// @Produce  json
// @Param id path int true "CoveredPerson ID"
// @Success 200 {object} ent.CoveredPerson
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coveredpersons/{id} [get]
func (ctl *CoveredPersonController) GetCoveredPerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.CoveredPerson.
		Query().
		Where(coveredperson.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, b)
}
 // ListCoveredPerson handles request to get a list of CoveredPerson entities
// @Summary List coveredperson entities
// @Description list coveredperson entities
// @ID list-coveredperson
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CoveredPerson
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coveredpersons [get]
func (ctl *CoveredPersonController) ListCoveredPerson(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)}
	}
  
	coveredpersons, err := ctl.client.CoveredPerson.
		Query().
		WithPatient().
		WithSchemeType().
		WithFund().
		WithCertificate().
		Limit(limit).
		Offset(offset).
		All(context.Background())

		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, coveredpersons)
	
 }

// DeleteCoveredPerson handles DELETE requests to delete a CoveredPerson entity
// @Summary Delete a coveredperson entity by ID
// @Description get coveredperson by ID
// @ID delete-coveredperson
// @Produce  json
// @Param id path int true "CoveredPerson ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coveredpersons/{id} [delete]
func (ctl *CoveredPersonController) DeleteCoveredPerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.CoveredPerson.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}
// UpdateCoveredPerson handles PUT requests to update a CoveredPerson entity
// @Summary Update a coveredperson entity by ID
// @Description update coveredperson by ID
// @ID update-coveredperson
// @Accept   json
// @Produce  json
// @Param id path int true "CoveredPerson ID"
// @Param coveredperson body ent.CoveredPerson true "CoveredPerson entity"
// @Success 200 {object} ent.CoveredPerson
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coveredpersons/{id} [put]
func (ctl *CoveredPersonController) UpdateCoveredPerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.CoveredPerson{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "CoveredPerson binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.CoveredPerson.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, b)
}
 
 // NewCoveredPersonController creates and registers handles for the coveredperson controller
func NewCoveredPersonController(router gin.IRouter, client *ent.Client) *CoveredPersonController {
	da := &CoveredPersonController{
		client: client,
		router: router,
	}

	da.register()

	return da

}

func (ctl *CoveredPersonController) register() {
	coveredpersons := ctl.router.Group("/coveredpersons")

	coveredpersons.POST("", ctl.CreateCoveredPerson)
	coveredpersons.GET("", ctl.ListCoveredPerson)
	coveredpersons.DELETE(":id", ctl.DeleteCoveredPerson)

}
 
  