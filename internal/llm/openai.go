package llm

import (
	"context"
	"errors"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// GenerateSQL 根据用户问题生成 SQL
func GenerateSQL(question string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// 创建 OpenAI 客户端
	client := openai.NewClient(apiKey)

	systemPrompt := `
你是一个 SQL 生成助手。
你只能根据用户问题生成 MySQL 查询语句。
要求：
1. 只能生成 SELECT 查询
2. 不能生成 INSERT、UPDATE、DELETE、DROP、ALTER、TRUNCATE
3. 必须带 LIMIT，默认 LIMIT 10
4. 只返回 SQL，不要解释
5. 表结构如下：

users 表：
- id: 用户ID
- name: 用户名
- phone: 手机号
- created_at: 注册时间

orders 表：
- id: 订单ID
- user_id: 用户ID
- amount: 订单金额
- status: 订单状态，0待支付，1支付成功，2支付失败，3已退款
- created_at: 下单时间
`

	// 调用模型
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4.1-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	// 判断模型是否返回内容
	if len(resp.Choices) == 0 {
		return "", errors.New("AI 未返回结果")
	}
	return resp.Choices[0].Message.Content, nil
}
