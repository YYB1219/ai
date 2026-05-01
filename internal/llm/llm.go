package llm

// Client 定义统一的大模型接口
// 后面无论接 Ollama、OpenAI、Anthropic，都实现这个接口
type Client interface {
	// GenerateSQL 根据用户问题生成 SQL
	GenerateSQL(question string) (string, error)
}
