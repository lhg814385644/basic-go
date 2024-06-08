#!/usr/bin/bash

# 初始化SSH代理
eval "$(ssh-agent -s)"

# 添加私钥
ssh-add ~/.ssh/id_ed25519