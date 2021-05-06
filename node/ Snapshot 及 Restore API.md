有一种情况是我们的所有的 node，或者有一部分 node 失败，可能会造成我们的数据的丢失。也就是说 replca 不能提供一种灾难性的保护机制。我们需要一种完整的备份机制。

注册仓库
在一个 snapshot 可以被使用之前，我们必须注册一个仓库（repository)。

使用 _snapshot 终点
文件夹必须对所有的 node 可以访问
path.repo 必须在所有的 node 上进行配置，针对一个 fs 的 repository 来说
PUT _snapshot/my_repo 
{
  "type": "fs",
   "settings": {
   "location": "/mnt/my_repo_folder"
  } 
}
这里 /mnt/my_repo_folder 必须加进所有 node 的 elasticsearch.yml 文件中。

