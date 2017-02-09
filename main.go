package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ejholmes/cloudwatch"
)

func main() {
	var echoInput = false
	var groupName, streamName string

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s [-t] log-group-name log-stream-name\n", os.Args[0])
		os.Exit(1)
	}

	// This should really be cleaned up, but since there's only one optional
	// argument, it's "good enough" for now
	if os.Args[1] == "-t" {
		echoInput = true
		groupName = os.Args[2]
		streamName = os.Args[3]
	} else {
		groupName = os.Args[1]
		streamName = os.Args[2]
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	client := cloudwatchlogs.New(sess)
	if _, err := client.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{LogGroupName: aws.String(groupName)}); err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// It's OK if the log group exists already
			if awsErr.Code() != "ResourceAlreadyExistsException" {
				log.Fatal(awsErr)
			}
		} else {
			log.Fatal(err)
		}
	}

	group := cloudwatch.NewGroup(groupName, client)
	w, err := group.Create(streamName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(w io.Writer) {
		// Ensure we flush any remaining buffered logs to stream
		if writer, ok := w.(*cloudwatch.Writer); ok {
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}
		}
	}(w)

	var r io.Reader

	if echoInput {
		r = io.TeeReader(os.Stdin, os.Stdout)
	} else {
		r = os.Stdin
	}

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
