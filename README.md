# WxRobotGo
webchatRobot

微信公众号聊天机器人

  基于图灵机器人接口
  聊天数据存储到mysql
  自动建表，一键执行

  配置:
    go get .
 
  运行:
    go run main.go -config=config.yaml
 
  后台服务(默认linux):
    go build main.go
    nohup ./main -config=config.yaml >log 2>&1 &
