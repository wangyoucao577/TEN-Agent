package internal

import (
	"bytes"
	"log/slog"
	"text/template"
)

type PromptTemplate struct {
	TaskPromptTemplate     string
	CustomerPromptTemplate string

	taskTmpl     *template.Template
	customerTmpl *template.Template
}

func (p *PromptTemplate) Init() error {
	var err error
	p.taskTmpl, err = template.New("task_prompt").Parse(p.TaskPromptTemplate)
	if err != nil {
		slog.Error("parse taskPromptTemplate failed", slog.Any("error", err))
		return err
	}
	slog.Info("task prompt template initalized", slog.String("taskPromptTemplate", p.TaskPromptTemplate))

	p.customerTmpl, err = template.New("customer_prompt").Parse(p.CustomerPromptTemplate)
	if err != nil {
		slog.Error("parse customerPromptTemplate failed", slog.Any("error", err))
		return err
	}
	slog.Info("customer prompt template initalized", slog.String("customerPromptTemplate", p.CustomerPromptTemplate))

	return nil
}

func executePromptTemplate(tmpl *template.Template, customerInfo CustomerInfo) string {

	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, customerInfo)
	if err != nil {
		slog.Error("execute template failed", slog.Any("error", err))
		return ""
	}

	return buf.String()
}

func (p PromptTemplate) GeneratePrompt(customerInfo CustomerInfo) string {
	customerPrompt := executePromptTemplate(p.customerTmpl, customerInfo)
	taskPrompt := executePromptTemplate(p.taskTmpl, customerInfo)

	return customerPrompt + taskPrompt
}
