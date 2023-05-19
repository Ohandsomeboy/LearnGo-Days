# 从零开始Golang学习之路
这个项目旨在记录平时刷题和一些小项目的心得，也是自己对自己的一个激励

# 2023.4.14
对之前的leetcode里面的题目进行整理和复习，开始查找如何创建一个库
leetcode1023:驼峰匹配，法1：双指针法  法2：正则

# 2023.4.17
leetcode 387:第一个不重复的字符：

# 2023.4.18
leetcode 1.3: 二刷1.3题

# 2023.4.19
leetcode 1768:交替合并字符串

# 2023.4.22
leetcode 1431：拥有最多糖果的孩子

# 2023.4.23
leetcode 603：种花问题（贪心问题）

# 2023.4.24
leetcode 345：反转字符串中的元音字母 strings.continer()检查是否存在字串

# 2023.5.4
review leetcode golang 1.3 SQL:175：复盘面试；
开组会；今天收到了论文拒稿；

# 2023.5.5
多表联结：（7种）
select ... from 表1 as a left join
表 2 as b on a.列名=b.列名

select ... from 表1 as a right join
表 2 as b on a.列名=b.列名

select ... from 表1 as a inner join
表 2 as b on a.列名=b.列名

select ... from 表1 as a left join
表 2 as b on a.列名=b.列名 where b.列名 is null

select ... from 表1 as a right join
表 2 as b on a.列名=b.列名 where b.列名 is null

1、char_length(str)
（1）计算单位：字符
（2）不管汉字还是数字或者是字母都算是一个字符
2、length(str)
（1）计算单位：字节
（2）utf8编码：一个汉字三个字节，一个数字或字母一个字节。
（3）gbk编码：一个汉字两个字节，一个数字或字母一个字节。

空值判断：
MySQL 使用三值逻辑 —— TRUE, 
FALSE 和 UNKNOWN。任何与 NULL 值进行的比较都会与第三种值 UNKNOWN 做比较。
这个“任何值”包括 NULL 本身！
这就是为什么 MySQL 
提供 IS NULL 和 IS NOT NULL 两种操作来对 NULL 特殊判断。

# 2023.5.5
leetcode 9 15
sql 1378 1068
整理面经

# 2023.5.7
leetcode 15 20 21
sql 197
方法一：lag()+datediff()
lag()和lead()这两个函数可以查询我们得到的结果集上下偏移相应行数的相应的结果。

lag()函数：
查询当前行向上偏移n行对应的结果
该函数有三个参数：第一个为待查询的参数列名，第二个为向上偏移的位数，第三个参数为超出最上面边界的默认值。

lead()函数：
查询当前行向下偏移n行对应的结果
该函数有三个参数：第一个为待查询的参数列名，第二个为向下偏移的位数，第三个参数为超出最下面边界的默认值。

DATEDIFF()

DATEDIFF() 函数返回两个日期之间的时间。

day=DiffDate=enddate−startdate=1，把示例的day换成year，就是按年计算


select id from
(select id, temperature, recordDate, lag(recordDate,1) 
over(order by recordDate) as last_date,
lag(temperature,1) over(order by recordDate) as last_temperature
from Weather) a where temperature > last_temperature and 
datediff(recordDate, last_date) = 1

方法二：笛卡尔积
获取所有比前一天温度高的记录并记录其id
找相关联表 题干涉及到的表只有单表Weather
分析：

我们想要拿记录的温度与前一天进行相对比，那必然需要将前一天的记录放到一行中，不然单纯的select是无法进行比较的；所以想到用笛卡儿积，然后再通过前置知识进去过滤+温度条件

三、写SQL 将1步骤对应条件拼接到步骤2后面即可得到SQL

小知识：笛卡尔积
笛卡尔积是指在数学中，两个集合X和Y的笛卡尓积（Cartesian product），又称直积，表示为X×Y，
第一个对象是X的成员而第二个对象是Y的所有可能有序对的其中一个成员。
笛卡尔积又叫笛卡尔乘积，是一个叫笛卡尔的人提出来的。 简单的说就是两个集合相乘的结果。

假设集合A={a, b}，集合B={0, 1, 2}，则两个集合的笛卡尔积为{(a, 0), (a, 1), (a, 2), (b, 0), (b, 1), (b, 2)}。

inner join(等值连接) 只返回两个表中联结字段相等的行on表示链接条件

select b.Id from weather a inner join weather b
where DATEDIFF(b.recordDate,a.recordDate)=1 and b.Temperature > a.Temperature;

方法三：TIMESTAMPDIFF()
TIMESTAMPDIFF能干什么，可以计算相差天数、小时、分钟和秒，相比于datediff函数要灵活很多。格式是时间小的前，时间大的放在后面。 计算相差天数：

select w1.Id
from Weather as w1, Weather as w2
where TIMESTAMPDIFF(DAY, w2.RecordDate, w1.RecordDate) = 1
AND w1.Temperature > w2.Temperature

方法四：adddate()函数
ADDDATE()函数用于将指定的时间间隔添加到日期值date是表示日期的值，它可以是 String、DATE（YEAR、MONTH 和 DAY）、DATETIME（HOURS、MINUTES 或 SECONDS）或 TIMESTAMP 类型。

ADDDATE( date , INTERVAL value addunit )

ADDDATE( date , days )

# 2023.5.9
leetcode 34.21.283(双指针)
1.了解了Context和channel+select的使用场景

# 2023.5.11
leetcode 自己记录得前5题
golang 整理相关面试资料，重新梳理做过得项目

# 2023.5.12
leetcode 15.34.20
开组会
小林:图解MySQL

# 2023.5.18
上周三轮面试，真的地狱。累死娃了
这周开始小林图解：计网、Mysql、系统
同时开始，一天至少一个专题的刷leetcode（在六月前刷完75题，去拿牌子）
