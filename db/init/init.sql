CREATE DATABASE pomodork_db;
CREATE USER 'pomodork_root'@'%' identified by 'password';
CREATE USER 'pomodork_dev'@'%' identified by 'developer';

GRANT ALL ON *.* TO 'root'@'%' WITH GRANT OPTION;
GRANT ALL ON pomodork_db.* TO 'pomodork_root'@'%' WITH GRANT OPTION;
GRANT SELECT,INSERT,UPDATE,DELETE ON pomodork_db.* TO 'pomodork_dev'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
