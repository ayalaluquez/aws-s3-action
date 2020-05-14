package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/credentials"

    "fmt"
    "os"
)
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

func main() {

    // Access Inputs as environment vars
	aws_access_key_id := os.Getenv("aws-access-key-id")
	aws_secret_access_key := os.Getenv ("aws-secret-access-key")
	aws_region := os.Getenv("aws-region")
	aws_bucket := os.Getenv("aws-bucket")	
	filename := os.Getenv("filename")	


	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}
	
	defer file.Close()
	
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id , aws_secret_access_key),
	})
	_, err := sess.Config.Credentials.Get()

	uploader := s3manager.NewUploader(sess)


	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(aws_bucket),
		Key: aws.String(filename),
		Body: file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, aws_bucket, err)
	}
	
	fmt.Printf("Successfully uploaded %q to %q\n", filename, aws_bucket)
}