# WxRobotGo

## 微信公众号聊天机器人<br>

  基于图灵机器人接口<br>
  聊天数据存储到mysql<br>
  自动建表，一键执行<br>
  
配置:<br>
  go get .<br>

运行:<br>
  go run main.go -config=config.yaml<br>
 
后台服务(默认linux):<br>
  go build main.go<br>
  nohup ./main -config=config.yaml >log 2>&1 &
