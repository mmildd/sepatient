package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/mmildd_s/app/ent"
   "github.com/mmildd_s/app/ent/certificate"
   "github.com/gin-gonic/gin"
)
 
// CertificateController defines the struct for the certificate controller
type CertificateController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateCertificate handles POST requests for adding certificate entities
// @Summary Create certificate
// @Description Create certificate
// @ID create-certificate
// @Accept   json
// @Produce  json
// @Param certificate body ent.Certificate true "Certificate entity"
// @Success 200 {object} ent.Certificate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certificates [post]
func (ctl *CertificateController) CreateCertificate(c *gin.Context) {
	obj := ent.Certificate{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Certificate binding failed",
		})
		return
	}
  
	me, err := ctl.client.Certificate.
		Create().
		SetCertificateName(obj.CertificateName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, me)
 }
 // GetCertificate handles GET requests to retrieve a certificate entity
// @Summary Get a certificate entity by ID
// @Description get certificate by ID
// @ID get-certificate
// @Produce  json
// @Param id path int true "Certificate ID"
// @Success 200 {object} ent.Certificate
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certificates/{id} [get]
func (ctl *CertificateController) GetCertificate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	me, err := ctl.client.Certificate.
		Query().
		Where(certificate.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, me)
 }
 // ListCertificate handles request to get a list of certificate entities
// @Summary List certificate entities
// @Description list certificate entities
// @ID list-certificate
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Certificate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certificates [get]
func (ctl *CertificateController) ListCertificate(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	certificates, err := ctl.client.Certificate.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, certificates)
 }
 // DeleteCertificate handles DELETE requests to delete a certificate entity
// @Summary Delete a certificate entity by ID
// @Description get certificate by ID
// @ID delete-certificate
// @Produce  json
// @Param id path int true "Certificate ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certificates/{id} [delete]
func (ctl *CertificateController) DeleteCertificate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Certificate.
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
 // UpdateCertificate handles PUT requests to update a certificate entity
// @Summary Update a certificate entity by ID
// @Description update certificate by ID
// @ID update-certificate
// @Accept   json
// @Produce  json
// @Param id path int true "Certificate ID"
// @Param certificate body ent.Certificate true "Certificate entity"
// @Success 200 {object} ent.Certificate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certificates/{id} [put]
func (ctl *CertificateController) UpdateCertificate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Certificate{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Certificate binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	me, err := ctl.client.Certificate.
		UpdateOne(&obj).
		SetCertificateName(obj.CertificateName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, me)
 }
 // NewCertificateController creates and registers handles for the certificate controller
func NewCertificateController(router gin.IRouter, client *ent.Client) *CertificateController {
	uc := &CertificateController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitCertificateController registers routes to the main engine
 func (ctl *CertificateController) register() {
	certificates := ctl.router.Group("/certificates")
  
	certificates.GET("", ctl.ListCertificate)
  
	// CRUD
	certificates.POST("", ctl.CreateCertificate)
	certificates.GET(":id", ctl.GetCertificate)
	certificates.PUT(":id", ctl.UpdateCertificate)
	certificates.DELETE(":id", ctl.DeleteCertificate)
 }
 
 