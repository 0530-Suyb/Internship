package main

import "fmt"

// 组合模式
// 组合模式是一种结构型设计模式，它允许你将对象组合成树形结构来表示部分-整体的层次结构。
// 组合模式使得客户端对单个对象和组合对象的使用具有一致性。
// 组合模式通常用于表示树形结构的数据，例如文件系统、组织结构等。
// 组合模式的核心思想是将对象组合成树形结构，以便于客户端以统一的方式处理单个对象和组合对象。

// 文件系统案例

type FileSystemComponent interface {
	// 显示文件或目录的信息
	ShowInfo() string
	// 组件类型
	Type() string
	// 添加子组件（仅对目录有效）
	Add(component FileSystemComponent)
	// 删除子组件（仅对目录有效）
	Delete(component FileSystemComponent)
	// 获取子组件（仅对目录有效）
	GetChild(index int) FileSystemComponent
	// 获取子组件的数量（仅对目录有效）
	GetChildCount() int
}

type File struct {
	name string
}

func (f *File) ShowInfo() string {
	return fmt.Sprintf("File: %s", f.name)
}

func (f *File) Type() string {
	return "File"
}

func (f *File) Add(component FileSystemComponent) {
	panic("Cannot add component to a file")
}

func (f *File) Delete(component FileSystemComponent) {
	panic("Cannot delete component from a file")
}

func (f *File) GetChild(index int) FileSystemComponent {
	panic("Cannot get child from a file")
}

func (f *File) GetChildCount() int {
	return 0
}

type Directory struct {
	name     string
	children []FileSystemComponent
}

func (d *Directory) ShowInfo() string {
	return fmt.Sprintf("Directory: %s", d.name)
}

func (d *Directory) Type() string {
	return "Directory"
}

func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

func (d *Directory) Delete(component FileSystemComponent) {
	for i, child := range d.children {
		if child == component {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) GetChild(index int) FileSystemComponent {
	if index < 0 || index >= len(d.children) {
		panic("Index out of bounds")
	}
	return d.children[index]
}

func (d *Directory) GetChildCount() int {
	return len(d.children)
}

func main() {
	// 创建文件和目录
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	dir1 := &Directory{name: "dir1"}
	dir2 := &Directory{name: "dir2"}

	// 添加文件到目录
	dir1.Add(file1)
	dir1.Add(file2)

	// 添加目录到根目录
	dir2.Add(dir1)

	dir2.Add(&File{name: "file3.txt"})
	dir2.Add(&File{name: "file4.txt"})

	// 显示信息
	// 使用dfs迭代遍历
	fmt.Println("Root Directory:")
	dfs(dir2, 0)

	// fmt.Println(dir2.ShowInfo())
	// for i := 0; i < dir2.GetChildCount(); i++ {
	// 	child := dir2.GetChild(i)
	// 	fmt.Println("  ", child.ShowInfo())
	// 	for j := 0; j < child.GetChildCount(); j++ {
	// 		fmt.Println("    ", child.GetChild(j).ShowInfo())
	// 	}
	// }
}

func dfs(component FileSystemComponent, depth int) {
	for i := 0; i < depth; i++ {
		if i == depth-1 {
			fmt.Print("└── ")
		} else {
			fmt.Print("    ")
		}
	}
	fmt.Println(component.ShowInfo())
	if component.Type() == "Directory" {
		for i := 0; i < component.GetChildCount(); i++ {
			child := component.GetChild(i)
			dfs(child, depth+1)
		}
	}
}
