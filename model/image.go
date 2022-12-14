package model

import (
	"my-imgur/lib/pcloud"
	"net/url"
	"os"
)

type Image struct {
	pCloudClient pcloud.IClient
}

func (i *Image) GetPublicThumbnailLink(fileName string, fileId int, width int, height int) (link string, err error) {
	link, err = i.pCloudClient.GetPublicThumbnail("/Public Asset/obsidian/"+fileName, fileId, width, height)
	return
}

func (i *Image) UploadFile(fileName string, data []byte) (path string, err error) {
	res, err := i.pCloudClient.UploadFile(pcloud.OBSIDIAN, fileName, data, pcloud.UploadFileOption{RenameIfExists: true})
	if err != nil {
		return "", err
	}
	path = os.Getenv("PUBLIC_DOMAIN") + "/v1/temp-link/obsidian/" + url.PathEscape(res.Metadata[0].Name)
	//path = os.Getenv("PUBLIC_DOMAIN") + "/obsidian/" + res.Metadata[0].Name
	return
}

func NewImage() IImage {
	client := pcloud.NewClient()
	return &Image{
		pCloudClient: client,
	}
}

//go:generate mockgen -destination=image_mock.go -package=model . IImage
type IImage interface {
	UploadFile(fileName string, data []byte) (path string, err error)
	GetPublicThumbnailLink(fileName string, fileId int, width int, height int) (link string, err error)
}
