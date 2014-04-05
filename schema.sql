CREATE DATABASE frontline;
CREATE USER 'user'@'localhost' IDENTIFIED BY 'pass'; -- give a real user and pass
GRANT ALL ON frontline.* TO 'user'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE users(
    `id` int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `email` varchar(255) NOT NULL,
    `pass` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL DEFAULT "Player"
)Engine=InnoDB;
