## 部署位置
nc066:/disk1/suns/gowork/src/searchq-operationsys
## 启动方式
项目目录下执行：bee run
## 框架
beego 框架、golang语言
## 脚本
### 位置：
ns014：/disk1/suns/dict/dict-searchq/target/day.py
### 作用：
实现将hadoop上的每天日志导入到mysql中，利用了dict-searchq工程中的parser-uq-log2和parseLog2MysqlTool两个工具类
## 数据库
### 位置
mysql（hd016）：log
### 表格
- log

 `id` int(11) NOT NULL AUTO_INCREMENT,  //日志id

  `user` varchar(30) NOT NULL,      用户客户端
  
  `imei` varchar(50) NOT NULL,      设备号
  
  `action` varchar(50) NOT NULL,    使用接口
  
  `image` varchar(100) NOT NULL,    图片链接
  
  `date` varchar(10) NOT NULL,      日志时间
  
  `result` varchar(3000) NOT NULL,  日志内容
  
  PRIMARY KEY (`id`)
  
- review  评测表，记录一次评测总体数据

 `id` int(11) NOT NULL AUTO_INCREMENT,   评测ID
 
  `name` varchar(50) NOT NULL,          评测名字
  
  `beginTime` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, 评测时间段的起点
  
  `endTime` varchar(40) NOT NULL,        评测时间段的起点
  
  `summary` varchar(300) DEFAULT NULL,   评测的总结评论
  
  `num` int(11) DEFAULT NULL,            评测中包含的样本数量
  
  `type` int(11) DEFAULT NULL COMMENT '0 all, 1 multiple, 2 single ', 评测中的样本类型
  
- reviewQuestion 评测的样本数据
- 
  `id` int(11) NOT NULL AUTO_INCREMENT,  

  `idReview` int(11) NOT NULL,     评测表中的id
  
  `idQuestion` int(11) NOT NULL,    log表中的id
  
  `resultType` int(11) DEFAULT '2',  样本评测结果：有效（1）、无效（0）、未评（2） 
  
  `errorType` int(11) DEFAULT NULL,  无效类型 其他（0），模糊（1）、非K12（2）、横屏拍摄（3）、纯口算、计算（4）、纯手写作业（5）
  
  `grade` int(11) DEFAULT NULL,      年级 无法判断（0）、小学（1）、中学（2）
  
  `subject` int(11) DEFAULT NULL,    学科 其他（0）、理科（1）、数学（2）、英语（3）、文科（4）、语文（5）

  `allNumber` int(11) DEFAULT NULL,  页面题目个数
  
  `cutNumber` int(11) DEFAULT NULL,  切出题目个数
  
  `cutAccurateNumber` int(11) DEFAULT NULL,  切准题目个数
  
  `searchTrueNumber` int(11) DEFAULT NULL,   搜对题目个数:
  
   PRIMARY KEY (`id`)
   
 - reviewPartQuestion 样本切题的评测数据

  `id` int(11) NOT NULL AUTO_INCREMENT,
  
  `idQuestion` int(11) NOT NULL,  log的id
  
  `sequenceId` int(11) NOT NULL,  样本中切题的序列号
  
  `similarId` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, 匹配的相似模板号
  
  `cutAccurateNumber` int(11) NOT NULL, 是否切题切对
  
  `searchTrueNumber` int(11) NOT NULL,  是否切题切对
  
   PRIMARY KEY (`id`)
   
## 项目结构
db.go 数据库操作
review.go  评测的相关操作
bean.go 结构体
info 主页中的详情查看
default 主页的相关操作
## 页面结构
wel 主页
review 评测主页
wel2 添加评测后的样本页面
reviewinfo 样本的评测页面
wel3 评测的进度页面
re_result 评测的结果统计页面
