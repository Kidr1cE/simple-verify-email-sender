# simple-email-verify
## Overview
基于[jordan-wright/email](https://github.com/jordan-wright/email)封装，通过yml配置文件的网易163验证邮箱发送模块。  
## Usage
1. 开启网易邮箱`IMAP/SMTP服务`并记录认证密码
2. 配置`config.yml`
```yaml
address: smtp.163.com:25
auth:
  host: smtp.163.com
  username: xxx@163.com
  password: pass
verify:
  type: token #token/code
template:
  name: Alco
  subject: Verify Email
```
3. 编写自定义HTML邮件(可选)
4. 使用
```go
func main() {
	config, err := verify.LoadConfig("./config.yml")
	if err != nil {
		return 
	}
	var sender verify.VerifyEmailSender

	err = sender.ReadConfig(config)
	if err != nil {
		return
	}
}
```