# pfinal_cli_tools

**pfinal_cli_tools** 是一个集成了多种网络安全和性能测试工具的命令行界面工具包，提供了统一的交互式操作体验。

## 功能特性

- **统一的命令行界面**：集成多种常用工具，提供一致的操作方式
- **终端彩色输出**：使用彩色文本提升用户体验
- **工具自动检测**：检查依赖工具是否已安装，并提供安装指引
- **交互式配置**：通过问答方式帮助用户配置命令参数
- **pfinalclub 品牌标识**：显示特色logo

## 支持的工具

- **nmap**：网络扫描和安全审计工具
- **wrk**：HTTP性能测试工具
- **sqlmap**：SQL注入检测和利用工具

## 项目结构

```
pfinal_cli_tools/
├── cmd/
│   ├── nmap.go    # nmap命令实现
│   ├── root.go    # 根命令和主要逻辑
│   ├── sqlmap.go  # sqlmap命令实现
│   └── wrk.go     # wrk命令实现
├── internal/
│   ├── config/    # 配置相关
│   ├── ui/        # 用户界面相关
│   └── utils/     # 工具函数
│       ├── checker.go  # 工具检测功能
│       └── logo.go     # Logo显示功能
├── main.go        # 程序入口
├── go.mod         # Go模块定义
├── go.sum         # 依赖锁定
└── README.md      # 项目文档
```

## 安装方法

### 前提条件

- 安装 [Go](https://golang.org/) 1.16 或更高版本
- 安装所需的依赖工具：nmap、wrk、sqlmap（可通过本工具的检测功能引导安装）

### 编译安装

```bash
# 克隆仓库
git clone https://github.com/pfinal/pfinal_cli_tools.git
cd pfinal_cli_tools

# 编译项目
go build -o pfinal_cli_tools

# 移动到系统路径（可选）
sudo mv pfinal_cli_tools /usr/local/bin/
```

## 使用方法

### 基本用法

```bash
# 查看帮助信息
./pfinal_cli_tools --help

# 查看特定工具的帮助
./pfinal_cli_tools nmap --help
```

### 工具检测

```bash
# 检测所有依赖工具是否已安装
./pfinal_cli_tools -c
# 或
./pfinal_cli_tools --check-tools
```

### 运行特定工具

```bash
# 运行nmap工具（将进入交互式配置）
./pfinal_cli_tools nmap

# 运行wrk工具（将进入交互式配置）
./pfinal_cli_tools wrk

# 运行sqlmap工具（将进入交互式配置）
./pfinal_cli_tools sqlmap
```

## 工具检测功能

该工具会自动检测您系统中是否安装了所需的依赖工具（nmap、wrk、sqlmap），如果未安装，将提供相应的安装命令提示。

- **macOS**：使用brew安装
- **Linux**：使用apt、yum等包管理器安装
- **Windows**：提供下载链接或安装方法

## 开发说明

如果您想扩展此工具包，添加新的命令或功能，请遵循以下步骤：

1. 在`cmd/`目录下创建新的命令文件（如`newtool.go`）
2. 在`internal/utils/`目录下添加必要的工具函数
3. 在`main.go`中注册新命令

## 许可证

本项目使用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 贡献

欢迎提交问题和改进建议！如果您有兴趣为项目贡献代码，请提交Pull Request。

## 关于

本工具包由 pfinalclub 开发维护，旨在简化网络安全测试和性能评估工作流程。