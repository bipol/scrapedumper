package martaapi

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ScheduleToWriteRequest(s Schedule) (*dynamodb.WriteRequest, error) {
	date, err := time.Parse("1/02/2006 3:04:05 PM", s.EventTime)
	if err != nil {
		return nil, err
	}
	s.PrimaryKey = fmt.Sprintf("%s_%s_%s", s.Station, s.Destination, date.Format("2006-01-02"))
	attr, err := dynamodbattribute.MarshalMap(s)
	if err != nil {
		return nil, err
	}
	return &dynamodb.WriteRequest{
		PutRequest: &dynamodb.PutRequest{
			Item: attr,
		},
	}, nil
}

const (
	PutRequestKey = "PutRequest"
)

func DigestScheduleResponse(r io.Reader) (*dynamodb.BatchWriteItemInput, error) {
	var (
		inp dynamodb.BatchWriteItemInput
	)
	requestItems := make(map[string][]*dynamodb.WriteRequest)

	dec := json.NewDecoder(r)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return nil, err
	}

	for dec.More() {
		var s Schedule
		err = dec.Decode(&s)
		if err != nil {
			return nil, err
		}
		wr, err := ScheduleToWriteRequest(s)
		if err != nil {
			return nil, err
		}
		requestItems[PutRequestKey] = append(requestItems[PutRequestKey], wr)
	}

	_, err = dec.Token()
	if err != nil {
		return nil, err
	}

	return inp.SetRequestItems(requestItems), err
}
