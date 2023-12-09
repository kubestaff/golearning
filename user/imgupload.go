package user

import (

	"net/http"
	"mime/multipart"
	"github.com/gin-gonic/gin"
)

type UserImage struct {
	Image      *multipart.FileHeader `form:"image" binding:"required"`
 }

func Upload(c *gin.Context) {
	body := UserImage{}
   
	//to bind the request payload to the userProfileDTO struct instance.
	if err := c.ShouldBind(&body); err != nil {
	 c.AbortWithStatusJSON(400, "Bad Request")
	}
	//save the uploaded file to the specified folder "asset/".
	err := c.SaveUploadedFile(body.Image, "static/uploads/"+body.Image.Filename)
	if err != nil {
	 c.JSON(http.StatusInternalServerError, "Something went wrong!")
	 return
	}
   
	c.JSON(200, gin.H{
	 "image": body,
	})
   }