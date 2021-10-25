package uploadfileservices

import (
	"bytes"
	"context"
	"fmt"
	"github.com/foody-go-api/aws/s3"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/uploadfiles/uploadfilemodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type UploadServiceInterface interface {
	CreateImage(context context.Context, data *common.Image) error
}

type UploadService struct {
	provider s3.AwsS3Provider
	imgStore UploadServiceInterface
}

func NewUploadService(provider s3.AwsS3Provider, imgStore UploadServiceInterface) *UploadService {
	return &UploadService{provider: provider, imgStore: imgStore}
}

func (biz *UploadService) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadfilemodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadfilemodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	return nil, uploadfilemodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
