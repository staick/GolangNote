package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//定义用户属性
type User struct {
	id   string
	name string
	msg  chan string
}

//定义用户集合
var allUser = make(map[string]User)

//定义一个全局的message通道
var message = make(chan string, 10)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	fmt.Println("服务器启动成功！")

	go boardcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		go handler(conn)
	}
}

//业务函数
func handler(conn net.Conn) {
	fmt.Println("业务处理启动成功！")

	//获取客户端地址
	clientAddr := conn.RemoteAddr().String()

	//创建用户
	newUser := User{
		id:   clientAddr,
		name: clientAddr,
		msg:  make(chan string, 10),
	}

	//将用户添加到map
	allUser[clientAddr] = newUser

	//定义一个退出信号，用于标识client退出
	var isQuit = make(chan bool)

	//创建一个定时器管道，用于告知watch函数，用户正在输入
	var resetTimer = make(chan bool)

	//启动go程，负责监听客户端退出信号
	go watch(&newUser, conn, isQuit, resetTimer)

	//启动go程，将msg信息返回给客户端
	go writeBack(&newUser, conn)

	//上线通知
	loginInfo := fmt.Sprintf("[%s]:[%s]上线了！\n", newUser.id, newUser.name)
	message <- loginInfo

	for {
		buf := make([]byte, 1024)

		cnt, err := conn.Read(buf)

		if cnt == 0 {
			fmt.Println("客户端主动关闭Ctrl + c. 准备退出！")
			isQuit <- true
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("接收客户端的消息：", string(buf[:cnt]), "，数据长度：", cnt-1)
		// message <- fmt.Sprintf("[%s]:%s",newUser.name, string(buf[:cnt]))

		//查询当前所有的用户/who
		userInput := string(buf[:cnt-1]) //用户输入的数据，最后一个是回车，记得去掉
		//a.先判断接收的数据是不是/who ===> 长度&&字符串
		if len(userInput) == 4 && userInput == "/who" {
			//b.遍历allUsers这个map，将id和那么拼接成字符串，返回给客户端
			fmt.Println("用户即将查询所有信息！")

			//这个切片包含所有的用户信息
			var userInfos []string

			for _, user := range allUser {
				userInfo := fmt.Sprintf("userid:%s, username:%s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}

			//最终写到管道中，一定是一个字符串
			r := strings.Join(userInfos, "\n") //连接数字切片，生成字符串

			//将数据返回给查询的客户端
			newUser.msg <- r

			//修改用户名/rename
			//规则：/rename Duke
			//1.读取数据判断长度大于7，判断字符是/rename
			//2.使用 进行分割，获取 后面的内容，作为名字
			//3.更新用户名字newUser.name = Duke
			//4.通知客户端，更新成功
		} else if len(userInput) > 7 && userInput[:7] == "/rename" {
			newUser.name = strings.Split(userInput, " ")[1]
			allUser[newUser.id] = newUser	//更新map中的user

			newUser.msg <- "rename successfully!\n"	//通知客户端，更新成功！
		} else if len(userInput) == 5 && userInput == "/quit" {
			isQuit <- true
		} else {
			message <- fmt.Sprintf("[%s]:%s\n", newUser.name, userInput)
		}

		resetTimer <- true
	}

}

func writeBack(user *User, conn net.Conn) {
	fmt.Printf("user:%s的go程正在监听自己的msg管道\n", user.name)
	for data := range user.msg {
		_, _ = conn.Write([]byte(data))
	}
}

func boardcast() {
	fmt.Println("广播启动成功!")
	defer fmt.Println("boardcast程序退出！")
	for {
		//从message中读取数据
		fmt.Println("boardcast监听message中...")
		info := <-message

		fmt.Println("message接收消息：", info)

		//将数据写入到每一个用户的msg管道中
		for _, user := range allUser {
			user.msg <- info
		}
	}
}

//每个用户都有自己的watch go程，负责监听退出信号
func watch(user *User, conn net.Conn, isQuit <- chan bool, resetTimer <- chan bool) {
	fmt.Println("启动监听退出信号的go程...")
	defer fmt.Println("watch go程退出！")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s 下线了!\n", user.name)
			user.msg <- "您已下线，请按任意键继续..."
			time.Sleep(500)
			fmt.Println("删除当前用户：", user.name)
			delete(allUser, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-time.After(60 * time.Second):
			logoutInfo := fmt.Sprintf("%s 停滞太久，强制下线!\n", user.name)
			user.msg <- "您因停滞太久已强制下线，请按任意键继续..."
			time.Sleep(500)
			fmt.Println("删除当前用户：", user.name)
			delete(allUser, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-resetTimer:
			fmt.Println("连接%s重置计时器！\n", user.name)
		}
	}
}
