package entities

type IdealType struct {

	// age
	Age *IdealTypeAge `json:"age,omitempty"`

	// annual income
	AnnualIncome *IdealTypeAnnualIncome `json:"annual_income,omitempty"`

	// 体型
	BodyBuild []string `json:"body_build"`

	// お酒を飲むか？
	Drinking []string `json:"drinking"`

	// 学歴
	Education []string `json:"education"`

	// height
	Height *IdealTypeHeight `json:"height,omitempty"`

	// 休日
	Holiday []string `json:"holiday"`

	// 出身地
	HomeState []string `json:"home_state"`

	// 仕事
	Job []string `json:"job"`

	// 居住地
	ResidenceState []string `json:"residence_state"`

	// タバコを吸うか？
	Smoking []string `json:"smoking"`
}

type IdealTypeAge struct {

	// 最低年齢
	From int64 `json:"from,omitempty"`

	// 最高年齢
	To int64 `json:"to,omitempty"`
}

type IdealTypeAnnualIncome struct {

	// 最低年収
	From int64 `json:"from,omitempty"`

	// 最高年収
	To int64 `json:"to,omitempty"`
}

type IdealTypeHeight struct {

	// 最低身長
	From int64 `json:"from,omitempty"`

	// 最高身長
	To int64 `json:"to,omitempty"`
}
