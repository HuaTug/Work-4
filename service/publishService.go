package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type UploadVideoService struct {
	ContentType string `form:"content_type" json:"content_type"`
	ObjectName  string `form:"object_name" json:"object_name"`
	BucketName  string `form:"bucket_name" json:"bucket_name"`
}

func (service *UploadVideoService) UploadVideo(file *multipart.FileHeader, author uint) serializer.Response {
	var video model.Video
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	minioClient, err := minio.New("127.0.0.1:9000", &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	fmt.Println(file.Header)
	fmt.Println(file.Filename)
	bucketName := service.BucketName
	objectName := service.ObjectName
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(context.Background(), bucketName)
		if err == nil && exists {
			log.Printf("Bucket:%s is already exist\n", bucketName)
		} else {
			return serializer.Response{
				Status: 500,
				Msg:    "err",
			}
		}
	}
	filePath := "C:\\Users\\0\\Downloads\\Video\\" + file.Filename
	//ToDo 如何实现任意的文件上传并且获得其绝对路径
	src, err1 := os.Open(filePath)
	if err1 != nil {
		fmt.Println("Open文件出错")
		log.Fatalln(err)
	}
	defer src.Close()
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	publish := &model.Video{
		Title:       file.Filename,
		PublishTime: time.Now(),
		AuthorId:    author,
		CoverUrl:    filePath,
		PlayUrl:     filePath,
	}
	log.Println("视频文件上传成功")
	if err := model.DB.Model(&video).Create(publish).Error; err != nil {
		panic(err)
	}
	return serializer.Response{
		Status: 200,
		Msg:    "视频文件上传成功",
	}
}
