/*
编写一个 SQL 查询，查找所有至少连续出现三次的数字。

+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+

例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。

+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+

*/

/*
注意,这里是连续出现3次,而不是说一共出现3次,如是是一共出现3次应该怎么搞呢
*/


select distinct l1.Num as `ConsecutiveNums`
from Logs l1,
     Logs l2,
     Logs l3
where l1.Id = l2.Id - 1
  and l2.Id = l3.Id - 1
  and l1.Num = l2.Num
  and l2.Num = l3.Num;

/*
#①首先遍历一遍整张表，找出每个数字的连续重复次数
#具体方法为：
    #初始化两个变量，一个为pre，记录上一个数字；一个为count，记录上一个数字已经连续出现的次数。
    #然后调用if()函数，如果pre和当前行数字相同，count加1极为连续出现的次数；如果不同，意味着重新开始一个数字，count重新从1开始。
    #最后，将当前的Num数字赋值给pre，开始下一行扫描。
    select
        Num,    #当前的Num 数字
        if(@pre=Num,@count := @count+1,@count := 1) as nums, #判断 和 计数
        @pre:=Num   #将当前Num赋值给pre
    from Logs as l ,
        (select @pre:= null,@count:=1) as pc #这里需要别名
    #上面这段代码执行结果就是一张三列为Num,count as nums,pre的表。

#②将上面表的结果中，重复次数大于等于3的数字选出，再去重即为连续至少出现三次的数字。
    select
        distinct Num as ConsecutiveNums
    from
        (select Num,
                if(@pre=Num,@count := @count+1,@count := 1) as nums,
                @pre:=Num
            from Logs as l ,
                (select @pre:= null,@count:=1) as pc
        ) as n
    where nums >=3;

#注意：pre初始值最好不要赋值为一个数字，因为不确定赋值的数字是否会出现在测试表中。
*/


# Write your MySQL query statement below
select distinct Num as ConsecutiveNums
from (select Num,
             if(@pre = Num, @count := @count + 1, @count := 1) as nums,
             @pre := Num
      from Logs as l,
           (select @pre := null, @count := 1) as pc
     ) as n
where nums >= 3;


# 另一种写法
SELECT DISTINCT a.Num ConsecutiveNums
FROM (
         SELECT t.Num,
                @cnt := IF(@pre = t.Num, @cnt + 1, 1) cnt,
                @pre := t.Num                         pre
         FROM Logs t,
              (SELECT @pre := null, @cnt := 0) b) a
WHERE a.cnt >= 3;


SELECT DISTINCT num AS ConsecutiveNums
FROM (SELECT id,
             num,
             CASE
                 WHEN @n = num THEN @c
                 ELSE @c := @c + 1
                 END AS rk
      FROM logs as l
               JOIN (SELECT @n := NULL, @c := 0) as tmp
      ORDER BY num, id) t1
GROUP BY num, (id - rk)
HAVING count(*) >= 3;

# 代码块
select distinct dd.Num ConsecutiveNums
from (
         select d.Num,
                @n := if(@pre = Num, @n + 1, @n := 1) count,
                @pre := Num
         from Logs d,
              (select @pre := null, @n := 1) r) dd
where dd.count >= 3;






