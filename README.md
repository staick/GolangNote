# GolangNote

[TOC]

## 相关网站

Go官网：https://go.dev/

Go中文网：https://studygolang.com/

## 0x00 VSCode开发环境配置

1. 安装插件

2. 设置代理

   Go1.13版本之后，推荐使用"https://goproxy.cn"代理

	```shell
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	```

3. 安装工具包

	在VSCode中使用`Ctrl+Shift+P`

## 0x01 基础语法

1. 编译并运行程序

   ```shell
   go run test.go
   ```

2. 构建程序

   ```shell
   go build test.go
   ```

### 1. 类型

1. go是强类型的语言，在声明变量的时候需要指定变量的类型。

2. 对于bool值，go不允许使用1和0代表true和false

#### 1.1 类型

   - 整数
     - int32
     - int64
   - 浮点数
     - float32
     - float64，常用
   - 字符串
     - string，不能对字符串执行数学运算
   - 数组

#### 1.2 检查变量类型

 使用`reflect.TypeOf(变量)`来显示变量的类型

```go
var i int = 3
fmt.Println(reflect.TypeOf(i))
```

#### 1.3 类型转换

- go不支持隐式类型转换

- 可以用类似C语言的强制类型转换

  ```go
  var f float64 = 1.1
  fmt.Println(int64(f))
  ```
  
  实测上述方法不支持字符串和布尔类型的转换
- strconv包提供了一系列类型转换的方法

  ```go
	var s string = "true"
	//将s转化为bool类型
	b, err := strconv.ParseBool(s)
	```
	


### 2. 变量

#### 2.1 声明变量

1. 变量声明的方式，在函数内使用简短变量声明，再函数外使用省略类型

   ```go
   var s string = "Hello World"
   
   //省略变量类型
   var s = "Hello World"
   
   var t string
   t = "Hello World"
   
   //简短变量声明，常用方式，但是只能再函数内部使用
   u := "Hello World"
   ```

2. 快捷生成类型不同的变量

   ```go
   var (
       s string = "foo"
       i int = 4
   )
   ```

3. 变量的默认值（零值）

   ```go
   var i int		//0
   var f float64	//0
   var b bool		//false
   var s string	//""
   ```

#### 2.2 指针

再变量名前加上&可以获取变量再计算机内存中的地址

   ```go
   s := "Hello World"
   fmt.Println(&s)
   ```

#### 2.3 声明常量

常量是不允许修改的

```go
const s string = "Hello World"
```

### 3. 函数

#### 3.1 返回单个值

```go
func function() bool {
    
}
```

#### 3.2 返回多个值

```go
func function() (int, string) {
    
}

//调用时，可以使用两个变量接收
a, b := function()
```

#### 3.3 不定参函数

```go
//函数中，num是一个包含所有参数的切片
func function(num...int) int {
    
}
```

#### 3.4 具名返回值

```go
//使用具名返回值时无需显示的返回相应的变量
func function() (x, y string) {
    //只需return即可
}
```

#### 3.5 递归函数

```go
func selfLoop(n int) {
	if n <= 0 {
		fmt.Println("循环结束，退出！")
		return
	}
	fmt.Printf("循环第%d\n", 11-n)
	n -= 1
	selfLoop(n)
}
```

#### 3.6 函数作为值传递

Go将函数是为一种类型，可以将函数赋值给变量，通过变量来调用它

```go
fn := func() {
    
}
fn()
```

函数作为传入参数

```go
//参数中的string为返回值类型
func function(f func() string) string {
    
}
```

### 4. 流程控制

#### 4.1 选择结构

1. if-else if-else

   ```go
   if 判断条件 {
       
   } else if 判断条件 {
       
   } else {
       
   }
   ```

2. switch-case

   ```go
   switch 变量 {
   	case 判断:
       	语句
   	case 判断:
       	语句
       default:
       	语句
   } 
   ```

#### 4.2 循环结构

1. 死循环

   ```go
   for{
       
   }
   ```

2. 常用循环

   ```go
   for i:=0;i<10;i++ {
       
   }
   ```

3. for-range循环

   ```go
   nums := []int{1, 2, 3, 4}
   for i, n := range nums {
       
   }
   ```

#### 4.3 defer语句

1. defer能够在函数返回前执行另一个函数，常用于执行清理操作或确保操作

	```go
	func main() {
    	defer fmt.Println("2")
    	fmt.Println("1")
	}
	```

2. 对于多个defer标记的语句，从后往前执行

### 5. 数据结构

#### 5.1 数组(Array)

声明数组并给它赋值

```go
var cheese [2]string
cheese[0] = "1"
cheese[1] = "2"
```

#### 5.2 切片(Slice)

切片相对于数组来说可以添加删除元素，还可以复制切片中的元素，感觉类似于动态数组。

1. 定义切片

	```go
	var cheeses = make([]string, 2)
	```

2. 添加元素

   ```go
   //添加一个元素
   cheeses = append(cheeses, "Camembert")
   
   //添加多个元素
   cheeses = append(cheeses, "Camembert", "Reblochon", "Picodon")
   ```

3. 删除元素

   ```go
   //切片没有删除的方法，但是可以使用append()配合切片实现删除
   //注意调用的其实是如slice = append(slice, anotherSlice...)所示的函数，必须加最后的三个.
   cheeses = append(cheeses[:2], cheeses[2+1:]...)
   ```

4. 复制元素

   ```go
   //使用copy()函数，第一个参数要粘贴的切片，第二个参数是要复制的切片
   copy(dest, source)
   ```

#### 5.3 集合(Map)

1. 定义集合

   ```go
   var players = make(map[string]int)
   ```

2. 添加/修改元素

   ```go
   players["cook"] = 30
   ```

3. 删除元素

   ```go
   del(players,"cook")
   ```

#### 5.4 结构体(Structure)

1. 创建结构体

   ```go
   //声明结构体
   type Movie struct {
       Name 	string
       Rating	float64 
   }
   
   //创建方式一：使用类型的方式
   var m Movie
   m.Name = "Metropolis"
   m.Rating = 0.99
   
   //创建方式二：使用new关键字
   m := new(Movie)
   m.Name = "Metropolis"
   m.Rating = 0.99
   
   //创建方式三：使用简短变量赋值（常用方式）
   m := Movie{
       Name: "Metropolis",
       Rating: 0.99,
   }
   ```

2. 嵌套结构体

   ```go
   type Superhero struct {
       Name	string
       Age		int
       Address	Address
   }
   
   type Address struct {
       Number	int
       Street	string
       City	string
   }
   ```

3. 初始化结构体

   ```go
   //结构体被创建后，结构体内的属性有默认值，默认值与变量的默认值相同，其他类型的默认值为nil
   //所以我们最好对结构体进行初始化，最常用的方法就是使用构造函数
   type Alarm struct {
       Time	string
       Sound	string
   }
   
   func NewAlarm(time string) Alarm {
       a := Alarm{
           Time:	time,
           Sound:	"Klaxon",
       }
       return a
   }
   ```

4. 比较结构体

   ```go
   //对于相同类型的结构体可以使用相等或不等运算符来进行比较
   //对于不同类型的运算符不能进行比较
   //可以使用reflect.TypeOf()查看结构体的类型
   ```

5. 公有和私有

   对于go语言来说，公有和私有主要的区别在于表示符首字符是否是大写的

   - 标识符首字母大写：公有
   - 标识符首字母小写：私有

## 0x02 面向对象

## 0x03 错误处理

## 0x04 多线程

## 0x05 Socket

### 1.Server Demo(单次连接)

1. 确定监听地址

  ```go
  ip := "127.0.0.1"
  port := 4321
  address := fmt.Sprintf("%s:%d", ip, port)
  ```

2. 创建监听

  ```go
  //func Listen(network, address string) (Listener, error)
  listener, err := net.Listen("tcp", address)
  if err != nil {
  	fmt.Println("net.Listen err:", err)
  	return
  }
  ```

3. 接受连接

  ```go
  //Accept() (Conn, error)
  conn, err := listener.Accept()
  if err != nil {
  	fmt.Println("listener.Accept err:", err)
  	return
  }
  ```

4. 创建一个容器，用于接收读取到的数据

	```go
	buf := make([]byte, 1024)
	```

5. 接收读取到的数据

  ```go
  //Read(b []byte) (n int, err error)
  cnt, err := conn.Read(buf)
  if err != nil {
  	fmt.Println("conn.Read err:", err)
  	return
  }
  ```

6. 对数据进行处理

  ```go
  //func ToUpper(s string) string
  upperData := strings.ToUpper(string(buf[0:cnt]))
  
  //Write(b []byte) (n int, err error)
  cnt, err = conn.Write([]byte(upperData))
  if err != nil {
  	fmt.Println("conn.Write err:", err)
  	return
  }
  ```

8. 关闭连接

   ```go
   conn.Close()
   ```
   

### 2.Client Demo（单次连接）

1. 建立连接

   ```go
   ip := "127.0.0.1"
   port := 4321
   address := fmt.Sprintf("%s:%d", ip, port)
   
   //func Dial(network, address string) (Conn, error)
   //ip地址默认为127.0.0.1
   //net.Dial("tcp",":4321")
   conn, err := net.Dial("tcp", address)
   if err != nil {
   	fmt.Println("net.Dial err:", err)
   	return
   }
   fmt.Println("client与server连接建立成功！")
   ```

2. 向Server发送数据

   ```go
   sendData := []byte("helloworld")
   //Write(b []byte) (n int, err error)
   	cnt, err := conn.Write(sendData)
   	if err != nil {
   		fmt.Println("conn.Write err:", err)
   		return
   	}
   fmt.Println("Client ====> Server cnt:", cnt, ", data:", string(sendData))
   ```

3. 接收服务器返回的数据

   ```go
   buf := make([]byte, 1024)
   //Read(b []byte) (n int, err error)
   cnt, err = conn.Read(buf)
   if err != nil {
   	fmt.Println("conn.Read err:", err)
   	return
   }
   fmt.Println("Client <==== Server cnt:", cnt, ", data:", string(buf[0:cnt]))
   ```

4. 关闭连接

   ```go
   conn.Close()
   ```

### 3.Server Demo(多次连接)

1. 确定监听地址

   ```go
   ip := "127.0.0.1"
   port := 8848
   address := fmt.Sprintf("%s:%d", ip, port)
   ```

2. 创建监听

   ```go
   //func Listen(network, address string) (Listener, error)
   listener, err := net.Listen("tcp", address)
   if err != nil {
   	fmt.Println("net.Listen err:", err)
   	return
   }
   ```

3. 主go程

   ```go
   for {
   		fmt.Println("监听中...")
   
   		//3.接受连接
   		//Accept() (Conn, error)
   		conn, err := listener.Accept()
   		if err != nil {
   			fmt.Println("listener.Accept err:", err)
   			return
   		}
   
   		fmt.Println("连接建立成功！")
   		//go程
   		go handleConnection(conn)
   	}
   ```

4. go程函数

   ```go
   func handleConnection(conn net.Conn) {
   	for {
   		//4.创建一个容器，用于接收读取到的数据
   		buf := make([]byte, 1024)
   
   		//5.接收读取到的数据
   		//Read(b []byte) (n int, err error)
   		cnt, err := conn.Read(buf)
   		if err != nil {
   			fmt.Println("conn.Read err:", err)
   			return
   		}
   		fmt.Println("Client ====> Server，长度：", cnt, "，数据：", string(buf[0:cnt]))
   		//6.对数据进行处理
   		//func ToUpper(s string) string
   		upperData := strings.ToUpper(string(buf[0:cnt]))
   
   		//Write(b []byte) (n int, err error)
   		cnt, err = conn.Write([]byte(upperData))
   		if err != nil {
   			fmt.Println("conn.Write err:", err)
   			return
   		}
   		fmt.Println("Client <==== Server，长度：", cnt, "，数据：", upperData)
   	}
   }
   ```

   

### 4.Client Demo(多次连接) 

1. 建立连接

   ```go
   ip := "127.0.0.1"
   port := 4321
   address := fmt.Sprintf("%s:%d", ip, port)
   
   //func Dial(network, address string) (Conn, error)
   //ip地址默认为127.0.0.1
   //net.Dial("tcp",":4321")
   conn, err := net.Dial("tcp", address)
   if err != nil {
   	fmt.Println("net.Dial err:", err)
   	return
   }
   fmt.Println("client与server连接建立成功！")
   ```

   将下面两部分包进for循环里面即可，后面添加延时，防止数据交互过快

2. 向Server发送数据

   ```go
   sendData := []byte("helloworld")
   //Write(b []byte) (n int, err error)
   	cnt, err := conn.Write(sendData)
   	if err != nil {
   		fmt.Println("conn.Write err:", err)
   		return
   	}
   fmt.Println("Client ====> Server cnt:", cnt, ", data:", string(sendData))
   ```

3. 接收服务器返回的数据

   ```go
   buf := make([]byte, 1024)
   //Read(b []byte) (n int, err error)
   cnt, err = conn.Read(buf)
   if err != nil {
   	fmt.Println("conn.Read err:", err)
   	return
   }
   fmt.Println("Client <==== Server cnt:", cnt, ", data:", string(buf[0:cnt]))
   ```

## 0x06 http

### 1.请求报文格式

1. 请求行

   1. 格式：方法 URL 协议版本号
   2. 实例：POST /chapter17/user HTTP/1.1
   3. 请求方法：
      1. GET：获取数据
      2. POST：上传数据
      3. PUT：修改数据
      4. DELETE：删除数据

2. 请求头

   1. 格式：key: value
   2. 可以有很多键值对
   3. 常见重要头：
      1. Accept：接收数据的格式
      2. User-Agent：描述用户浏览器的信息
      3. Connection：Keep-Alive(长链接)，Close(短链接)
      4. Accept-Encoding：可以接收的编码
      5. Cookie：由服务器设置的key value数据，客户端下次请求的时候可以携带
      6. Content-Type：
         1. appliction/-form(表示上传的数据是表单格式)
         2. application/json(表示body的数据格式是json格式)
      7. 用户可以自定义的
         1. name: Duke
         2. age: 18

3. 空行

   告诉服务器请求头结束了，用于分割

4. 请求体(可选的)

   1. 一般在POST方法时，会提供BODY
   2. 不建议在GET的时候使用
   3. 上传两种数据格式：
      1. 表单
      2. json数据格式

### 2.响应消息格式

1. 状态行

   1. 协议格式： 协议版本号 状态码 状态描述
   2. 实例：HTTP/1.1 200 OK
   3. 实例2：HTTP/1.1 404 not found
   4. 常用状态码：
      - 1xx：客户端可以继续发送请求（一般感知不到）
      - 2xx：正常访问，200
      - 3xx：重定向
      - 4xx
        - 401：未授权 not authorized
        - 404：Not found
      - 5xx
        - 501：Internal Error（服务器内部错误）

2. 响应头

   1. Content-Type: application/json
   2. Server: Apache
   3. Data: Mon, 12 Sep
   4. ...

3. 空行

   用于分割，表示下面没有响应头了

4. 响应包体

   通常返回json数据

### 3.http client

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//1.设置请求URL
	client := http.Client{}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("clent.Get err:", err)
		return
	}

	//2.获取请求头的相关信息
	//使用beego,gin等web框架，可以使用更加简单的方式获取下面的信息
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("header : ", resp.Header)

	fmt.Println("content-type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)

    
	body := resp.Body
	fmt.Println("body 111:", body)
    //body是接口类型，需要使用下面的方法转换成字节类型
	//func ReadAll(r io.Reader) ([]byte, error)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return
	}

	fmt.Println("body string:", string(readBodyStr))
}
```

### 4.http server

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//1.注册路由 router
	//xxxx/user ===> func1
	//xxxx/name ===> func2
	//xxxx/id ===> func3

	//https://127.0.0.1:8080/user，func是回调函数，用于路由的响应，这个回调函数原型是固定的
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request：===> 包含客户端发来的数据
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		//这里是具体处理业务逻辑

		//write：===> 通过writer将数据返回给客户端
		//func WriteString(w Writer, s string) (n int, err error)
		_, _ = io.WriteString(writer, "这是/user请求返回的数据！")
	})

	//https://127.0.0.1:8080/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/name请求返回的数据！")
	})

	//https://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("用户请求详情：")
		fmt.Println("request:", request)

		_, _ = io.WriteString(writer, "这是/id请求返回的数据！")
	})

	fmt.Println("Http Server start ...")
	//func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}

```

## 0x07 网络聊天室

实现一个网络聊天室

功能分析：

1. 上线下线
2. 聊天，其他人、自己都可以看到聊天消息
3. 查询当前聊天室的用户名who
4. 修改自己名字rename|Duke
5. 超时踢出

技术点分析：

1. socket tcp编程
2. map结构
   1. 存储所有的用户
   2. map遍历
   3. map删除
3. go程、channel
4. select（超时退出，主动退出）
5. timer定时器

### 1.tcp socket，建立多个连接

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("服务器启动成功...")

	for {
		fmt.Println("服务器监听中...")

		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		fmt.Println("连接建立成功！")

		//业务处理
		go handler(conn)
	}
}

//业务处理函数
func handler(conn net.Conn) {
	for {
		fmt.Println("业务处理开始...")

		buf := make([]byte, 1024)

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("接收客户端发送的数据：", string(buf[:cnt-1]),"，数据长度：", cnt-1)
	}
}
```

### 2.聊天室基本功能实现

### 3.who命令实现

### 4.rename命令实现

### 5.quit命令实现

### 6.map上锁

map不允许同时读写，如果有不同go程同时操作map，需要对map上锁，不上锁同时对一个map读和写会发生错误

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var idnames = make(map[int]string)
var lock sync.RWMutex

func main() {
	go func() {
		for {
			fmt.Println("111111")
			lock.Lock()
			fmt.Println("222222")
			idnames[0] = "duke"
			fmt.Println("333333")
			lock.Unlock()
		}
	}()

	go func() {
		for {
			fmt.Println("aaaaaa")
			lock.Lock()
			fmt.Println("bbbbbb")
			name := idnames[0]
			fmt.Println("name:", name)
			lock.Unlock()
		}
	}()

	for {
		fmt.Println("OVER")
		time.Sleep(1 * time.Second)
	}
}
```

## 0x08 微服务

### 1.特性

- 优点：
  1. 职责单一
  2. 轻量级通信
  3. 独立性
  4. 迭代开发
- 缺点：
  1. 运维成本高
  2. 分布式复杂度
  3. 接口成本高
  4. 重复性劳动
  5. 业务分离困难

### 2.RPC协议

IPC：进程间通信

RPC：Remote Procedure Call Protocol

- 理解：**像调用本地函数一样，去调用远程函数**
  - 通过rpc协议，传递：函数名、函数参数。达到本地，调用远端函数，得返回值到本地得目标
- 微服务使用RPC：
  - 每个服务都被封装成进程，彼此独立
  - 进程和进程之间，可以使用不同的语言实现

#### 2.1 RPC入门使用

##### 服务端

1. 注册rpc服务对象，给对象绑定方法

   1. 定义类

   2. 绑定类方法

   ```go
   rpc.RegisterName("服务名", 回调对象)
   ```

2. 创建监听器

   ```go
   listener, err := net.Listen()
   ```

3. 建立连接

   ```go
   conn, err := listener.Accept()
   ```

4. 将连接绑定rpc服务

   ```go
   rpc.ServeConn(conn)
   ```

##### 客户端

1. 用rpc连接服务器（rpc.Dial()）

   ```go
   conn, err := rpc.Dial()
   ```

2. 调用远程函数

   ```go
   conn.Call("服务名.方法名", 传入参数, 传出参数)
   ```

#### 2.2 RPC相关函数

1. 注册rpc服务

   ```go
   func (server *Server) RegisterName(name string, rcvr interface{}) error
   /*
   name:服务名，字符串类型
   rcvr:对应的rpc对象。该对象绑定方法要满足如下条件：
   	1.方法必须是导出的，包外可见（首字母大写）
   	2.方法必须有两个参数，都是导出类型、内建类型
   	3.方法的第二个参数必须是“指针”（传出参数）
   	4.方法只有一个error借口类型的返回值
   举例:func (this *World) HelloWorld (name string, resp *string) error {
   }
   	rpc.RegisterName("服务名", new(World))
   */	
   ```

2. 绑定rpc服务

   ```go
   func (server *Server) ServeConn(conn io.ReadWriteCloser)
   /*
   conn:成功建立好的socket
   */
   ```

3. 调用远程函数：

   ```go
   func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
   /*
   serviceMethod:"服务名.方法名"
   args:传入参数。方法需要的数据
   reply:传出参数。定义var 变量, &变量名 完成传参
   */
   ```

#### 2.3 编码实现

#### 2.4 json版rpc

- 使用nc -l 127.0.0.1 8800 充当服务器
- 02-client.go 充当客户端，发起通信 ---> 乱码
  - 原因：RPC使用go语言特有的数据序列化gob。其他编程语言不能解析
- 使用通用的序列化、反序列化 ---- json、protobuf
  - 修改客户端，使用jsonrpc

    ```go
    conn, err := jsonrpc.Dial("tcp", "8800")
    ```

    使用nc -l 127.0.0.1 8800充当服务器

    结果：

    ```json
    {"method":"hello.HelloWorld","params":["李白"],"id":0}
    ```

  - 修改服务器，修改绑定服务的部分，使用jsonrpc

    ```go
    jsonrpc.ServeConn(conn)
    ```

    受用nc 127.0.0.1充当客户端

    ```shell
    echo -e '{"method":"hello.HelloWorld","params":["李白"],"id":0}' | nc 127.0.0.1 8800
    ```

    如果返回值的error不为空，无论传出参数是否有值，服务端都不会返回数据

#### 2.5 rpc封装

## 补充

Herman Schaaf基准测试，测试字符串拼接各种方法的性能

