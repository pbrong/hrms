# hrms
基于go、gorm、gin、mysql及layui构建的人力资源管理系统。提供员工管理、考试管理、通知管理、薪资考勤管理、招聘管理、权限管理及分公司分库数据隔离等功能。欢迎Star或提Issue。

# 开源声明
本项目用于Go爱好者学习和参考，不能直接用于生产环境，转载使用请说明出处。golang开发交流微信群，如二维码过期可添加微信arong2048，备注github。
<img width="265" alt="image" src="https://github.com/pbrong/hrms/assets/38597513/72a20f9e-0d0c-40c2-afa9-fde09110d98c">


# 维护日志
- [x] 2021-02-15 需求分析设计、数据库设计及项目搭建(go、gin、gorm、layui)
- [x] 2021-03-21 RBAC及模板引擎实现分权限化模块管理（supersys、admin、normal)
- [x] 2021-03-24 分公司分库数据隔离模块（数据库实例动态筛选）
- [x] 2021-03-25 部门管理、职级管理及员工管理模块
- [x] 2021-04-12 新闻管理及发布通知模块
- [x] 2021-03-25 薪酬管理、薪资发放、薪资套账模块
- [x] 2021-03-26 考勤管理、考勤上报模块
- [x] 2021-04-01 招聘管理、候选人管理模块
- [x] 2021-12-19 补充sql文件，完善说明
- [x] 2022-11-16 更新项目配置
- [x] 2022-12-03 新增员工批量导入模块

# 项目分层
- README.md// 项目说明
- build.sh // 编译脚本
- config   // 配置文件
- go.mod   // go依赖列表文件
- go.sum   // go依赖校验文件
- handler  // 路由层
- hrms_app // 编译后的打包文件
- main.go  // 启动文件
- model    // 实体层
- resource // 配置层
- service  // 业务层
- sql      // 所用到的sql文件
- excel    // 上传测试文件
- static   // 静态资源
- views    // 前端文件
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
