HOST="127.0.0.1"
PORT="3306"
DB_NAME="dummy_db"
USERNAME="root"
PASSWORD="password"
PROTOCOL="tcp"

echo "DROP DATABASE IF EXISTS ${DB_NAME};\n" >> temp.sql
echo "CREATE DATABASE IF NOT EXISTS ${DB_NAME};\n" >> temp.sql
echo "USE ${DB_NAME};\n" >> temp.sql
cat ./files/sql/schema.sql >> temp.sql
cat ./files/sql/seed.sql >> temp.sql

mysql -h ${HOST} -P ${PORT} -u ${USERNAME} --password=${PASSWORD} --protocol=${PROTOCOL} --verbose < temp.sql
rm temp.sql
