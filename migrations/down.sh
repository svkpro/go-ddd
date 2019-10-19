#bin/bash

docker exec -i ddd.postgres  createdb go-ddd -U postgres

for filename in . *up.sql;
   do
     docker exec -i ddd.postgres psql -U postgres go-ddd < $filename
 done