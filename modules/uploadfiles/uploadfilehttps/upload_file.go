package uploadfilehttps

import (
	"github.com/foody-go-api/aws/s3"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/uploadfiles/uploadfilerepo"
	"github.com/foody-go-api/modules/uploadfiles/uploadfileservices"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(s3 s3.AwsS3Provider, db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		conn := uploadfilerepo.NewSqlConn(db)

		biz := uploadfileservices.NewUploadService(s3, conn)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.NewSuccessResponseNoPaging(img))
	}
}
