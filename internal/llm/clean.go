package llm

import "strings"

// CleanSQL 清理模型返回的 SQL 文本
func CleanSQL(text string) string {
	sqlText := strings.TrimSpace(text)

	// 清理 Markdown SQL 代码块
	sqlText = strings.TrimPrefix(sqlText, "```sql")
	sqlText = strings.TrimPrefix(sqlText, "```")
	sqlText = strings.TrimSuffix(sqlText, "```")

	return strings.TrimSpace(sqlText)
}
