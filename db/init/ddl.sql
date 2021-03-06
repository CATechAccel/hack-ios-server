-- MySQL Script generated by MySQL Workbench
-- Fri Feb 14 23:09:20 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,ALLOW_INVALID_DATES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
-- -----------------------------------------------------
-- Table `hack_ios_server_api`.`users`
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS  `hack_ios_server_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `hack_ios_server_api`;

SET CHARSET utf8mb4;

CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '削除日時',
    `name` VARCHAR(64) UNIQUE NOT NULL COMMENT 'ユーザ名',
    `password` VARCHAR(128) NOT NULL COMMENT 'パスワード',
    PRIMARY KEY (`id`),
    INDEX `idx_created_at` (`created_at` ASC))
    ENGINE = InnoDB
    COMMENT = 'ユーザ';


-- -----------------------------------------------------
-- Table `hack_ios_server_api`.`tasks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tasks` (
    `id` VARCHAR(128) NOT NULL COMMENT 'タスクID',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL  COMMENT '削除日時',
    `name` VARCHAR(64) NOT NULL COMMENT 'タスク名',
    `description` VARCHAR(128) COMMENT 'タスクの詳細',
    `is_done`  TINYINT(1) NOT NULL COMMENT 'タスクの状態',
    `user_id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
    PRIMARY KEY (`id`),
    INDEX `idx_created_at` (`created_at` ASC),
    FOREIGN KEY fk_tasks_user_id(user_id) REFERENCES users(id))
    ENGINE = InnoDB
    COMMENT = 'タスク';

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;