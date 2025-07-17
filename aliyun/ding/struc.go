package ding

import "time"

type GetToken struct {
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	Errmsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
}

// 获取管理员列表
type AmdinUsersList struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  []struct {
		SysLevel int    `json:"sys_level"`
		Userid   string `json:"userid"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

// 获取用户信息
type UserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		Active        bool      `json:"active"`
		Admin         bool      `json:"admin"`
		Avatar        string    `json:"avatar"`
		Boss          bool      `json:"boss"`
		CreateTime    time.Time `json:"create_time"`
		DeptIDList    []int     `json:"dept_id_list"`
		DeptOrderList []struct {
			DeptID int   `json:"dept_id"`
			Order  int64 `json:"order"`
		} `json:"dept_order_list"`
		ExclusiveAccount bool `json:"exclusive_account"`
		HideMobile       bool `json:"hide_mobile"`
		LeaderInDept     []struct {
			DeptID int  `json:"dept_id"`
			Leader bool `json:"leader"`
		} `json:"leader_in_dept"`
		Name       string `json:"name"`
		RealAuthed bool   `json:"real_authed"`
		RoleList   []struct {
			GroupMe string `json:"group_me"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
		} `json:"role_list"`
		Senior  bool   `json:"senior"`
		Unionid string `json:"unionid"`
		Userid  string `json:"userid"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

// 表属性
type ExcelAttribute struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Visibility         string `json:"visibility"`
	LastNonEmptyRow    int    `json:"lastNonEmptyRow"`
	LastNonEmptyColumn int    `json:"lastNonEmptyColumn"`
	RowCount           int    `json:"rowCount"`
	ColumnCount        int    `json:"columnCount"`
}

// 表格内容

type ExcelData struct {
	Values           [][]string `json:"values"`
	Formulas         [][]string `json:"formulas"`
	DisplayValues    [][]string `json:"displayValues"`
	BackgroundColors [][]struct {
		Red       int    `json:"red"`
		Green     int    `json:"green"`
		Blue      int    `json:"blue"`
		HexString string `json:"hexString"`
	} `json:"backgroundColors"`
}
