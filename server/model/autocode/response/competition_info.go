package response

type ComSelectList struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	CName    string `gorm:"column:c_name;type:varchar(255)" json:"c_name"` // 竞赛名称
	CId      int64  `gorm:"column:c_id;type:int(11);NOT NULL" json:"c_id"` // 竞赛编号
	Version  string `gorm:"column:version;type:varchar(12)" json:"version"`
	FVersion string `json:"f_version"`
	SId      int64  `json:"s_id"`                                      // 竞赛编号
	Level    string `gorm:"column:level;type:varchar(4)" json:"level"` // 竞赛级别
}

type ComSelectListTree struct {
	Id       int64  `json:"id" primaryKey:"yes"`
	CName    string `gorm:"column:c_name;type:varchar(255)" json:"label"` // 竞赛名称
	Children []struct {
		CId      int64  `gorm:"column:c_id;type:int(11);NOT NULL" json:"-"  fid:"Id"` // 竞赛编号
		Version  string `gorm:"column:version;type:varchar(12)" json:"label"  primaryKey:"yes"`
		Children []struct {
			FVersion string ` json:"-" fid:"Version"`
			SId      int64  `json:"value" primaryKey:"yes"`
			Level    string `gorm:"column:level;type:varchar(4)" json:"label"  ` // 竞赛级别
		} `json:"children"`
	} `json:"children"`
}
