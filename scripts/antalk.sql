DROP DATABASE IF EXISTS antalk;
--#CREATE DATABASE antalk DEFAULT CHARSET SET utf8 COLLATE utf8_general_ci;
CREATE DATABASE antalk DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE antalk;

DROP TABLE IF EXISTS  `User`;
CREATE TABLE `User` (
    `userId` int(11) NOT NULL,
    `userName` varchar(4096) NOT NULL,
    `password` varchar(4096) NOT NULL,
    PRIMARY KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS `Msg`;
CREATE TABLE `Msg` (
	`msgId` int(11) NOT NULL, 
	`srcUid` varchar(4096) NOT NULL,
	`dstUid` varchar(4096) NOT NULL,
	`msgBody` varchar(4096) COLLATE utf8mb4_bin DEFAULT '',
	PRIMARY KEY (`msgId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
