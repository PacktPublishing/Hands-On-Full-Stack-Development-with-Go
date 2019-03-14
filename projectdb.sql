-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.12 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             9.5.0.5196
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for gomusic
CREATE DATABASE IF NOT EXISTS `gomusic` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `gomusic`;

-- Dumping structure for table gomusic.customers
CREATE TABLE IF NOT EXISTS `customers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(50) NOT NULL DEFAULT '0',
  `lastname` varchar(50) NOT NULL DEFAULT '0',
  `email` varchar(100) NOT NULL DEFAULT '0',
  `pass` varchar(100) NOT NULL DEFAULT '0',
  `cc_customerid` varchar(50) NOT NULL DEFAULT '0',
  `loggedin` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gomusic.customers: ~4 rows (approximately)
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` (`id`, `firstname`, `lastname`, `email`, `pass`, `cc_customerid`, `loggedin`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Mal', 'Zein', 'mal.zein@email.com', '$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC', 'cus_EL08toK8pfDcom', 0, '2018-08-14 07:52:54', '2019-03-03 19:01:16', NULL),
	(2, 'River', 'Sam', 'river.sam@email.com', '$2a$10$mNbCLmfCAc0.4crDg3V3fe0iO1yr03aRfE7Rr3vdfKMGVnnzovCZq', '', 0, '2018-08-14 07:52:55', '2019-01-12 22:39:01', NULL),
	(3, 'Jayne', 'Ra', 'jayne.ra@email.com', '$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC', 'cus_EL4GpQmVjwvUUZ', 0, '2018-08-14 07:52:55', '2019-01-13 21:56:05', NULL),
	(19, 'John', 'Doe', 'john.doe@bla.com', '$2a$10$T4c8rmpbgKrUA0sIqtHCaO0g2XGWWxFY4IGWkkpVQOD/iuBrwKrZu', '', 0, '2019-01-13 08:43:44', '2019-01-13 15:12:25', NULL);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;

-- Dumping structure for table gomusic.orders
CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `purchase_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gomusic.orders: ~11 rows (approximately)
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` (`id`, `customer_id`, `product_id`, `price`, `purchase_date`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, 90, '2018-12-29 23:34:32', '2018-12-29 23:35:36', '2018-12-29 23:35:36', NULL),
	(2, 1, 2, 299, '2018-12-29 23:34:53', '2018-12-29 23:35:48', '2018-12-29 23:35:48', NULL),
	(3, 1, 3, 16000, '2018-12-29 23:35:05', '2018-12-29 23:35:57', '2018-12-29 23:35:57', NULL),
	(4, 2, 1, 95, '2018-12-29 23:36:18', '2018-12-29 23:36:18', '2018-12-29 23:36:18', NULL),
	(5, 2, 2, 299, '2018-12-29 23:36:39', '2018-12-29 23:36:39', '2018-12-29 23:36:39', NULL),
	(6, 2, 4, 205, '2018-12-29 23:37:01', '2018-12-29 23:38:13', '2018-12-29 23:38:13', NULL),
	(7, 3, 4, 210, '2018-12-29 23:37:28', '2018-12-29 23:38:19', '2018-12-29 23:38:19', NULL),
	(8, 3, 5, 200, '2018-12-29 23:37:41', '2018-12-29 23:38:28', '2018-12-29 23:38:28', NULL),
	(9, 3, 6, 1000, '2018-12-29 23:37:54', '2018-12-29 23:38:32', '2018-12-29 23:38:32', NULL),
	(10, 19, 6, 1000, '2018-12-29 23:37:54', '2019-01-13 00:44:55', '2019-01-13 00:44:55', NULL),
	(11, 1, 3, 17000, '0000-00-00 00:00:00', '2019-01-14 06:03:08', '2019-01-14 06:03:08', NULL);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;

-- Dumping structure for table gomusic.products
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image` varchar(100) DEFAULT NULL,
  `smallimg` varchar(100) DEFAULT NULL,
  `imgalt` varchar(50) DEFAULT NULL,
  `description` text,
  `productname` varchar(50) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `promotion` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table gomusic.products: ~6 rows (approximately)
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` (`id`, `image`, `smallimg`, `imgalt`, `description`, `productname`, `price`, `promotion`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'img/strings.png', 'img/img-small/strings.png', 'string', '', 'Strings', 100, NULL, '2018-08-14 07:54:19', '2019-01-11 00:28:40', NULL),
	(2, 'img/redguitar.jpeg', 'img/img-small/redguitar.jpeg', 'redg', '', 'Red Guitar', 299, 240, '2018-08-14 07:54:20', '2019-01-11 00:29:11', NULL),
	(3, 'img/drums.jpg', 'img/img-small/drums.jpg', 'drums', '', 'Drums', 17000, NULL, '2018-08-14 07:54:20', '2019-01-11 22:05:42', NULL),
	(4, 'img/flute.jpeg', 'img/img-small/flute.jpeg', 'flute', '', 'Flute', 210, 190, '2018-08-14 07:54:20', '2019-01-11 00:29:53', NULL),
	(5, 'img/blackguitar.jpeg', 'img/img-small/blackguitar.jpeg', 'Black guitar', '', 'Black Guitar', 200, NULL, '2018-08-14 07:54:20', '2019-01-11 00:30:12', NULL),
	(6, 'img/saxophone.jpeg', 'img/img-small/saxophone.jpeg', 'Saxophone', '', 'Saxophone', 1000, 980, '2018-08-14 07:54:20', '2019-01-11 00:30:35', NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
