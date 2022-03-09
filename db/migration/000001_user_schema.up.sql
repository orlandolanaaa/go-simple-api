CREATE TABLE `db_entry_task`.`users` (
                                         `id` INT NOT NULL AUTO_INCREMENT,
                                         `username` VARCHAR(32) NOT NULL,
                                         `password` VARCHAR(255) NOT NULL,
                                         `email` VARCHAR(255) NOT NULL,
                                         `nickname` VARCHAR(50) NULL,
                                         `profile_picture` VARCHAR(255) NULL,
                                         `created_at` DATETIME NULL,
                                         `updated_at` DATETIME NULL,
                                         `deleted_at` DATETIME NULL,
                                         UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
                                         PRIMARY KEY (`id`),
                                         UNIQUE INDEX `username_UNIQUE` (`username` ASC) INVISIBLE,
                                         UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);