CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(100),
  `occupation` varchar(100),
  `email` varchar(50) UNIQUE,
  `password_hash` varchar(255),
  `avatar_filename` varchar(255),
  `role` varchar(20),
  `token` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp
);

CREATE TABLE `campaigns` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int,
  `name` varchar(255),
  `short_description` varchar(255),
  `description` text,
  `goal_amount` int,
  `current_amount` int,
  `backer_count` int,
  `perks` text,
  `slug` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp
);

CREATE TABLE `campaign_images` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `campaign_id` int,
  `image_filename` varchar(255),
  `is_primary` bool DEFAULT false,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp
);

CREATE TABLE `transactions` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int,
  `campaign_id` int,
  `amount` int,
  `status` varchar(50),
  `code` varchar(100),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp
);

ALTER TABLE `campaigns` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `campaign_images` ADD FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`);
