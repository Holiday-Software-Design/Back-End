# Back-End
寒假软件设计T3 后端项目仓库

使用viper，gin、gorm、goredis，并且由于mgo已经停止维护，故选择mongo-driver

接口文档地址

https://apifox.com/apidoc/shared-61642332-f0fa-4a35-8853-99662e0d9bcf

# 项目规范
go语言项目
- 一般变量使用小驼峰camelCase命名方式，避免缩写，避免使用下划线“_”
- 包外可见的变量名称首字母大写
- 命名常量时候全部大写
- 结构体内变量采用大驼峰CamelCase命名方式，首字母大写，以便其他包可见
- 接口命名为动词+er的形式，例如：Reader、Writer、Formatter。这样的命名模式能够清晰地表示接口的功能

接口
- 统一使用小驼峰camelCase命名方式
- 使用明确的相应状态码，其标准如下：
    - 2**成功
        - 200 OK： 请求成功。
        - 201 Created： 请求已创建成功，通常用于 POST 请求创建资源。
        - 204 No Content： 请求成功，但响应中没有内容，用于更新或删除资源
    - 3** 重定向
        - 301 Moved Permanently： 资源被永久移动到新位置。
        - 302 Found： 资源被临时移动到新位置。
        - 304 Not Modified： 客户端缓存仍有效，可以直接使用
    - 4xx 客户端错误：
        - 400 Bad Request： 请求无效或参数错误。
        - 401 Unauthorized： 需要身份验证。
        - 403 Forbidden： 服务器理解请求，但拒绝执行。
        - 404 Not Found： 请求的资源不存在。
    - 5xx 服务器错误：
        - 500 Internal Server Error： 服务器遇到错误，无法完成请求。
        - 502 Bad Gateway： 充当网关或代理的服务器从上游服务器收到无效响应。
        - 503 Service Unavailable： 服务器当前无法处理请求，通常是临时性的。


# 申报表部分
下列可申请项目结尾用（*）标识<br>
而且有分项的需要进行汇总计算得分
- 德育素质
    - 基本评定分D1
    - 记实加减分D2
        - 集体评定等级分（*）
        - 社会责任记实分（*）
        - 思政学习加减分
        - 违纪违规扣分
        - 学生荣誉称号加减分（*）

- 智育素质
    - 智育平均学分绩点

- 体育素质
    - 体育课程成绩T1
    - 课外体育活动成绩T2
        - 体育竞赛获奖得分（*）
        - 早锻炼得分

- 美育素质
    - 文化艺术实践成绩M1（*）
    - 文化艺术竞赛获奖得分M2（*）

- 劳育素质
    - 日常劳动分L1
        - 寝室日常考核基本分
        - “文明寝室”创建、寝室风采展等活动加分（*）
        - 寝室行为表现与卫生状况加减分
    - 志愿服务分L2
    - 实习实训L3

- 创新与实践素质
    - 创新创业成绩C1
        - 创新创业竞赛获奖得分（*）
        - 水平等级考试（*）
    - 社会实践活动C2（*）
    - 社会工作C3

通过观察文档我们可以看出，每个模块只有一个父项有子项，那么我们就可以通过在父项处加入F-开头的字段进行标识，在子相处加入L-开头字段命名表示

### tips
context.WithValue一般用来设置不用改变的值
而c.Set("")的方式一般用来设置可以改变的值