/*
表1: Person
+-------------+---------+
| 列名         | 类型     |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId 是上表主键
表2: Address
+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId 是上表主键
编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：
FirstName, LastName, City, State
*/

select FirstName, LastName, City, State from Person left join Address on Person.PersonId=Address.PersonId;


# 以后采用这种写法更好
select a.FirstName as FirstName,a.LastName as LastName,b.City as City,b.State as State from Person a left   join  Address b on a.PersonId = b.PersonId;
