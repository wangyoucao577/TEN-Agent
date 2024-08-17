package internal

type CustomerInfo map[string]string

func (u CustomerInfo) Empty() bool {
	return len(u) == 0
}

func GetCustomerProperties() map[string]string {
	return customerInfoFieldsMappingEN2CN
}

var customerInfoFieldsMappingEN2CN = map[string]string{
	"id": "id",

	"name": "名字",

	// 基本信息
	"age":                "基本信息_年龄",
	"gender":             "基本信息_性别",
	"occupation":         "基本信息_职业",
	"income_level":       "基本信息_收入水平",
	"place_of_residence": "基本信息_居住地",

	// 家庭情况
	"marital_status":                  "家庭情况_婚姻状态",
	"number_of_children":              "家庭情况_子女数量",
	"health_status_of_family_members": "家庭情况_家庭成员健康状况",

	"health_status":   "健康状况",
	"lifestyle":       "生活方式",
	"medical_history": "医疗历史",

	// 保险
	"insurance_awareness":  "保险_保险意识",
	"insurance_experience": "保险_保险经历",
	"insurance_knowledge":  "保险_保险知识",

	"personality_type": "性格_性格类型",

	// 经济状况
	"economic_level":      "经济状况_经济水平",
	"spending_priorities": "经济状况_支出优先级",

	// 日常生活
	"living_habits": "日常生活_生活习惯",
	"hobbies":       "日常生活_兴趣爱好",

	// 情绪和社交
	"current_emotions":      "情绪和社交_当前情绪",
	"future_expectations":   "情绪和社交_未来期望",
	"gossip_and_complaints": "情绪和社交_八卦抱怨",
	"family_relationships":  "情绪和社交_家庭关系",
}
