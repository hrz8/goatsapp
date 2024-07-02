/*
  Warnings:

  - You are about to drop the column `client_id` on the `devices` table. All the data in the column will be lost.
  - Added the required column `client_device_id` to the `devices` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "devices" DROP COLUMN "client_id",
ADD COLUMN     "client_device_id" VARCHAR(50) NOT NULL;
