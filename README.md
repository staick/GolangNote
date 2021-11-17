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

