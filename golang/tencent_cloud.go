一面
————————————————————————————————————————
1.自我介绍
2.项目介绍，问的挺多的，每个项目的架构以及实现方式

3.etcd相关

是什么？如何保持高可用性，选举机制，脑裂如何解决

4.k8s相关

哪些常用组件，发起一个pod的创建整个通路，service有哪些，

一个请求到达pod的过程、configmap、 dockerfile


线程和协程：

线程是系统级实现的，依赖于操作系统，调度方法是时间片轮转法；

协程是应用软件级实现的，当一个协程调度到另一个协程时，将上一个协程的上下文信息压入堆；

golang的协程调度基于GPM模型实现；


5.golang基础

a 数组和切片的区别

数组是定长，数组是“值类型”，赋值将会赋予值拷贝，

切片是可变长度，切片是“引用类型”，赋值将会赋予指针地址，切片初始化时候cap和len是一样的，当cap不够用时候扩容是按照每2倍数len进行扩容；

b 协程同步的方式

其实这个问题等于是说，如何使用多协程的情况保证并发安全；

1 Mutex互斥锁，用互斥锁防止同时被多个协程操作同一个变量，是多个协程变得有序；

2 golang的channel，利用channel的无缓冲管道，一个输入的必须有一个接收的形成有序，比如生产者消费者模型；

3 WaitGroup的等待次数阻塞

c waitgroup和context区别

都可以用来保证所有协程执行完毕；

但是waitgroup主要是用于阻塞主协程的；

WaitGroup 在 sync 包中，用于阻塞主线程执行直到添加的 goroutine 全部执行完毕；
而context可以多个goroutine传递信号，超时、取消、截止、k-v

Context 包不仅实现了在程序单元之间共享状态变量的方法，同时能通过简单的方法，使
我们在被调用程序单元的外部，通过设置ctx变量值，将过期或撤销这些信号传递给被调用的程序单元；子协程也可以通过监听关闭信号进行回收关闭goroutine


d 如何处理异常defer

defer先进先出fifo，return之前会调用，主要会用于资源清理；

关于异常和错误，我们上层使用recover捕获异常防止继续往上抛出（创建一个goroutine处理一个request时候），错误会透传，
而异常（文件权限错误、文件打开异常、redis或者mysql连接失败，我们都会认为是一个异常）会直接panic，并发送邮件告知开发者；


e 通用的http请求日志打印如何封装

通常是创建一个log中间件在路由处理前置中间件，按照请求地址、参数、请求时间、方式等记录

f init函数介绍

g 简述golang的包管理

go mod init将当前项目定义为一个模块，下面每一个目录的定义为一个package,引入包规则为模块名+路径格式
注意 import后面带的是完全等价于module name + 文件夹路径的;而一个文件夹里面只能一个package,但是这个package的名字可以和文件夹名字不一致

6.tcp三次握手四次挥手 可靠性如何保证

首先我们理一下UDP和TCP是什么：

UDP，特点是不需要建立连接，远程主机收到报文以后也不需要给出任何确认；因为不需要建立连接保证交付所以省去很多开销，他的速度更快，对应的应用层的协议有：DNS(域名服务)、TFTp（简单文件传输协议）

TCP：面向连接、传送数据之前要先连接，传送完成以后要释放连接，SMTP、Telnet、HTTP、ftp

tcp的3次握手，其实可以看出来，服务端的资源更为珍贵，客户端请求服务端连接，服务端响应可以连接，客户端再响应已经收到可以连接的信号。其实是最大程度保证服务端的资源不浪费，因为如果不珍惜服务端资源的话，在
第一次服务端回应可以连接的时候，已经建立连接了；

而4次挥手的话，通俗的解释就是：首先客户端发送连接释放的信号；服务端确认可以关闭连接，客户端收到服务端的确认请求（客户端进入终止等待状态），服务端将最后的数据发送完毕后，向客户端发送连接释放报文，此时又要客户端

对服务端的释放连接信号发出确认，服务端收到客户端的确认才算是进入closed状态，总结一来一回的发生了4次，叫4次挥手：

1 客户端发出释放连接信号，服务端响应并进入关闭等待  --  挥手一次  ||  顾客发送停止服务信号，卖家进入服务收尾（关闭等待）状态

2 服务端发送进入关闭等待状态的确认信号，客户端接收信号进入终止等待状态 || 卖家回应已经准备在服务收尾中（算账单），顾客进入终止等待状态（准备付款二维码）

3 服务端将所有数据发送完毕后发送连接释放信号，客户端响应信号并发出确认，进入closed || 卖家已经算完账单服务结束了，告诉顾客，顾客付完款已经离开店了

4 服务端接收到客户端发送的确认，进入closed状态 || 卖家确认收款，把店关了

其实很简单，就是一个互相确认的过程，最重要的是记住：四次挥手之中每一次，服务端或者客户端都有一个状态变更；

而可靠性如何保证呢？

序列号和确认应答信号
超时重发控制
连接管理
滑动窗口控制
流量控制
拥塞控制

至于这几个属于的解释呢，我啥也不知道


举一反三：tcp/IP属于计算机网络OSI七层模型之中的哪一层呢 - 传输层（建立管理端与端之前的连接）
而http/ftp之类属于应用层（为程序提供网络服务）

TCP 是"面向连接"的
UDP 是"面向报文"的

7.redis数据类型

回看《redis入门指南》

1、字符串类型 key-value
    简单的一对一映射关系，刚开始采用序列化数组为字符串存储 { name => 'jck' }
2、散列类型
    存储对象，并且可以直接对对象的某一个属性进行增删改查 {id => {'name':'jack','title'=>'taitannike','time'=>'2018-08-08 15:20'}}
3、列表类型（改变某一个的顺序需要把列表的每一个元素全部重新排列）
    获取文章列表分页数据的时候，需要用列表，但是依然需要解决排序、取中间片段速度缓慢的问题  [1,2,3,4,5]
4、集合类型
    存储文章标签的时候，如何做到标签唯一，有就无操作无就存入. id => {"美食","旅行","装饰"}
5、有序集合类型 - (散列+跳跃表实现，所以中间存取也是极快)
    号称最高级的数据类型，就是序号加值，但是他可以做到按序号大小获取中间片段，以及按大小排序 {score => {['89','tom'],['99','marry']}}

8.linux查看端口占用命令

netstat -tnpl | grep 8899

7.mysql相关

存储引擎 区别 索引的种类 查询较慢的时候如何分析

8.算法

两个栈实现一个队列
————————————————————————————————————————
二面（主要是项目的深入考察）
————————————————————————————————————————
1.自我介绍
2.项目介绍
3.tcp四次挥手、time_wait状态
4.linux常用命令，使用shell拆分一个ip地址
5.网络不可达如何排查，例如我当前打不开qq.com
6.k8s内部请求到达外部服务的过程
————————————————————————————————————————
我又来更新了^_^
三面（主要是个人规划、看法）
————————————————————————————————————————
1.自我介绍
2.项目介绍（具体干啥、作用、背景、成果、遇到的问题、解决方案）
3.基础相关
    tcp、udp区别、进程线程区别
4.项目相关
    etcd mvcc、k8s pod之间如何通信
5.linux相关
    如何排查网络问题、命令
5.个人规划
6.对于部门工作的了解（不了解，卒）
7.对自己的要求
————————————————————————————————————————
本人小菜鸡1年golang后台开发经验

整体来说 问的不难 更多注重项目、基础 所有问题基本都是从项目中出发，牛客保佑，希望能有hr面😃