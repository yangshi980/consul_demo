package demo
import (
   "fmt"
  consulapi "github.com/hashicorp/consul/api"
  consulwatch "github.com/hashicorp/consul/api/watch"
)

func Watch(client *consulapi.Client,key string) {
  params := map[string]interface{}{}
  params["type"] = "key"
  params["key"] = key
  plan, err := consulwatch.Parse(params)
  if err != nil {
    // 错误处理
  }

  // 当所监视的key发生任何更改时调用该handler函数
  plan.Handler = handler

  //  阻塞调用，因此这段代码应该在goroutine中运行
  plan.RunWithClientAndHclog(client, nil)
}
func handler(index uint64, result interface{}) {
  fmt.Printf("watch data: %s", result)
  // 检查返回的键是否有sessionID
  // 如果session ID存在，则key被某个节点获取
  // 如果没有，目前没有领导者，尝试获取key
}
