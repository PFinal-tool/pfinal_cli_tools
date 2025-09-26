package config

// ParamType 定义参数类型
const (
	ParamTypeString  = "string"
	ParamTypeBool    = "bool"
	ParamTypeSelect  = "select"
	ParamTypeInt     = "int"
)

// ParamOption 定义选项参数的选项
type ParamOption struct {
	Value string
	Label string
}

// Parameter 定义参数结构
type Parameter struct {
	Name        string
	Type        string
	Label       string
	Description string
	Default     string
	Required    bool
	Options     []ParamOption
}

// ToolTemplate 定义工具模板
type ToolTemplate struct {
	Name       string
	Parameters []Parameter
}

// GetNmapTemplate 获取nmap工具的参数模板
func GetNmapTemplate() ToolTemplate {
	return ToolTemplate{
		Name: "nmap",
		Parameters: []Parameter{
			{
				Name:        "target",
				Type:        ParamTypeString,
				Label:       "扫描目标",
				Description: "要扫描的IP地址或主机名",
				Default:     "",
				Required:    true,
			},
			{
				Name:        "scanType",
				Type:        ParamTypeSelect,
				Label:       "扫描类型",
				Description: "选择扫描类型",
				Default:     "-sS",
				Required:    true,
				Options: []ParamOption{
					{Value: "-sS", Label: "TCP SYN扫描 (默认)"},
					{Value: "-sT", Label: "TCP连接扫描"},
					{Value: "-sU", Label: "UDP扫描"},
					{Value: "-sA", Label: "ACK扫描"},
					{Value: "-sW", Label: "窗口扫描"},
				},
			},
			{
				Name:        "ports",
				Type:        ParamTypeString,
				Label:       "端口范围",
				Description: "要扫描的端口，如 22,80,443 或 1-1000",
				Default:     "",
				Required:    false,
			},
			{
				Name:        "osDetect",
				Type:        ParamTypeBool,
				Label:       "操作系统检测",
				Description: "是否进行操作系统检测",
				Default:     "false",
				Required:    false,
			},
			{
				Name:        "serviceDetect",
				Type:        ParamTypeBool,
				Label:       "服务版本检测",
				Description: "是否进行服务版本检测",
				Default:     "false",
				Required:    false,
			},
			{
				Name:        "verbose",
				Type:        ParamTypeBool,
				Label:       "详细输出",
				Description: "是否显示详细输出",
				Default:     "false",
				Required:    false,
			},
		},
	}
}

// GetWrkTemplate 获取wrk工具的参数模板
func GetWrkTemplate() ToolTemplate {
	return ToolTemplate{
		Name: "wrk",
		Parameters: []Parameter{
			{
				Name:        "url",
				Type:        ParamTypeString,
				Label:       "测试URL",
				Description: "要测试的HTTP URL",
				Default:     "",
				Required:    true,
			},
			{
				Name:        "threads",
				Type:        ParamTypeInt,
				Label:       "线程数",
				Description: "使用的线程数量",
				Default:     "2",
				Required:    true,
			},
			{
				Name:        "connections",
				Type:        ParamTypeInt,
				Label:       "连接数",
				Description: "保持的HTTP连接数量",
				Default:     "100",
				Required:    true,
			},
			{
				Name:        "duration",
				Type:        ParamTypeString,
				Label:       "测试持续时间",
				Description: "测试持续时间，如 10s, 1m, 2h",
				Default:     "10s",
				Required:    true,
			},
		},
	}
}

// GetSqlmapTemplate 获取sqlmap工具的参数模板
func GetSqlmapTemplate() ToolTemplate {
	return ToolTemplate{
		Name: "sqlmap",
		Parameters: []Parameter{
			{
				Name:        "url",
				Type:        ParamTypeString,
				Label:       "目标URL",
				Description: "包含SQL注入漏洞的URL",
				Default:     "",
				Required:    true,
			},
			{
				Name:        "method",
				Type:        ParamTypeSelect,
				Label:       "HTTP方法",
				Description: "HTTP请求方法",
				Default:     "GET",
				Required:    true,
				Options: []ParamOption{
					{Value: "GET", Label: "GET"},
					{Value: "POST", Label: "POST"},
				},
			},
			{
				Name:        "data",
				Type:        ParamTypeString,
				Label:       "POST数据",
				Description: "POST请求的数据，如 'id=1'",
				Default:     "",
				Required:    false,
			},
			{
				Name:        "level",
				Type:        ParamTypeSelect,
				Label:       "测试级别",
				Description: "测试级别(1-5)",
				Default:     "1",
				Required:    true,
				Options: []ParamOption{
					{Value: "1", Label: "级别1 (最快)"},
					{Value: "2", Label: "级别2"},
					{Value: "3", Label: "级别3"},
					{Value: "4", Label: "级别4"},
					{Value: "5", Label: "级别5 (最全面)"},
				},
			},
			{
				Name:        "risk",
				Type:        ParamTypeSelect,
				Label:       "风险级别",
				Description: "风险级别(1-3)",
				Default:     "1",
				Required:    true,
				Options: []ParamOption{
					{Value: "1", Label: "风险1 (最安全)"},
					{Value: "2", Label: "风险2"},
					{Value: "3", Label: "风险3 (可能导致数据修改)"},
				},
			},
		},
	}
}