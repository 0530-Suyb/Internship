package main

import "fmt"

// 命令模式
// 命令模式（Command Pattern）是一种行为设计模式，它将请求封装为对象，从而使您可以使用不同的请求、排队请求或记录请求，并支持可撤销的操作。
// 在命令模式中，您将请求封装在一个对象中，该对象包含执行请求所需的信息。
// 该对象通常实现一个接口，该接口定义了执行请求的方法。
// 这样，您可以将请求作为参数传递给其他对象，并在需要时执行请求。命令模式通常用于实现撤销/重做功能、宏命令和日志记录等场景。

// 数据库事件案例

type Command interface {
	Execute() string
	Undo() string
}

type InsertCommand struct {
	db   *Database
	data string
}

func (c *InsertCommand) Execute() string {
	c.db.Insert(c.data)
	return fmt.Sprintf("Inserted: %s", c.data)
}
func (c *InsertCommand) Undo() string {
	c.db.Delete(c.data)
	return fmt.Sprintf("Deleted: %s", c.data)
}
func NewInsertCommand(db *Database, data string) *InsertCommand {
	return &InsertCommand{db: db, data: data}
}

type DeleteCommand struct {
	db   *Database
	data string
}

func (c *DeleteCommand) Execute() string {
	c.db.Delete(c.data)
	return fmt.Sprintf("Deleted: %s", c.data)
}
func (c *DeleteCommand) Undo() string {
	c.db.Insert(c.data)
	return fmt.Sprintf("Inserted: %s", c.data)
}
func NewDeleteCommand(db *Database, data string) *DeleteCommand {
	return &DeleteCommand{db: db, data: data}
}

type Database struct {
	data []string
}

func (db *Database) Insert(data string) {
	db.data = append(db.data, data)
}

func (db *Database) Delete(data string) {
	for i, d := range db.data {
		if d == data {
			db.data = append(db.data[:i], db.data[i+1:]...)
			break
		}
	}
}
func (db *Database) GetData() []string {
	return db.data
}
func (db *Database) PrintData() {
	for _, d := range db.data {
		fmt.Println(d)
	}
}
func NewDatabase() *Database {
	return &Database{data: []string{}}
}

type Transaction struct {
	cmds []Command
}

func NewTransaction() *Transaction {
	return &Transaction{cmds: []Command{}}
}
func (t *Transaction) AddCommand(cmd Command) {
	t.cmds = append(t.cmds, cmd)
}
func (t *Transaction) Execute() {
	for _, cmd := range t.cmds {
		fmt.Println(cmd.Execute())
	}
}
func (t *Transaction) Undo() {
	for i := len(t.cmds) - 1; i >= 0; i-- {
		fmt.Println(t.cmds[i].Undo())
	}
}

func main() {
	db := NewDatabase()
	insertCmd1 := NewInsertCommand(db, "Data1")
	insertCmd2 := NewInsertCommand(db, "Data2")
	deleteCmd := NewDeleteCommand(db, "Data1")

	transaction := NewTransaction()
	transaction.AddCommand(insertCmd1)
	transaction.AddCommand(insertCmd2)
	transaction.AddCommand(deleteCmd)

	fmt.Println("Executing transaction:")
	transaction.Execute()

	fmt.Println("Current database data:")
	db.PrintData()

	fmt.Println("Undoing transaction:")
	transaction.Undo()

	fmt.Println("Current database data after undo:")
	db.PrintData()
}
