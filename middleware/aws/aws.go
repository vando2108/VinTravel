package aws

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
  AWS_REGION = "ap-southeast-1"
  AWS_BUCKET = "vintravel"
)

func Init() (*session.Session, error) {
  session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_REGION)})
  return session, err
}

func Upload(session *session.Session, fileDir string) (string, error) {
  upFile, err := os.Open(fileDir)
  if err != nil {
    return "", err
  }
  defer upFile.Close()

  fileInfo, _ := upFile.Stat()
  var fileSize int64 = fileInfo.Size()
  buffer := make([]byte, fileSize)
  upFile.Read(buffer)
  fmt.Println(buffer)

  _, err = s3.New(session).PutObject(&s3.PutObjectInput{
    Bucket: aws.String(AWS_BUCKET),
    Key: aws.String(fileInfo.Name()),
    ACL: aws.String("public-read"),
    Body: bytes.NewReader(buffer),
    ContentLength: aws.Int64(fileSize),
    ContentType: aws.String(http.DetectContentType(buffer)),
    ContentDisposition: aws.String("attachment"),
    ServerSideEncryption: aws.String("AES256"),
  }) 
  
  url := "https://" + AWS_BUCKET + ".s3-" + AWS_REGION + ".amazonaws.com/" + fileInfo.Name();
  return url, err
}
