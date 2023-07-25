package miniocontroller

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Nunu-Nugroho/golang-first-project/package"
	"github.com/gin-gonic/gin"
	// "golang.org/x/tools/go/packages"
	minioPackage "github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
)

func Test(c *gin.Context) {
	// var client minio.MinioConnection
	client, err := minio.MinioConnection()
	ctx := context.Background()
	buckets, err := client.ListBuckets(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to connect to MinIO server",
		})
		return
	}
	var bucketNames []string
	for _, bucket := range buckets {
		bucketNames = append(bucketNames, bucket.Name)
	}
	c.JSON(http.StatusOK, gin.H{"buckets": bucketNames})
}
func CreateImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Image not found",
		})
		return
	}

	// Ambil nama folder dari form data "folder"
	folder := c.PostForm("folder")
	minioClient, err := minio.MinioConnection()
	// Buat folder di dalam bucket (opsional, jika folder belum ada)
	err = minioClient.MakeBucket(c, os.Getenv("MINIO_BUCKET_NAME"), minioPackage.MakeBucketOptions{
		Region: "object",
	})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(c, "object")
		if errBucketExists == nil && exists {
			log.Println("Bucket already exists.")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create bucket.",
			})
			return
		}
	}

	// Buka file yang diunggah
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to open the image",
		})
		return
	}
	defer src.Close()

	// Simpan file ke folder di bucket di server MinIO
	_, err = minioClient.PutObject(c, "object", folder+"/"+file.Filename, src, file.Size, minioPackage.PutObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload the image to MinIO",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Image uploaded successfully",
	})

}

func GetImage(c *gin.Context) {
	imagePath := c.Param("path")
	bucketName := os.Getenv("MINIO_BUCKET_NAME") // Ganti dengan nama bucket MinIO Anda
	minioClient, err := minio.MinioConnection()
	ctx := context.Background()
	// Dapatkan objek dari bucket MinIO

	obj, err := minioClient.GetObject(ctx, bucketName, imagePath, minioPackage.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gambar tidak ditemukan"})
		return
	}
	defer obj.Close()

	// Set header response untuk tipe konten gambar
	c.Header("Content-Type", "image/jpeg") // Ganti sesuai tipe gambar Anda

	// Salin isi objek MinIO ke ResponseWriter
	_, err = io.Copy(c.Writer, obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim gambar"})
		return
	}
}

func ListContent(c *gin.Context) {
	minioClient, err := minio.MinioConnection()
	// Dapatkan daftar objek di bucket di server MinIO
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	bucketObjects := minioClient.ListObjects(ctx, bucketName, minioPackage.ListObjectsOptions{
		Recursive: false,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list Gambar tidak ditemukan"})
		return
	}
	// defer bucketObjects.Close()

	var contentsList []gin.H
	for object := range bucketObjects {
		if object.Err != nil {
			log.Println("Error:", object.Err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to list contents",
			})
			return
		}

		// Skip the folder itself if it is the root folder
		if object.Key != "" {
			// Check if the object is a folder
			isFolder := strings.HasSuffix(object.Key, "/")

			contentsList = append(contentsList, gin.H{
				"type":     "folder",
				"name":     object.Key,
				"isFolder": isFolder,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"contents": contentsList,
	})
}

func ListContentbyFolder(c *gin.Context) {
	// Ambil nama folder dari parameter URL
	folder := c.Param("folder")
	minioClient, err := minio.MinioConnection()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list Content tidak ditemukan"})
		return
	}
	// Pastikan folder memiliki trailing slash ("/") jika tidak kosong
	if folder != "" && !strings.HasSuffix(folder, "/") {
		folder += "/"
	}

	// Dapatkan daftar objek di bucket di server MinIO
	bucketObjects := minioClient.ListObjects(c, "object", minioPackage.ListObjectsOptions{
		Prefix:    folder,
		Recursive: false,
	})

	var contentsList []gin.H
	for object := range bucketObjects {
		if object.Err != nil {
			log.Println("Error:", object.Err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to list contents",
			})
			return
		}

		// Jika object adalah folder pada folder yang dimaksud, tambahkan ke daftar folder
		if strings.HasPrefix(object.Key, folder) {
			objectName := strings.TrimPrefix(strings.TrimPrefix(object.Key, folder), "/")

			// Skip the object itself (the folder) if it is the same as the folder parameter
			if objectName != "" {
				// Check if the object is a folder
				isFolder := strings.HasSuffix(object.Key, "/")

				contentsList = append(contentsList, gin.H{
					"type":     "folder",
					"name":     objectName,
					"isFolder": isFolder,
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"contents": contentsList,
	})
}