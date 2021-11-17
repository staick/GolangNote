# GolangNote
My way to learn Golang

## Socket

### 1.Server Demo(单次连接)

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

### 4.Client Demo(多次连接) 
