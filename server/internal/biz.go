package internal

import (
	"fmt"
	"log/slog"
)

var taskPrompt string

func SetTaskPrompt(prompt string) {
	taskPrompt = prompt
	slog.Info("task prompt set", slog.String("prompt", taskPrompt))
}

func generatePrompt(customerInfo CustomerInfo) string {
	if customerInfo == nil {
		return taskPrompt
	}

	var prompt string
	if s, ok := customerInfo["name"]; ok {
		prompt += fmt.Sprintf("# 人物: %s\n", s)
	}
	prompt += "## 描述：\n\n"

	prompt += "### 基本信息：\n"
	if s, ok := customerInfo["age"]; ok {
		prompt += fmt.Sprintf("- 年龄：%s\n", s)
	}
	if s, ok := customerInfo["gender"]; ok {
		prompt += fmt.Sprintf("- 性别：%s\n", s)
	}
	if s, ok := customerInfo["occupation"]; ok {
		prompt += fmt.Sprintf("- 职业：%s\n", s)
	}
	if s, ok := customerInfo["income_level"]; ok {
		prompt += fmt.Sprintf("- 收入水平：%s\n", s)
	}
	if s, ok := customerInfo["place_of_residence"]; ok {
		prompt += fmt.Sprintf("- 居住地：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 家庭情况：\n"
	if s, ok := customerInfo["marital_status"]; ok {
		prompt += fmt.Sprintf("- 婚姻状态：%s\n", s)
	}
	if s, ok := customerInfo["number_of_children"]; ok {
		prompt += fmt.Sprintf("- 子女信息：%s\n", s)
	}
	if s, ok := customerInfo["health_status_of_family_members"]; ok {
		prompt += fmt.Sprintf("- 家庭成员健康状况：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 健康状况：\n"
	if s, ok := customerInfo["health_status"]; ok {
		prompt += fmt.Sprintf("- 当前健康状况：%s\n", s)
	}
	if s, ok := customerInfo["medical_history"]; ok {
		prompt += fmt.Sprintf("- 医疗历史：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 经济状况：\n"
	if s, ok := customerInfo["economic_level"]; ok {
		prompt += fmt.Sprintf("- 经济状况：%s\n", s)
	}
	if s, ok := customerInfo["spending_priorities"]; ok {
		prompt += fmt.Sprintf("- 支出优先级：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 保险意识：\n"
	if s, ok := customerInfo["insurance_awareness"]; ok {
		prompt += fmt.Sprintf("- 保险意识：%s\n", s)
	}
	if s, ok := customerInfo["insurance_experience"]; ok {
		prompt += fmt.Sprintf("- 保险经历：%s\n", s)
	}
	if s, ok := customerInfo["insurance_knowledge"]; ok {
		prompt += fmt.Sprintf("- 保险认知：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 性格：\n"
	if s, ok := customerInfo["personality_type"]; ok {
		prompt += fmt.Sprintf("- 性格类型：%s\n", s)
	}

	prompt += "### 生活方式：\n"
	if s, ok := customerInfo["living_habits"]; ok {
		prompt += fmt.Sprintf("- 生活习惯：%s\n", s)
	}
	if s, ok := customerInfo["hobbies"]; ok {
		prompt += fmt.Sprintf("- 兴趣爱好：%s\n", s)
	}
	prompt += "\n"

	prompt += "### 情绪和社交：\n"
	if s, ok := customerInfo["current_emotions"]; ok {
		prompt += fmt.Sprintf("- 当前情绪：%s\n", s)
	}
	if s, ok := customerInfo["future_expectations"]; ok {
		prompt += fmt.Sprintf("- 未来期望：%s\n", s)
	}
	if s, ok := customerInfo["gossip_and_complaints"]; ok {
		prompt += fmt.Sprintf("- 家庭关系：%s\n", s)
	}
	if s, ok := customerInfo["family_relationships"]; ok {
		prompt += fmt.Sprintf("- 社会关系：%s\n", s)
	}
	prompt += "\n"

	return prompt + taskPrompt
}
