CREATE DATABASE IF NOT EXISTS gameshop;
USE gameshop;
DROP TABLE IF EXISTS game;
CREATE TABLE game (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(128) NOT NULL,
  `detail` varchar(128) NOT NULL,
  `genre` varchar(128) NOT NULL,
  `price` int(9) NOT NULL,
  `release_date` DATETIME NOT NULL,
  `cero` varchar(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO game (title, detail, genre, price, release_date, cero) VALUES ('スーパーメリオ', '姫を助ける配管工の話', 'アクション', 4800, '2020-9-15 23:59:59.999999', 'A');
INSERT INTO game (title, detail, genre, price, release_date, cero) VALUES ('ポケモヌ', 'ポケモヌを集めて対戦しよう', 'RPG', 4800, '2020-9-15 23:59:59.999999', 'A');