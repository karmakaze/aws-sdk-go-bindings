package sqs

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreateQueueInput(t *testing.T) {

	out, err := NewCreateQueueInput(queueName)

	assert.NoError(t, err)
	assert.Equal(t, queueName, *out.CreateQueueInput.QueueName)

	_, shouldBeErrEmptyParameter := NewCreateQueueInput("")

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, shouldBeErrEmptyParameter.Error())

}

func TestNewGetQueueAttributesInput(t *testing.T) {

	out, err := NewGetQueueAttributesInput(queueName)

	assert.NoError(t, err)
	assert.Equal(t, queueName, *out.GetQueueAttributesInput.QueueUrl)

	_, shouldBeEmptyErr := NewGetQueueAttributesInput("")

	assert.Error(t, shouldBeEmptyErr)
	assert.Equal(t, ErrEmptyParameter, shouldBeEmptyErr.Error())

}

func TestNewSendMessageInput(t *testing.T) {

	queueUrl := "some_url"
	s := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	out, err := NewSendMessageInput(
		s,
		queueUrl,
	)

	assert.NoError(t, err)
	assert.Equal(t, queueUrl, *out.SendMessageInput.QueueUrl)
	assert.Contains(t, *out.SendMessageInput.MessageBody, val1, val2)

	var m TestSQSUtilType

	unmarshalErr := json.Unmarshal([]byte(*out.SendMessageInput.MessageBody), &m)

	assert.NoError(t, unmarshalErr)
	assert.Equal(t, val1, m.SomeParam1)
	assert.Equal(t, val2, m.SomeParam2)

	_, shouldBeErrEmptyParameter := NewSendMessageInput(
		s,
		"",
	)

	assert.Error(t, shouldBeErrEmptyParameter)
	assert.Equal(t, ErrEmptyParameter, shouldBeErrEmptyParameter.Error())

	_, shouldBeErrNoPointerParameterAllowed := NewSendMessageInput(
		&s,
		queueUrl,
	)

	assert.Error(t, shouldBeErrNoPointerParameterAllowed)
	assert.Equal(t, ErrNoPointerParameterAllowed, shouldBeErrNoPointerParameterAllowed.Error())

}

func TestMarshalStructToJsonString(t *testing.T) {

	s := TestSQSUtilType{
		SomeParam1: val1,
		SomeParam2: val2,
	}

	res, err := marshalStructToJsonString(s)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Contains(t, res, val1, val2)

	var out TestSQSUtilType

	unmarshalErr := json.Unmarshal([]byte(res), &out)

	assert.NoError(t, unmarshalErr)
	assert.Equal(t, val1, out.SomeParam1)
	assert.Equal(t, val2, out.SomeParam2)

	_, shouldBeErrNoPointerParameterAllowed := marshalStructToJsonString(&s)

	assert.Error(t, shouldBeErrNoPointerParameterAllowed)
	assert.Equal(t, ErrNoPointerParameterAllowed, shouldBeErrNoPointerParameterAllowed.Error())

}