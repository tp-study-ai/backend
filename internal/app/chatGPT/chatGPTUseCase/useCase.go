package chatGPTUseCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	che "github.com/tp-study-ai/backend/internal/app/chatGPT"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
)

type UseCaseChatGPT struct {
	Repo    che.Repository
	Secret1 string
	Secret2 string
	Secret3 string
	Secret4 string
	Secret5 string
}

func NewUseCaseChatGPT(TaskRepo che.Repository, secret string, secret1 string, secret2 string, secret3 string, secret4 string) *UseCaseChatGPT {
	return &UseCaseChatGPT{
		Repo:    TaskRepo,
		Secret1: secret,
		Secret2: secret1,
		Secret3: secret2,
		Secret4: secret3,
		Secret5: secret4,
	}
}

func (u *UseCaseChatGPT) Chat(Message models.ChatGPT) (*models.Message, error) {
	fmt.Println(Message.TaskId)
	task, err := u.Repo.GetTaskForChatGPT(Message.TaskId)
	if err != nil {
		return nil, err
	}

	MessageRequest := &models.ChatGPTRequest{
		UserMessage:    Message.Message,
		Statement:      task.Description,
		UserSolution:   Message.Code,
		MasterSolution: task.MasterSolution,
	}

	result, err := json.Marshal(MessageRequest)
	if err != nil {
		return nil, err
	}

	req := bytes.NewBuffer(result)
	resp, err := http.Post(u.Secret5, "application/json", req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var ChatGPTResponse models.Message

	err = json.Unmarshal(body, &ChatGPTResponse)
	if err != nil {
		return nil, errors.Errorf("1511 " + err.Error())
	}

	return &ChatGPTResponse, nil
}
