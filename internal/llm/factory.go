package llm

import "fmt"

// Config 定义大模型配置
type Config struct {
	// Provider 表示模型供应商
	Provider string
	// Model 表示模型名称
	Model string
	// BaseURL 表示接口地址
	BaseURL string
	// APIKey 表示 API Key，本地 Ollama 可以为空
	APIKey string
}

// NewClient 根据配置创建不同的大模型客户端
func NewClient(cfg Config) (Client, error) {
	switch cfg.Provider {
	case "ollama":
		return NewOllamaClient(cfg), nil

	case "openai":
		// 后面接 OpenAI 时实现
		return nil, fmt.Errorf("openai 暂未实现")

	case "anthropic":
		// 后面接 Anthropic 时实现
		return nil, fmt.Errorf("anthropic 暂未实现")

	default:
		return nil, fmt.Errorf("不支持的 LLM Provider: %s", cfg.Provider)
	}
}
