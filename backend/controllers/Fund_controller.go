package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/mmildd_s/app/ent"
   "github.com/mmildd_s/app/ent/fund"
   "github.com/gin-gonic/gin"
)
 
// FundController defines the struct for the fund controller
type FundController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateFund handles POST requests for adding fund entities
// @Summary Create fund
// @Description Create fund
// @ID create-fund
// @Accept   json
// @Produce  json
// @Param fund body ent.Fund true "Fund entity"
// @Success 200 {object} ent.Fund
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /funds [post]
func (ctl *FundController) CreateFund(c *gin.Context) {
	obj := ent.Fund{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fund binding failed",
		})
		return
	}
  
	f, err := ctl.client.Fund.
		Create().
		SetFundName(obj.FundName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, f)
 }
 // GetFund handles GET requests to retrieve a fund entity
// @Summary Get a fund entity by ID
// @Description get fund by ID
// @ID get-fund
// @Produce  json
// @Param id path int true "Fund ID"
// @Success 200 {object} ent.Fund
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /funds/{id} [get]
func (ctl *FundController) GetFund(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	f, err := ctl.client.Fund.
		Query().
		Where(fund.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, f)
 }
 // ListFund handles request to get a list of fund entities
// @Summary List fund entities
// @Description list fund entities
// @ID list-fund
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Fund
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /funds [get]
func (ctl *FundController) ListFund(c *gin.Context) {
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
  
	funds, err := ctl.client.Fund.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, funds)
 }
 // DeleteFund handles DELETE requests to delete a fund entity
// @Summary Delete a fund entity by ID
// @Description get fund by ID
// @ID delete-fund
// @Produce  json
// @Param id path int true "fund ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /funds/{id} [delete]
func (ctl *FundController) DeleteFund(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Fund.
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
 // UpdateFund handles PUT requests to update a fund entity
// @Summary Update a fund entity by ID
// @Description update fund by ID
// @ID update-fund
// @Accept   json
// @Produce  json
// @Param id path int true "fund ID"
// @Param fund body ent.Fund true "fund entity"
// @Success 200 {object} ent.Fund
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /funds/{id} [put]
func (ctl *FundController) UpdateFund(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Fund{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fund binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Fund.
		UpdateOne(&obj).
		SetFundName(obj.FundName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, f)
 }
 // NewFundController creates and registers handles for the fund controller
func NewFundController(router gin.IRouter, client *ent.Client) *FundController {
	uc := &FundController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitFundController registers routes to the main engine
 func (ctl *FundController) register() {
	funds := ctl.router.Group("/funds")
  
	funds.GET("", ctl.ListFund)
	// CRUD
	funds.POST("", ctl.CreateFund)
	funds.GET(":id", ctl.GetFund)
	funds.PUT(":id", ctl.UpdateFund)
	funds.DELETE(":id", ctl.DeleteFund)
 }
 
 