select a.id as ID, a.name as UserName, b.name as ParentUserName
from user a left join user b on b.id = a.parent