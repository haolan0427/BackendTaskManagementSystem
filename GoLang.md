# Go语言学习笔记

## 目录
1. [基础语法](#基础语法)
2. [数据类型](#数据类型)
3. [流程控制](#流程控制)
4. [函数](#函数)
5. [结构体与方法](#结构体与方法)
6. [接口](#接口)
7. [并发编程](#并发编程)
8. [错误处理](#错误处理)
9. [包管理](#包管理)
10. [文件操作](#文件操作)
11. [网络编程](#网络编程)
12. [数据库操作](#数据库操作)
13. [测试](#测试)

---

## 基础语法

### Hello World
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 变量声明
```go
package main

import "fmt"

func main() {
    var name string = "Go"
    var age int = 10
    
    var (
        city string = "San Francisco"
        country string = "USA"
    )
    
    shortName := "Golang"
    fmt.Println(name, age, city, country, shortName)
}
```

### 常量
```go
package main

import "fmt"

const (
    PI = 3.14159
    MaxSize = 100
)

func main() {
    const greeting = "Hello"
    fmt.Println(PI, MaxSize, greeting)
}
```

---

## 数据类型

### 基本数据类型
```go
package main

import "fmt"

func main() {
    var (
        b bool = true
        i int = 42
        f float64 = 3.14
        s string = "Go语言"
    )
    
    fmt.Printf("bool: %t\n", b)
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float: %f\n", f)
    fmt.Printf("string: %s\n", s)
}
```

### 数组与切片
```go
package main

import "fmt"

func main() {
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3, 4, 5}
    
    slice = append(slice, 6)
    slice = append(slice, 7, 8, 9)
    
    subSlice := slice[1:4]
    
    fmt.Println("数组:", arr)
    fmt.Println("切片:", slice)
    fmt.Println("子切片:", subSlice)
    
    makeSlice := make([]int, 5, 10)
    fmt.Println("make创建的切片:", makeSlice, len(makeSlice), cap(makeSlice))
}
```

### 映射（Map）
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["apple"] = 5
    m["banana"] = 3
    
    value, exists := m["apple"]
    fmt.Println("apple的数量:", value, "存在:", exists)
    
    delete(m, "banana")
    fmt.Println("删除后的map:", m)
    
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }
}
```

### 指针
```go
package main

import "fmt"

func main() {
    x := 10
    p := &x
    
    fmt.Println("x的值:", x)
    fmt.Println("x的地址:", p)
    fmt.Println("通过指针访问x的值:", *p)
    
    *p = 20
    fmt.Println("修改后x的值:", x)
}
```

---

## 流程控制

### 条件语句
```go
package main

import "fmt"

func main() {
    score := 85
    
    if score >= 90 {
        fmt.Println("优秀")
    } else if score >= 80 {
        fmt.Println("良好")
    } else if score >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }
    
    if num := 42; num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }
}
```

### 循环
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    sum := 1
    for sum < 100 {
        sum += sum
    }
    fmt.Println("sum:", sum)
    
    names := []string{"Alice", "Bob", "Charlie"}
    for index, name := range names {
        fmt.Printf("索引 %d: %s\n", index, name)
    }
    
    for _, name := range names {
        fmt.Println(name)
    }
    
    i := 0
    for i < 3 {
        fmt.Println("while循环:", i)
        i++
    }
}
```

### Switch语句
```go
package main

import "fmt"

func main() {
    day := 3
    
    switch day {
    case 1:
        fmt.Println("星期一")
    case 2:
        fmt.Println("星期二")
    case 3:
        fmt.Println("星期三")
    default:
        fmt.Println("其他")
    }
    
    score := 85
    switch {
    case score >= 90:
        fmt.Println("A")
    case score >= 80:
        fmt.Println("B")
    case score >= 70:
        fmt.Println("C")
    default:
        fmt.Println("D")
    }
}
```

---

## 函数

### 基本函数
```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func swap(a, b string) (string, string) {
    return b, a
}

func namedReturn() (result int) {
    result = 42
    return
}

func multipleReturn() (int, string, bool) {
    return 1, "Go", true
}

func variadic(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println("加法:", add(3, 5))
    
    first, second := swap("hello", "world")
    fmt.Println(first, second)
    
    fmt.Println("命名返回值:", namedReturn())
    
    a, b, c := multipleReturn()
    fmt.Println(a, b, c)
    
    fmt.Println("可变参数:", variadic(1, 2, 3, 4, 5))
}
```

### 闭包与匿名函数
```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    
    for i := 0; i < 5; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }
    
    func(msg string) {
        fmt.Println("匿名函数:", msg)
    }("Hello")
}
```

### 递归函数
```go
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("5的阶乘:", factorial(5))
    fmt.Println("第10个斐波那契数:", fibonacci(10))
}
```

### defer语句
```go
package main

import "fmt"

func main() {
    defer fmt.Println("最后执行")
    fmt.Println("首先执行")
    
    fmt.Println("\n多个defer:")
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    
    fileOperation()
}

func fileOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("恢复:", r)
        }
    }()
    
    fmt.Println("打开文件")
    defer fmt.Println("关闭文件")
    
    panic("发生错误")
}
```

---

## 结构体与方法

### 结构体定义与使用
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
    Address
}

type Address struct {
    City    string
    Country string
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d岁) - %s, %s", p.Name, p.Age, p.City, p.Country)
}

func main() {
    p1 := Person{
        Name: "Alice",
        Age:  25,
        Address: Address{
            City:    "Beijing",
            Country: "China",
        },
    }
    
    p2 := Person{Name: "Bob", Age: 30}
    p2.City = "Shanghai"
    p2.Country = "China"
    
    fmt.Println(p1)
    fmt.Println(p2)
    
    p := &p1
    p.Name = "Alice Updated"
    fmt.Println(p1)
}
```

### 方法定义
```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

type Shape interface {
    Area() float64
}

func printArea(s Shape) {
    fmt.Printf("面积: %.2f\n", s.Area())
}

func main() {
    rect := Rectangle{Width: 3, Height: 4}
    fmt.Println("矩形面积:", rect.Area())
    
    rect.Scale(2)
    fmt.Println("缩放后矩形面积:", rect.Area())
    
    circle := Circle{Radius: 5}
    printArea(rect)
    printArea(circle)
}
```

---

## 接口

### 接口定义与实现
```go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "汪汪"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "喵喵"
}

func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "旺财"}
    cat := Cat{Name: "咪咪"}
    
    makeSound(dog)
    makeSound(cat)
    
    var speaker Speaker = dog
    fmt.Printf("类型: %T, 值: %v\n", speaker, speaker)
}
```

### 空接口与类型断言
```go
package main

import "fmt"

func printAnything(v interface{}) {
    fmt.Printf("值: %v, 类型: %T\n", v, v)
    
    if str, ok := v.(string); ok {
        fmt.Println("字符串:", str)
    } else if num, ok := v.(int); ok {
        fmt.Println("整数:", num)
    }
}

func main() {
    printAnything("Hello")
    printAnything(42)
    printAnything(3.14)
    printAnything(true)
    
    var i interface{} = "hello"
    
    s := i.(string)
    fmt.Println("类型断言:", s)
    
    s, ok := i.(string)
    fmt.Println("安全类型断言:", s, ok)
    
    n, ok := i.(int)
    fmt.Println("失败的类型断言:", n, ok)
}
```

### 接口组合
```go
package main

import "fmt"

type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}

type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(s string) {
    f.content = s
}

func process(rw ReadWriter) {
    rw.Write("Hello, Go!")
    fmt.Println("读取内容:", rw.Read())
}

func main() {
    file := &File{}
    process(file)
}
```

---

## 并发编程

### Goroutine
```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
    
    time.Sleep(time.Second)
}
```

### Channel
```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    
    x, y := <-c, <-c
    
    fmt.Println(x, y, x+y)
}
```

### 缓冲Channel
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    
    ch <- 1
    ch <- 2
    
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

### Channel方向
```go
package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    
    ping(pings, "passed message")
    pong(pings, pongs)
    
    fmt.Println(<-pongs)
}
```

### Select语句
```go
package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    
    fibonacci(c, quit)
}
```

### Mutex互斥锁
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := SafeCounter{}
    
    for i := 0; i < 1000; i++ {
        go counter.Inc()
    }
    
    time.Sleep(time.Second)
    fmt.Println("计数器值:", counter.Value())
}
```

### WaitGroup
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d 开始工作\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("所有worker完成")
}
```

---

## 错误处理

### 基本错误处理
```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Println("结果:", result)
    }
    
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Println("结果:", result)
    }
}
```

### 自定义错误
```go
package main

import "fmt"

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("错误 %d: %s", e.Code, e.Message)
}

func process(value int) error {
    if value < 0 {
        return &MyError{Code: 400, Message: "值不能为负数"}
    }
    if value > 100 {
        return &MyError{Code: 403, Message: "值不能超过100"}
    }
    return nil
}

func main() {
    err := process(-5)
    if err != nil {
        fmt.Println("处理错误:", err)
        
        if myErr, ok := err.(*MyError); ok {
            fmt.Printf("错误代码: %d\n", myErr.Code)
        }
    }
}
```

### Panic与Recover
```go
package main

import "fmt"

func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("恢复:", r)
        }
    }()
    
    panic("发生panic!")
}

func main() {
    mayPanic()
    fmt.Println("程序继续执行")
}
```

---

## 包管理

### 创建和使用包
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
    fmt.Println(math.Sqrt(16))
}
```

### 自定义包
```go
package mypackage

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
```

### Go Modules
```bash
go mod init myproject
go mod tidy
go build
```

---

## 文件操作

### 读写文件
```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    content := []byte("Hello, Go文件操作!")
    
    err := ioutil.WriteFile("test.txt", content, 0644)
    if err != nil {
        fmt.Println("写入文件错误:", err)
        return
    }
    
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("读取文件错误:", err)
        return
    }
    
    fmt.Println("文件内容:", string(data))
    
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("打开文件错误:", err)
        return
    }
    defer file.Close()
    
    stat, err := file.Stat()
    if err != nil {
        fmt.Println("获取文件信息错误:", err)
        return
    }
    
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println("读取文件错误:", err)
        return
    }
    
    fmt.Println("使用os包读取:", string(bs))
}
```

### 目录操作
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Println("创建目录错误:", err)
    }
    
    files, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("读取目录错误:", err)
        return
    }
    
    for _, file := range files {
        fmt.Println(file.Name())
    }
    
    err = os.RemoveAll("testdir")
    if err != nil {
        fmt.Println("删除目录错误:", err)
    }
}
```

---

## 网络编程

### HTTP客户端
```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.github.com")
    if err != nil {
        fmt.Println("请求错误:", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("读取响应错误:", err)
        return
    }
    
    fmt.Println("响应状态:", resp.Status)
    fmt.Println("响应内容:", string(body))
}
```

### HTTP服务器
```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
    Status  int    `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "API响应成功",
        Status:  200,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/api", apiHandler)
    
    fmt.Println("服务器启动在 :8080")
    http.ListenAndServe(":8080", nil)
}
```

---

## 数据库操作

### SQLite示例
```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT);
    DELETE FROM users;
    `
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }
    
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    
    stmt, err := tx.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    
    for i := 1; i <= 3; i++ {
        _, err = stmt.Exec(fmt.Sprintf("User%d", i), fmt.Sprintf("user%d@example.com", i))
        if err != nil {
            log.Fatal(err)
        }
    }
    tx.Commit()
    
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    for rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("用户: %+v\n", user)
    }
}
```

---

## 测试

### 单元测试
```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; 期望 %d", result, expected)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-1, -2)
    expected := -3
    
    if result != expected {
        t.Errorf("Add(-1, -2) = %d; 期望 %d", result, expected)
    }
}
```

### 基准测试
```go
package main

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(i, i+1)
    }
}
```

### 表驱动测试
```go
package main

import "testing"

func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"正数相加", 2, 3, 5},
        {"负数相加", -1, -2, -3},
        {"正负相加", 5, -3, 2},
        {"零相加", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; 期望 %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

---

## 进阶话题

### Context使用
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("worker停止:", ctx.Err())
            return
        default:
            fmt.Println("worker工作中...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    go worker(ctx)
    
    time.Sleep(2 * time.Second)
    cancel()
    
    time.Sleep(time.Second)
}
```

### JSON处理
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

func main() {
    p := Person{
        Name: "Alice",
        Age:  25,
    }
    
    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println("JSON编码错误:", err)
        return
    }
    
    fmt.Println("JSON:", string(jsonData))
    
    var p2 Person
    err = json.Unmarshal(jsonData, &p2)
    if err != nil {
        fmt.Println("JSON解码错误:", err)
        return
    }
    
    fmt.Printf("解码后: %+v\n", p2)
}
```

### 反射
```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("类型:", reflect.TypeOf(x))
    fmt.Println("值:", reflect.ValueOf(x))
    
    v := reflect.ValueOf(x)
    fmt.Println("可设置:", v.CanSet())
    
    p := reflect.ValueOf(&x)
    v = p.Elem()
    v.SetFloat(7.1)
    fmt.Println("修改后:", x)
}
```

---

## 最佳实践

1. **错误处理**: 始终检查错误，不要忽略
2. **并发**: 使用channel进行goroutine间通信
3. **接口**: 设计小而专注的接口
4. **命名**: 使用清晰、描述性的名称
5. **测试**: 为关键功能编写测试
6. **文档**: 为公开的函数和包添加文档注释

---

## 学习资源

- [Go官方文档](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go标准库](https://golang.org/pkg/)

---

*最后更新: 2026-01-22*
