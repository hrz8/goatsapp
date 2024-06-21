/*
  Warnings:

  - You are about to alter the column `client_id` on the `devices` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(50)`.
  - You are about to alter the column `name` on the `devices` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(50)`.
  - You are about to alter the column `phone_number` on the `devices` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(20)`.
  - You are about to alter the column `description` on the `devices` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(140)`.
  - You are about to alter the column `name` on the `projects` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(50)`.
  - You are about to alter the column `alias` on the `projects` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(20)`.
  - You are about to alter the column `description` on the `projects` table. The data in that column could be lost. The data in that column will be cast from `Text` to `VarChar(140)`.

*/
-- AlterTable
ALTER TABLE "devices" ALTER COLUMN "client_id" SET DATA TYPE VARCHAR(50),
ALTER COLUMN "name" SET DATA TYPE VARCHAR(50),
ALTER COLUMN "phone_number" SET DATA TYPE VARCHAR(20),
ALTER COLUMN "description" SET DATA TYPE VARCHAR(140);

-- AlterTable
ALTER TABLE "projects" ALTER COLUMN "name" SET DATA TYPE VARCHAR(50),
ALTER COLUMN "alias" SET DATA TYPE VARCHAR(20),
ALTER COLUMN "description" SET DATA TYPE VARCHAR(140);
