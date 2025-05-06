package cmd

import (
	"bufio"
	"fmt"
	"github.com/scliangx/ginbuilder/builder"
	"github.com/scliangx/ginbuilder/builder/tools"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// NewCommandOptions 存储new命令的所有选项
type NewCommandOptions struct {
	projectName string
	packageName string
	path        string // 添加路径参数
}

// NewCommand 创建并配置new命令
func NewCommand() *cobra.Command {
	opts := &NewCommandOptions{}
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new Gin web project",
		Long: `Create a new Gin web project with a standard layout and common functionalities.
        The project will be created in $GOPATH/src directory.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	// 添加命令参数
	opts.AddFlags(cmd)
	return cmd
}

// AddFlags 添加命令行参数
func (o *NewCommandOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.projectName, "project", "p", "", "Gin project name")
	cmd.Flags().StringVar(&o.packageName, "package", "pkg", "Gin project package name (defaults to project name)")
	cmd.Flags().StringVarP(&o.path, "directory", "d", ".", "Directory to create the project in")
}

// Run 执行命令逻辑
func (o *NewCommandOptions) Run() error {
	if o.projectName == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	if o.packageName == "" {
		fmt.Println("[INFO]: Package name is empty, using project name as package name")
		o.packageName = o.projectName
	}

	if o.path == "" {
		gopath := tools.GetGOPath()
		o.path = filepath.Join(gopath, "src")
	}

	// 这里需要修改 builder.RenderFile 方法以支持自定义路径
	dir, err := builder.RenderFile(o.projectName, o.packageName, o.path)
	if err != nil {
		return fmt.Errorf("create project [%s] failed: %v", o.projectName, err)
	}
	err = os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("cd project directory failed: %v", err)
	}
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir

	// 获取命令的管道
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {

		return fmt.Errorf("无法获取标准输出管道: %v\n", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {

		return fmt.Errorf("无法获取错误输出管道: %v\n", err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {

		return fmt.Errorf("启动命令失败: %v\n", err)
	}

	fmt.Println("执行 go mod tidy...")

	// 实时读取并打印输出
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("命令执行失败: %v\n", err)
	}

	fmt.Println("go mod tidy 执行完成")
	fmt.Println("[SUCCESS]: Project created successfully")
	return nil
}

// 在init中注册命令
func init() {
	RegisterCommand(NewCommand())
}
