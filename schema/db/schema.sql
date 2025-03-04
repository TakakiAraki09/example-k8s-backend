-- Create "example_table" table
CREATE TABLE `example_table` (`id` int unsigned NOT NULL AUTO_INCREMENT, `title` varchar(100) NULL, `hoge` varchar(100) NULL, `fuga` varchar(100) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (`id` int unsigned NOT NULL AUTO_INCREMENT, `name` varchar(100) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "todos" table
CREATE TABLE `todos` (`id` int unsigned NOT NULL AUTO_INCREMENT, `text` varchar(100) NULL, `done` varchar(100) NULL, `user_id` int unsigned NOT NULL, PRIMARY KEY (`id`), INDEX `user_id` (`user_id`), CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
