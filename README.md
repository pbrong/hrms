# hrms
基于go、gorm、gin、mysql及layui构建的人力资源管理系统。提供员工管理、考试管理、薪资考勤管理、权限管理及分公司分库数据隔离等功能。欢迎Star或提Issue。

# 使用方式
- git clone https://github.com/pbrong/hrms.git
- cd hrms && go mod tidy
- 按照sql文件的两个配置，分别建hrms1和hrms2分公司数据库
- 更新conf配置文件配置
- sh build.sh 执行脚本编译可执行文件执行 或 直接启动main.go运行

# 功能结构
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-47-31.png)


# 系统架构
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-32-26.png)

# 数据库设计
共14张数据库表，ER关系如下：
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-29-52.png)

# 分公司分库设计
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-58-07.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-58-27.png)

# 权限设计
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-32-41.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-32-15.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Jan%2001%2012-32-51.png)

# 界面展示
- 分公司员工登陆
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-51-43.png)
- 超级管理员、企业管理员及普通员工
- ![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-52-57.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-52-28.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-52-41.png)
- 权限管理
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-53-42.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-54-24.png)
- 薪酬管理
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-54-53.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-55-03.png)
- 考勤上报
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-55-37.png)
- 招聘管理
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-56-06.png)
- 候选人管理
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-56-14.png)
- 考试管理
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-56-31.png)
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-56-47.png)
- 考试答题
![](https://github.com/pbrong/pbrong/blob/main/Screenshot%20at%20Dec%2015%2021-57-01.png)

