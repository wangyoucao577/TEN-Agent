package internal

import (
	"reflect"
	"testing"
)

func TestNewUserInfoFromCSVRecord(t *testing.T) {
	record := []string{
		"王女士", "31", "女", "公关专员", "中等", "郑州市",
		"未婚", "0", "良好", "健康状况良好", "工作与社交活动频繁",
		"无重大疾病史", "中等", "有基本的健康保险", "了解保险的基本功能",
		"外向，乐观，适应力强", "稳定", "个人兴趣，社交活动，生活消费",
		"喜欢社交活动和健身", "旅行，社交，阅读", "对工作和生活感到满意",
		"希望能在职业上有所突破，并建立稳定的个人生活", "对公司内部的一些问题有时感到不满",
		"与家人关系良好",
	}

	expected := UserInfo{
		Name:                        "王女士",
		Age:                         31,
		Gender:                      "女",
		Occupation:                  "公关专员",
		IncomeLevel:                 "中等",
		PlaceOfResidence:            "郑州市",
		MaritalStatus:               "未婚",
		NumberOfChildren:            0,
		HealthStatusOfFamilyMembers: "良好",
		HealthStatus:                "健康状况良好",
		Lifestyle:                   "工作与社交活动频繁",
		MedicalHistory:              "无重大疾病史",
		InsuranceAwareness:          "中等",
		InsuranceExperience:         "有基本的健康保险",
		InsuranceKnowledge:          "了解保险的基本功能",
		PersonalityType:             "外向，乐观，适应力强",
		EconomicLevel:               "稳定",
		SpendingPriorities:          "个人兴趣，社交活动，生活消费",
		LivingHabits:                "喜欢社交活动和健身",
		Hobbies:                     "旅行，社交，阅读",
		CurrentEmotions:             "对工作和生活感到满意",
		FutureExpectations:          "希望能在职业上有所突破，并建立稳定的个人生活",
		GossipAndComplaints:         "对公司内部的一些问题有时感到不满",
		FamilyRelationships:         "与家人关系良好",
	}

	user, err := NewUserInfoFromCSVRecord(record)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Clear the ID field for comparison purposes
	user.ID = ""
	expected.ID = ""

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("got %+v, want %+v", user, expected)
	}
}

func TestUserInfo_CSV(t *testing.T) {
	user := UserInfo{
		Name:                        "王女士",
		Age:                         31,
		Gender:                      "女",
		Occupation:                  "公关专员",
		IncomeLevel:                 "中等",
		PlaceOfResidence:            "郑州市",
		MaritalStatus:               "未婚",
		NumberOfChildren:            0,
		HealthStatusOfFamilyMembers: "良好",
		HealthStatus:                "健康状况良好",
		Lifestyle:                   "工作与社交活动频繁",
		MedicalHistory:              "无重大疾病史",
		InsuranceAwareness:          "中等",
		InsuranceExperience:         "有基本的健康保险",
		InsuranceKnowledge:          "了解保险的基本功能",
		PersonalityType:             "外向，乐观，适应力强",
		EconomicLevel:               "稳定",
		SpendingPriorities:          "个人兴趣，社交活动，生活消费",
		LivingHabits:                "喜欢社交活动和健身",
		Hobbies:                     "旅行，社交，阅读",
		CurrentEmotions:             "对工作和生活感到满意",
		FutureExpectations:          "希望能在职业上有所突破，并建立稳定的个人生活",
		GossipAndComplaints:         "对公司内部的一些问题有时感到不满",
		FamilyRelationships:         "与家人关系良好",
	}

	expected := []string{
		"王女士", "31", "女", "公关专员", "中等", "郑州市",
		"未婚", "0", "良好", "健康状况良好", "工作与社交活动频繁",
		"无重大疾病史", "中等", "有基本的健康保险", "了解保险的基本功能",
		"外向，乐观，适应力强", "稳定", "个人兴趣，社交活动，生活消费",
		"喜欢社交活动和健身", "旅行，社交，阅读", "对工作和生活感到满意",
		"希望能在职业上有所突破，并建立稳定的个人生活", "对公司内部的一些问题有时感到不满",
		"与家人关系良好",
	}

	record := user.CSV()

	if !reflect.DeepEqual(record, expected) {
		t.Errorf("got %+v, want %+v", record, expected)
	}
}
