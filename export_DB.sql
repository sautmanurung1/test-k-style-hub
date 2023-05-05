-- MySQL dump 10.13  Distrib 8.0.21, for Win64 (x86_64)
--
-- Host: localhost    Database: skincare
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `like_reviews`
--

DROP TABLE IF EXISTS `like_reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `like_reviews` (
  `id` varchar(191) NOT NULL,
  `review_id` varchar(255) DEFAULT NULL,
  `members_id` varchar(255) DEFAULT NULL,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_like_reviews_deleted_at` (`deleted_at`),
  KEY `fk_like_reviews_review_product` (`review_id`),
  KEY `fk_like_reviews_member` (`members_id`),
  CONSTRAINT `fk_like_reviews_member` FOREIGN KEY (`members_id`) REFERENCES `members` (`id`),
  CONSTRAINT `fk_like_reviews_review_product` FOREIGN KEY (`review_id`) REFERENCES `review_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `like_reviews`
--

LOCK TABLES `like_reviews` WRITE;
/*!40000 ALTER TABLE `like_reviews` DISABLE KEYS */;
INSERT INTO `like_reviews` VALUES ('e99ec257-f72d-4cb5-b674-513d8009aec6','8fbe02c8-9672-4635-bbbf-5d624a632ba3','b01299ac-0785-45db-a194-393f8e68ffa8',1683285626,1683285626,0);
/*!40000 ALTER TABLE `like_reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `members` (
  `id` varchar(255) NOT NULL,
  `username` longtext,
  `gender` longtext,
  `skin_type` longtext,
  `skin_color` longtext,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_members_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES ('b01299ac-0785-45db-a194-393f8e68ffa8','sautmanurung','pria','sawo matang','coklat',1683285531,1683285531,0);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` varchar(255) NOT NULL,
  `product_name` longtext,
  `price` bigint DEFAULT NULL,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES ('e033edf3-3c79-4cb6-97bd-f2452a20a813','Sabun Muka Sari Ayu',100000,1683285538,1683285538,0);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `review_products`
--

DROP TABLE IF EXISTS `review_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `review_products` (
  `id` varchar(255) NOT NULL,
  `product_id` varchar(255) DEFAULT NULL,
  `member_id` varchar(255) DEFAULT NULL,
  `desc_review` longtext,
  `created_at` bigint DEFAULT NULL,
  `updated_at` bigint DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_review_products_deleted_at` (`deleted_at`),
  KEY `fk_review_products_member` (`member_id`),
  KEY `fk_review_products_product` (`product_id`),
  CONSTRAINT `fk_review_products_member` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`),
  CONSTRAINT `fk_review_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `review_products`
--

LOCK TABLES `review_products` WRITE;
/*!40000 ALTER TABLE `review_products` DISABLE KEYS */;
INSERT INTO `review_products` VALUES ('8fbe02c8-9672-4635-bbbf-5d624a632ba3','e033edf3-3c79-4cb6-97bd-f2452a20a813','b01299ac-0785-45db-a194-393f8e68ffa8','Muka ku menjadi halus dan glowing banget terimakasih skincare ini',1683285589,1683285589,0);
/*!40000 ALTER TABLE `review_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'skincare'
--

--
-- Dumping routines for database 'skincare'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-05 19:45:35
