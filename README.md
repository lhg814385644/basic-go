# basic-go
basic go learn

# github ssh配置（Linux Ubuntu22.04 LTS）

https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent

1. ssh-keygen -t ed25519 -C "your_email@example.com"
2. 特别注意Adding your SSH key to the ssh-agent
3. 执行eval "$(ssh-agent -s)" （这一步是在后台启动ssh-agent）
4. 执行exec ssh-agent bash
5. 执行ssh-add ~/.ssh/id_ed25519
6. 执行完如上操作后基本上就可以SSH连接了

# 安装Oh-My-Posh
https://ohmyposh.dev/docs/installation/linux
1. 安装到Home目录下（如果没有如下目录请先新建）
   curl -s https://ohmyposh.dev/install.sh | bash -s -- -d ~/ohmyposh/bin
2. 配置环境变量(在Home目录下编辑.profile文件)
   export PATH=$PATH:~/ohmyposh/bin
   eval "$(/home/hg/ohmyposh/bin/oh-my-posh init bash --config ~/.cache/oh-my-posh/themes/easy-term.omp.json)"
3. 重新加载配置文件
   source .profile

# 1期仓库
https://gitee.com/geektime-geekbang_admin/geektime-basic-go/tree/record


# Linux（Ubuntu22.04 LTS）上安装kubectl
官方文档：https://kubernetes.io/zh-cn/docs/tasks/tools/install-kubectl-linux/
1. 下载kubectl(由于我的k8s版本是1.29.2，所以下载的是1.29.2的kubectl)
curl -LO https://dl.k8s.io/release/v1.29.2/bin/linux/amd64/kubectl
2. 安装kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
3. 验证安装
kubectl version --client


### 初步在k8s中部署Web应用
1. 创建一个Deployment(创建完成后执行如下命令)
kubectl apply -f webook-deployment.yaml
2. 创建一个Service(创建完成后执行如下命令)
kubectl apply -f webook-service.yaml