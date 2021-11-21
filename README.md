<<<<<<< HEAD
# GolangNote
My way to learn Golang

## Socket

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

## http

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

## 网络聊天室

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

### 实现基础

1. tcp socket，建立多个连接

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

   



=======
# GolangNote
Go语言学习记录

## Socket

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

## http

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

>>>>>>> b0e492457a4de9e83a9d99a62e4a1fe8668c3846
