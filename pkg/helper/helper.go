package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type jwtClaims struct {
	jwt.StandardClaims
	Id   int64
	Role string
}

func AddImageToS3(file *multipart.FileHeader) (string, error) {

	//got an error while uploading the profile picture so loaded the env varibles directly with using viper

	os.Setenv("AWS_ACCESS_KEY_ID", viper.GetString("AWS_ACCESS_KEY_ID"))

	os.Setenv("AWS_SECRET_ACCESS_KEY", viper.GetString("AWS_SECRET_ACCESS_KEY"))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		fmt.Println("configuration error:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("time-peace"),
		Key:    aws.String(file.Filename),
		Body:   f,
		//ACL:    "public-read",
	})

	if uploadErr != nil {
		fmt.Println("uploading error:", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}
func ValidateToken(tokenstring string, key string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenstring, jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method:%v", t.Header["alg"])
		}
		return []byte(key), nil

	})
	return token, err

}
