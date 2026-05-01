package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// OllamaClient 定义 Ollama 客户端
type OllamaClient struct {
	Model   string
	BaseURL string
}

// OllamaRequest 定义请求 Ollama 的参数
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaResponse 定义 Ollama 返回的数据结构
type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// NewOllamaClient 创建 Ollama 客户端
func NewOllamaClient(cfg Config) *OllamaClient {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = "http://127.0.0.1:11434"
	}
	model := cfg.Model
	if model == "" {
		model = "qwen2.7:7b"
	}

	return &OllamaClient{
		Model:   model,
		BaseURL: strings.TrimRight(baseURL, "/"),
	}
}

// GenerateSQL 根据用户问题生成 SQL
func (c *OllamaClient) GenerateSQL(question string) (string, error) {
	// 构造提示词
	prompt := BuildSQLPrompt(question)

	reqBody := OllamaRequest{
		Model:  c.Model,
		Prompt: prompt,
		Stream: false,
	}

	// 转成 JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// 请求本地 Ollama
	resp, err := http.Post(
		c.BaseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", errors.New("请求 Ollama 失败，请确认 ollama serve 是否已启动")
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Ollama 返回错误：" + string(body))
	}

	var result OllamaResponse

	// 解析 JSON
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	sqlText := CleanSQL(result.Response)
	if sqlText == "" {
		return "", errors.New("Ollama 未返回 SQL")
	}

	return sqlText, nil
}
