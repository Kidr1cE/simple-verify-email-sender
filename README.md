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
  username: xxx@163.com # Your email addr
  password:  # Your SMTP password
verify:
  type: token # token/code
template:
  name: # Your name
  subject: Verify Email
```
3. 编写自定义HTML邮件(可选)
4. 使用
```go
func main() {
	// read config
	file, err := os.ReadFile("./config.yml")
	config, err := verify.LoadConfig(file)

	// inint email sender
	var sender verify.VerifyEmailSender
	err = sender.ReadConfig(config)
	if err != nil {
		fmt.Println(err)
	}

	// send email
	sender.SendTo("784312513@qq.com", "https://yourwebsite/verity/?token=114514")
}
```