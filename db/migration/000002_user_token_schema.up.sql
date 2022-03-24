CREATE TABLE `db_entry_task`.`user_tokens` (
                                               `id` INT NOT NULL AUTO_INCREMENT,
                                               `user_id` INT NULL,
                                               `token` VARCHAR(255) NULL,
                                               `expired_at` DATETIME NULL,
                                               `created_at` DATETIME NULL,
                                               `updated_at` DATETIME NULL,
                                               `deleted_at` DATETIME NULL,
                                               UNIQUE INDEX `id_UNIQUE` (`id` ASC),
                                               CONSTRAINT `user_id_fk_users`
                                                   FOREIGN KEY (`user_id`)
                                                       REFERENCES `db_entry_task`.`users` (`id`)
                                                       ON DELETE NO ACTION
                                                       ON UPDATE NO ACTION);