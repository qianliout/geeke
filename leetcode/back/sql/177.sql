/*
编写一个 SQL 查询，获取 Employee 表中第 n 高的薪水（Salary）。

+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+

例如上述 Employee 表，n = 2 时，应返回第二高的薪水 200。如果不存在第 n 高的薪水，那么查询应返回 null。

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+

*/


create function getNthHighestSalary(n int) returns int
begin
  declare p1 int;
  declare p2 int;
  if (n < 1)
  then
    set p1 = 0,p2 = 0;
  
  else
    set p1 = n - 1,p2 = 1;
  end if;
  
  return (
    select ifnull(
             (select distinct Salary from Employee order by Salary desc limit p2 offset p1),
             null
             )
  );
end;

