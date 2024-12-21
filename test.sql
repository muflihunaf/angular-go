/*
 Navicat Premium Data Transfer

 Source Server         : pgsql local
 Source Server Type    : PostgreSQL
 Source Server Version : 160006 (160006)
 Source Host           : localhost:5432
 Source Catalog        : test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160006 (160006)
 File Encoding         : 65001

 Date: 21/12/2024 11:44:25
*/


-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "users_id_seq";
CREATE SEQUENCE "users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default",
  "age" int8,
  "city" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO "users" ("id", "name", "age", "city") VALUES (28, 'CUT NYAK DIN', 19, 'MOJOKERTO'), (29, 'CUT NYAK DIN', 19, 'MOJOKERTO'), (30, 'CUT NYAK DIN', 19, 'BANDA ACEH'), (31, 'CUT NYAK DIN', 19, 'BANDA ACEH'), (32, 'CUT NYAK DIN', 19, 'BANDA ACEH'), (33, 'CUT NYAK DIN', 19, 'BANDA ACEH'), (34, 'CUT NYAK DIN', 19, 'BANDA ACEH'), (35, 'CUT NYAK DIN', 19, 'JAKARTA SELATAN'), (15, 'FIRMAN', 19, 'JAWA TIMUR'), (36, 'CUT NYAK DIN', 19, 'JAKARTA SELATAN');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "users_id_seq"
OWNED BY "users"."id";
SELECT setval('"users_id_seq"', 36, true);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
