package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/mmildd_s/app/ent"
   "github.com/mmildd_s/app/ent/medical"
   "github.com/gin-gonic/gin"
)
 
// MedicalController defines the struct for the medical controller
type MedicalController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateMedical handles POST requests for adding medical entities
// @Summary Create medical
// @Description Create medical
// @ID create-medical
// @Accept   json
// @Produce  json
// @Param medical body ent.Medical true "Medical entity"
// @Success 200 {object} ent.Medical
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicals [post]
func (ctl *MedicalController) CreateMedical(c *gin.Context) {
	obj := ent.Medical{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Medical binding failed",
		})
		return
	}
  
	d, err := ctl.client.Medical.
		Create().
		SetMedicalEmail(obj.MedicalEmail).
		SetMedicalPassword(obj.MedicalPassword).
		SetMedicalName(obj.MedicalName).
		SetMedicalTel(obj.MedicalTel).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, d)
 }
 // GetMedical handles GET requests to retrieve a medical entity
// @Summary Get a medical entity by ID
// @Description get medical by ID
// @ID get-medical
// @Produce  json
// @Param id path int true "Medical ID"
// @Success 200 {object} ent.Medical
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicals/{id} [get]
func (ctl *MedicalController) GetMedical(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	d, err := ctl.client.Medical.
		Query().
		Where(medical.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, d)
 }
 // ListMedical handles request to get a list of medical entities
// @Summary List medical entities
// @Description list medical entities
// @ID list-medical
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medical
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicals [get]
func (ctl *MedicalController) ListMedical(c *gin.Context) {
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
  
	medicals, err := ctl.client.Medical.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, medicals)
 }
 // DeleteMedicalhandles DELETE requests to delete a medical entity
// @Summary Delete a medical entity by ID
// @Description get medical by ID
// @ID delete-medical
// @Produce  json
// @Param id path int true "Medical ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicals/{id} [delete]
func (ctl *MedicalController) DeleteMedical(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Medical.
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
 // UpdateMedical handles PUT requests to update a medical entity
// @Summary Update a medical entity by ID
// @Description update medical by ID
// @ID update-medical
// @Accept   json
// @Produce  json
// @Param id path int true "Medical ID"
// @Param medical body ent.Medical true "Medical entity"
// @Success 200 {object} ent.Medical
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicals/{id} [put]
func (ctl *MedicalController) UpdateMedical(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Medical{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medical binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Medical.
		UpdateOne(&obj).
		SetMedicalEmail(obj.MedicalEmail).
		SetMedicalPassword(obj.MedicalPassword).
		SetMedicalName(obj.MedicalName).
		SetMedicalTel(obj.MedicalTel).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
 // NewMedicalController creates and registers handles for the medical controller
func NewMedicalController(router gin.IRouter, client *ent.Client) *MedicalController {
	uc := &MedicalController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitMedicalController registers routes to the main engine
 func (ctl *MedicalController) register() {
	medicals := ctl.router.Group("/medicals")
  
	medicals.GET("", ctl.ListMedical)
  
	// CRUD
	medicals.POST("", ctl.CreateMedical)
	medicals.GET(":id", ctl.GetMedical)
	medicals.PUT(":id", ctl.UpdateMedical)
	medicals.DELETE(":id", ctl.DeleteMedical)
 }
 
 