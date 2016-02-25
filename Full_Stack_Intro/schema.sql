CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `author` varchar(32) NOT NULL,
  `content` varchar(500) NOT NULL,
    PRIMARY KEY (`id`)
);