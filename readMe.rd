
# 拉取并整理依赖
go mod tidy

# 强制下载依赖（有时缓存问题）
go clean -modcache
go mod download

#查看当前的remote列表
git remote -v
推送base代码到github仓库
git remote add public-origin https://github.com/zoumaguanhu/king-base.git
# 需从本地master推送代码，注意dev合并代码到master
git push -u public-origin master