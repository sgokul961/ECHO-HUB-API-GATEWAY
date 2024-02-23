package helper

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JwtWrapper struct {
	SecretKey       string
	ExpirationHours int64
}
type jwtClaims struct {
	jwt.StandardClaims
	Id       int64
	Email    string
	Password string
	Username string
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

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
