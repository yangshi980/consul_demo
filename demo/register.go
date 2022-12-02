package demo
import (
   "fmt"
  consulapi "github.com/hashicorp/consul/api"
)
 
const (
   ConsulAddress = "127.0.0.1:8500"
   LocalIP   = "127.0.0.1"
   LocalPort = 8005
)
// 这里使用的是传参形式进行服务注册.可以根据需求进行函数调用.
func Reg(host, name, id string, port int) (*consulapi.Client,error) { //服务注册需要的服务名称,服务id(后续会有随机id产生) 服务当前对应的端口号(随机端口号)
   defaultConfig := consulapi.DefaultConfig()
   defaultConfig.Address = ConsulAddress
     client, err := consulapi.NewClient(defaultConfig) //生成consul客户端
     if err != nil {
        panic(err)
     }
     registration := &consulapi.AgentServiceRegistration{  //被注册服务不是consul服务本身
        ID:      id,  //这里是被注册的服务的id
        Name:    name, //被注册的服务的name
        Tags:    []string{}, //被注册服务的tags
        Port:    port, //被注册服务的port
        Address: host, //被注册服务的host
     }
     serverAddr := fmt.Sprintf("http://%s:%d/health", registration.Address, registration.Port)  //监控检测是检查的被注册服务.
     check := &consulapi.AgentServiceCheck{   //注意这里监控检查的配置要声明时间单位.不能只填写数字.只填数字默认单位不是秒
        Interval:                       "1s",  //检查间隔
        Timeout:                        "5s",  //检查超时时间
        HTTP:                           serverAddr,  //这里是被注册服务监控检测地址
        DeregisterCriticalServiceAfter: "20s",  //如果服务20s没有响应就将注册的服务删除.
     }
   
     registration.Check = check  //进行健康检查
    fmt.Printf("%T\n",client)
    return client,client.Agent().ServiceRegister(registration)  
  }
