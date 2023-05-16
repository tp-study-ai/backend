package chatGPTUseCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/chatGPT"
	"github.com/tp-study-ai/backend/internal/app/models"
	"io/ioutil"
	"net/http"
)

type UseCaseChatGPT struct {
	Repo     chatGPT.Repository
	CGSecret string
}

func NewUseCaseChatGPT(TaskRepo chatGPT.Repository, secret string) *UseCaseChatGPT {
	return &UseCaseChatGPT{
		Repo:     TaskRepo,
		CGSecret: secret,
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
	resp, err := http.Post(u.CGSecret, "application/json", req)
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
