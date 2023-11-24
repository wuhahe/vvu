package vvu

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

const (
	colorWrite   = 0  // 白
	colorBlack   = 30 // 黑
	colorRed     = 31 // 红
	colorGreen   = 32 // 绿
	colorYellow  = 33 // 黄
	colorBlue    = 34 // 蓝
	colorMagenta = 35 // 紫
	colorCyan    = 36 // 青
	colorGrey    = 37 // 灰
)

type Std struct {
	projectPath string
	switchDir   bool
	switchTime  bool
	lineNum     int
	content     []any
}

var (
	projectName string
	switchDir   bool
	switchTime  bool
)

// Config 全局配置
// 配置项目名称 name 以跟踪栈信息
// dir 与 time 分别为控制是否打印输出所在行和当前时间
// Example: vvu.Config("ProjectName", true, true)
func Config(name string, dir, time bool) {
	projectName = name
	switchDir = dir
	switchTime = time
}

// PrintC 带颜色的标准输出配置
// Example: vvu.PrintC("example").Write()
func PrintC(content ...any) *Std {
	newStd := Std{
		switchDir:  switchDir,
		switchTime: switchTime,
	}
	_, fileName, line, _ := runtime.Caller(1)
	tmpFile := regexp.MustCompile(`(?m)`+projectName+`.*`).FindAllString(fileName, -1)
	if tmpFile != nil {
		newStd.projectPath = tmpFile[0]
	} else {
		newStd.switchDir = false
	}
	newStd.lineNum = line
	newStd.content = content
	return &newStd
}

// outPut 执行输出
func (o *Std) outPut(color int) {
	// 输出时间和路径
	if o.switchTime {
		nowTime := time.Now().Format("15:04:05")
		fmt.Printf("\033[0;0;%dm%s\033[0m ", colorYellow, nowTime)
	}
	if o.switchDir {
		fmt.Print(o.projectPath, ":", o.lineNum, " ")
	}
	// 内容输出
	for _, c := range o.content {
		fmt.Printf("\033[0;0;%dm%s\033[0m ", color, c)
	}
	fmt.Printf("\n")
}

// Write 白色字体
func (o *Std) Write() *Std {
	o.outPut(colorWrite)
	return o
}

// Black 黑色字体
func (o *Std) Black() *Std {
	o.outPut(colorBlack)
	return o
}

// Red 红色字体
func (o *Std) Red() *Std {
	o.outPut(colorRed)
	return o
}

// Green 绿色字体
func (o *Std) Green() *Std {
	o.outPut(colorGreen)
	return o
}

// Yellow 黄色字体
func (o *Std) Yellow() *Std {
	o.outPut(colorYellow)
	return o
}

// Blue 蓝色字体
func (o *Std) Blue() *Std {
	o.outPut(colorBlue)
	return o
}

// Magenta 紫色字体
func (o *Std) Magenta() *Std {
	o.outPut(colorMagenta)
	return o
}

// Cyan 青色字体
func (o *Std) Cyan() *Std {
	o.outPut(colorCyan)
	return o
}

// Grey 灰色字体
func (o *Std) Grey() *Std {
	o.outPut(colorGrey)
	return o
}
