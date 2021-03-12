select u.id, u.name as userName, up.name as parentUserName
 from user u 
 left join user up on u.parentId = up.id 
 order by id asc