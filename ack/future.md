

# ACK

## optimize
- ack全系列产品的适配增强，自动拉取所有关联resource和datasource
- ack required字段自动推荐
- 针对ack产品的template增强，tags, runtime等带有elem的字段需要透出

### ack深度优化
**规则**
- 命令行只输入ack resource时触发优化，例如`terrafmtter --resource cs_kubernetes`。
- data source在前，resource在后面输出

一些可能的优化项:
- cs_kubernetes_node_pool、cs_kubernetes_autoscaler、cs_kubernetes_addon需要关联到cluster
- cs_kubernetes，和cluster有关的一系列resource需要关联到依赖资源

