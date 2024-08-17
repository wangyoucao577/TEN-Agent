package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"strings"
)

// 用户信息
type UserInfo struct {
	ID string `json:"id,omitempty"` // calculated

	Name string `json:"name,omitempty"` // 名字

	// 基本信息
	Age              int    `json:"age,omitempty"`                // 基本信息_年龄
	Gender           string `json:"gender,omitempty"`             // 基本信息_性别
	Occupation       string `json:"occupation,omitempty"`         // 基本信息_职业
	IncomeLevel      string `json:"income_level,omitempty"`       // 基本信息_收入水平
	PlaceOfResidence string `json:"place_of_residence,omitempty"` // 基本信息_居住地

	// 家庭情况
	MaritalStatus               string `json:"marital_status,omitempty"`        // 家庭情况_婚姻状态
	NumberOfChildren            int    `json:"number_of_children,omitempty"`    // 家庭情况_子女数量
	HealthStatusOfFamilyMembers string `json:"health_status_of_family_members"` // 家庭情况_家庭成员健康状况

	HealthStatus   string `json:"health_status,omitempty"`   // 健康状况
	Lifestyle      string `json:"lifestyle,omitempty"`       // 生活方式
	MedicalHistory string `json:"medical_history,omitempty"` // 医疗历史

	// 保险
	InsuranceAwareness  string `json:"insurance_awareness,omitempty"`  // 保险_保险意识
	InsuranceExperience string `json:"insurance_experience,omitempty"` // 保险_保险经历
	InsuranceKnowledge  string `json:"insurance_knowledge,omitempty"`  // 保险_保险知识

	PersonalityType string `json:"personality_type,omitempty"` // 性格_性格类型

	// 经济状况
	EconomicLevel      string `json:"economic_level,omitempty"`      // 经济状况_经济水平
	SpendingPriorities string `json:"spending_priorities,omitempty"` // 经济状况_支出优先级

	// 日常生活
	LivingHabits string `json:"living_habits,omitempty"` // 日常生活_生活习惯
	Hobbies      string `json:"hobbies,omitempty"`       // 日常生活_兴趣爱好

	// 情绪和社交
	CurrentEmotions     string `json:"current_emotions,omitempty"`      // 情绪和社交_当前情绪
	FutureExpectations  string `json:"future_expectations,omitempty"`   // 情绪和社交_未来期望
	GossipAndComplaints string `json:"gossip_and_complaints,omitempty"` // 情绪和社交_八卦抱怨
	FamilyRelationships string `json:"family_relationships,omitempty"`  // 情绪和社交_家庭关系
}

func NewUserInfoFromCSVRecord(record []string) (UserInfo, error) {
	var user UserInfo
	var err error

	// parse fields
	user.Name = record[0]
	user.Age, err = strconv.Atoi(record[1])
	if err != nil {
		return UserInfo{}, err
	}
	user.Gender = record[2]
	user.Occupation = record[3]
	user.IncomeLevel = record[4]
	user.PlaceOfResidence = record[5]
	user.MaritalStatus = record[6]
	user.NumberOfChildren, err = strconv.Atoi(record[7])
	if err != nil {
		return UserInfo{}, err
	}
	user.HealthStatusOfFamilyMembers = record[8]
	user.HealthStatus = record[9]
	user.Lifestyle = record[10]
	user.MedicalHistory = record[11]
	user.InsuranceAwareness = record[12]
	user.InsuranceExperience = record[13]
	user.InsuranceKnowledge = record[14]
	user.PersonalityType = record[15]
	user.EconomicLevel = record[16]
	user.SpendingPriorities = record[17]
	user.LivingHabits = record[18]
	user.Hobbies = record[19]
	user.CurrentEmotions = record[20]
	user.FutureExpectations = record[21]
	user.GossipAndComplaints = record[22]
	user.FamilyRelationships = record[23]

	// generate id
	sum := sha1.Sum([]byte(strings.Join(record, ",")))
	user.ID = hex.EncodeToString(sum[:])
	return user, err
}

func (u UserInfo) CSV() []string {
	var record []string
	record = append(record, u.Name)
	record = append(record, strconv.FormatInt(int64(u.Age), 10))
	record = append(record, u.Gender)
	record = append(record, u.Occupation)
	record = append(record, u.IncomeLevel)
	record = append(record, u.PlaceOfResidence)
	record = append(record, u.MaritalStatus)
	record = append(record, strconv.FormatInt(int64(u.NumberOfChildren), 10))
	record = append(record, u.HealthStatusOfFamilyMembers)
	record = append(record, u.HealthStatus)
	record = append(record, u.Lifestyle)
	record = append(record, u.MedicalHistory)
	record = append(record, u.InsuranceAwareness)
	record = append(record, u.InsuranceExperience)
	record = append(record, u.InsuranceKnowledge)
	record = append(record, u.PersonalityType)
	record = append(record, u.EconomicLevel)
	record = append(record, u.SpendingPriorities)
	record = append(record, u.LivingHabits)
	record = append(record, u.Hobbies)
	record = append(record, u.CurrentEmotions)
	record = append(record, u.FutureExpectations)
	record = append(record, u.GossipAndComplaints)
	record = append(record, u.FamilyRelationships)
	return record
}
