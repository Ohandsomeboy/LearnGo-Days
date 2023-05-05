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

