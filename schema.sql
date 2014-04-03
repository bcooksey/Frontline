CREATE DATABASE frontline;
CREATE USER 'user'@'localhost' IDENTIFIED BY 'pass'; -- give a real user and pass
GRANT ALL ON frontline.* TO 'user'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE users(
    `id` int unsigned NOT NULL PRIMARY KEY,
    `email` varchar(255) NOT NULL,
    `pass` varchar(255) NOT NULL
)Engine=InnoDB;
