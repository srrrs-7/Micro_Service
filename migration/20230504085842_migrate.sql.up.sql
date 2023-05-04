CREATE TABLE `Users` (
  `id` int PRIMARY KEY,
  `username` varchar(50),
  `password` varchar(255),
  `email` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `Calendar` (
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `date` date NOT NULL,
  `day` int NOT NULL,
  `month` int NOT NULL,
  `year` int NOT NULL,
  `holiday` varchar(50)
);

CREATE TABLE `Events` (
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `event` varchar(100) NOT NULL,
  `from_date` date NOT NULL,
  `to_date` date NOT NULL,
  `description` text,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `Notifications` (
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `event_id` int,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `WorkSpaces` (
  `id` int PRIMARY KEY,
  `work_space` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `Chats` (
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `work_space_id` int NOT NULL,
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

ALTER TABLE `Calendar` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Events` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Notifications` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Notifications` ADD FOREIGN KEY (`event_id`) REFERENCES `Events` (`id`);

ALTER TABLE `Chats` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Chats` ADD FOREIGN KEY (`work_space_id`) REFERENCES `WorkSpaces` (`id`);
