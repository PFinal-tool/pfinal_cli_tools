package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pfinal/pfinal_cli_tools/internal/config"
)

// Model 定义交互式配置的模型
type Model struct {
	template  config.ToolTemplate
	answers   map[string]string
	cursor    int
	done      bool
	errorMsg  string
}

// StartInteractiveConfig 启动交互式配置
func StartInteractiveConfig(template config.ToolTemplate) map[string]string {
	initialModel := Model{
		template: template,
		answers:  make(map[string]string),
		cursor:   0,
		done:     false,
		errorMsg: "",
	}

	// 初始化默认值
	for _, param := range template.Parameters {
		initialModel.answers[param.Name] = param.Default
	}

	p := tea.NewProgram(initialModel)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("启动交互式配置失败: %v\n", err)
		return make(map[string]string)
	}

	return finalModel.(Model).answers
}

// Init 初始化模型
func (m Model) Init() tea.Cmd {
	return nil
}

// Update 更新模型状态
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 清除错误消息
	m.errorMsg = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// 处理键盘输入
		switch msg.String() {
		case "ctrl+c", "q":
			// 退出程序
			m.done = true
			return m, tea.Quit

		case "enter":
			// 如果在最后一个问题，完成配置
			if m.cursor >= len(m.template.Parameters) {
				m.done = true
				return m, tea.Quit
			}

		case "up", "k":
			// 向上移动光标
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// 向下移动光标
			if m.cursor < len(m.template.Parameters)+1 {
				m.cursor++
			}

		case "right", "l":
			// 对于布尔类型和选择类型，向右选择
			if m.cursor < len(m.template.Parameters) {
				param := m.template.Parameters[m.cursor]
				if param.Type == config.ParamTypeBool {
					if m.answers[param.Name] == "true" {
						m.answers[param.Name] = "false"
					} else {
						m.answers[param.Name] = "true"
					}
				} else if param.Type == config.ParamTypeSelect {
					// 查找当前值的索引
					currentIndex := -1
					for i, option := range param.Options {
						if option.Value == m.answers[param.Name] {
							currentIndex = i
							break
						}
					}
					// 移动到下一个选项
					if currentIndex >= 0 {
						currentIndex = (currentIndex + 1) % len(param.Options)
						m.answers[param.Name] = param.Options[currentIndex].Value
					}
				}
			}

		case "left", "h":
			// 对于布尔类型和选择类型，向左选择
			if m.cursor < len(m.template.Parameters) {
				param := m.template.Parameters[m.cursor]
				if param.Type == config.ParamTypeBool {
					if m.answers[param.Name] == "true" {
						m.answers[param.Name] = "false"
					} else {
						m.answers[param.Name] = "true"
					}
				} else if param.Type == config.ParamTypeSelect {
					// 查找当前值的索引
					currentIndex := -1
					for i, option := range param.Options {
						if option.Value == m.answers[param.Name] {
							currentIndex = i
							break
						}
					}
					// 移动到上一个选项
					if currentIndex >= 0 {
						currentIndex = (currentIndex - 1 + len(param.Options)) % len(param.Options)
						m.answers[param.Name] = param.Options[currentIndex].Value
					}
				}
			}

		default:
			// 处理文本输入
			if m.cursor < len(m.template.Parameters) {
				param := m.template.Parameters[m.cursor]
				if param.Type == config.ParamTypeString || param.Type == config.ParamTypeInt {
					// 处理退格键
					if msg.String() == "backspace" {
						if len(m.answers[param.Name]) > 0 {
							m.answers[param.Name] = m.answers[param.Name][:len(m.answers[param.Name])-1]
						}
					} else {
						// 添加新字符
						m.answers[param.Name] += string(msg.Runes)
					}
				}
			}
		}

	case tea.WindowSizeMsg:
		// 处理窗口大小变化
		// 这里可以根据需要实现
	}

	return m, nil
}

// View 渲染视图
func (m Model) View() string {
	var s string

	// 显示标题
	s += "\n" + m.template.Name + " 工具参数配置\n"
	s += "------------------------\n\n"

	// 显示错误消息
	if m.errorMsg != "" {
		s += "错误: " + m.errorMsg + "\n\n"
	}

	// 显示参数列表
	for i, param := range m.template.Parameters {
		// 显示光标
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// 显示参数标签
		s += fmt.Sprintf("%s %s", cursor, param.Label)
		if param.Required {
			s += " *"
		}
		s += "\n"

		// 显示参数描述
		s += "  " + param.Description + "\n"

		// 根据参数类型显示不同的输入方式
		answer := m.answers[param.Name]
		s += "  当前值: "

		switch param.Type {
		case config.ParamTypeString:
				s += fmt.Sprintf("%s_", answer)

		case config.ParamTypeInt:
				s += fmt.Sprintf("%s_", answer)

		case config.ParamTypeBool:
				if answer == "true" {
					s += "[✓] 是   [ ] 否"
				} else {
					s += "[ ] 是   [✓] 否"
				}

		case config.ParamTypeSelect:
				// 查找当前选项的标签
				currentLabel := answer
				for _, option := range param.Options {
					if option.Value == answer {
						currentLabel = option.Label
						break
					}
				}
				s += currentLabel
		}

		s += "\n\n"
	}

	// 显示完成按钮
	cursor := " "
	if m.cursor == len(m.template.Parameters) {
		cursor = ">"
	}
	s += fmt.Sprintf("%s 完成配置\n", cursor)

	// 显示帮助信息
	s += "\n使用方向键移动，Enter确认，Ctrl+C退出\n"

	return s
}