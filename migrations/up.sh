#bin/bash
#docker exec -i ddd.postgres  createdb go-ddd -U postgres
#docker exec -i ddd.postgres psql -U postgres go-ddd < $(find -r . -name '*up.sql')

for filename in ./*up.sql;
   do
     docker exec -i ddd.postgres psql -U postgres go-ddd < $filename
 done