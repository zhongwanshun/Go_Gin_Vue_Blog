#                                    BLOG

### 一、初始化项目,配置参数

* **前端使用Vue**
* **后端使用Gin**
* **数据库用MySql**

**使用Git管理项目，**

**1.安装Gin 框架:**

```
go get -u github.com/gin-gonic/gin
```

**2.项目框架**

**config：管理网站参数的一个文件夹，管理网站的相关配置参数，集中管理。**

​		config.ini：存放网站的相关配置参数。

**model：管理数据库的存储，读写的包**

​		Article：文章模型

​		Category：文章分类模型

​		Comment：评论模型

​		db：数据库模型，数据库的入口文件

​		Profile：详情模型

​		Upload：上传模型

​		User：用户模型，对用户功能的具体实现

**api:控制器，前后端分离，下面可以建立多个版本，做api入口。根据模型来写对应的接口**

article

category

comment

login

profile

upload

user：对用户的具体操作的控制器

**middleware:中间件，用于解决前后端分离跨域问题.**

**routes:路由的一个接口**

​	routers:路由的具体操作

**utils:用于工具的一个文件夹，公共功能需要全局使用的时候放在里面**

​		errmsg：错误处理工具包

​		setting：做一些数据处理。

**upload：上传下载的一个目录，托管静态资源。**

**web：用于前端页面静态托管。**

### 二、配置数据库，数据模型

### 三、架构错误处理模块和路由接口

实现api里面的架构完成，和错误处理。

### 四、用户模块接口，实现初步验证和分页功能

**添加用户：**

![image-20211228110915445](README.assets/image-20211228110915445.png)

![image-20211228111002859](README.assets/image-20211228111002859.png)

**展示用户(设置当前页数和每页显示的数量)**

![image-20211228113739502](README.assets/image-20211228113739502.png)

### 五、用户密码加密

### 六、编写编辑用户和删除用户接口