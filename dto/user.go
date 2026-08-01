package dto

// User 用户
type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Avatar           string `json:"avatar"`
	Bot              bool   `json:"bot"`
	UnionOpenID      string `json:"union_openid"`       // 特殊关联应用的 openid
	UnionUserAccount string `json:"union_user_account"` // 机器人关联的用户信息，与union_openid关联的应用是同一个
	MemberRole       string `json:"member_role"`        // 权限

	/*Id           string `json:"id"`
	Username     string `json:"username"`
	Bot          bool   `json:"bot"`
	MemberOpenid string `json:"member_openid"`
	MemberRole   string `json:"member_role"`
	UnionOpenid  string `json:"union_openid"`*/
}
