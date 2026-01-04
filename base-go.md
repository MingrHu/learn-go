linux环境下：
1.设置基础
sudo apt update && sudo apt upgrade -y
sudo mkdir -p /usr/local/go

2.安装go编译器
# 替换版本号为你需要的Go版本（如go1.22.0）
cd ~
wget https://dl.google.com/go/go1.25.5.linux-amd64.tar.gz
# 解压
sudo tar -zxvf go1.21.6.linux-amd64.tar.gz -C /usr/local/

3.环境变量设置
sudo vi /etc/profile
# 配置GOROOT：Go的安装目录
export GOROOT=/usr/local/go
# 配置GOPATH：Go工作区目录
export GOPATH=$HOME/go
# 配置PATH：将Go的bin目录添加到系统环境变量，让系统识别go命令
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# 可选：开启Go模块支持（Go 1.11+推荐，用于依赖管理）
export GO111MODULE=on
# 可选：配置Go模块代理（加速国内依赖下载，推荐阿里云代理）
export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
# 激活
source /etc/profile

4.vscode的go插件设置
# 配置阿里云Go代理
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
vscode自动完成插件的工具链安装


go文件编译运行
# 编译go文件（生成可执行文件hello）
go build hello.go
# 运行可执行文件
./hello
# 也可直接运行（无需手动编译，go会自动临时编译并执行）
go run hello.go

MacOs环境下
1.下载go-1.25 darwin/arm-64bit的pkg
2.安装pkg 自动部署和设置环境变量
3.配置代理（阿里云代理，同样稳定）
echo 'export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct' >> ~/.zshrc
source ~/.zshrc
4.在vscode里面安装go插件和go的工具链，完成后可进行调试、补齐等