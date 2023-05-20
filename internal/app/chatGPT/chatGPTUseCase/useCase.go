package chatGPTUseCase

import (
	"fmt"
	"github.com/tp-study-ai/backend/internal/app/chatGPT"
	"github.com/tp-study-ai/backend/internal/app/models"
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

	fmt.Println(Message)

	//task, err := u.Repo.GetTaskForChatGPT(Message.TaskId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//MessageRequest := &models.ChatGPTRequest{
	//	UserMessage:    Message.Message,
	//	Statement:      task.Description,
	//	UserSolution:   Message.Code,
	//	MasterSolution: task.MasterSolution,
	//}
	//
	//result, err := json.Marshal(MessageRequest)
	//if err != nil {
	//	return nil, err
	//}
	//
	//req := bytes.NewBuffer(result)
	//resp, err := http.Post(u.CGSecret, "application/json", req)
	//if err != nil {
	//	return nil, err
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}

	ChatGPTResponse := &models.Message{
		Message: "Привет! Рада видеть, что ты решаешь задачи на обучающей платформе и развиваешь свои навыки.\nЯ вижу что ты решаешь задачу про игру Алисы и Боба, где должен определить победителя.\nНо твоя ошибка заключается в том, что вместо инкремента cnt1 и cnt2 ты их декрементируешь.\nПопробуй поменять свой код следующим образом:\n```cpp\nif (s[i] == '0') {\n    cnt1++;\n} else {\n    cnt2++;\n}```",
	}

	//err = json.Unmarshal(body, &ChatGPTResponse)
	//err := json.Unmarshal([]byte("Привет! Рада видеть, что ты решаешь задачи на обучающей платформе и развиваешь свои навыки.\nЯ вижу что ты решаешь задачу про игру Алисы и Боба, где должен определить победителя.\nНо твоя ошибка заключается в том, что вместо инкремента cnt1 и cnt2 ты их декрементируешь.\nПопробуй поменять свой код следующим образом:\n```cpp\nif (s[i] == '0') {\n    cnt1++;\n} else {\n    cnt2++;\n}```"), &ChatGPTResponse)
	//if err != nil {
	//	return nil, errors.Errorf("1511 " + err.Error())
	//}

	return ChatGPTResponse, nil
}
