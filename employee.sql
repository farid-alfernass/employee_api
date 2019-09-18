-- phpMyAdmin SQL Dump
-- version 4.7.6
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Sep 18, 2019 at 09:15 PM
-- Server version: 5.7.15
-- PHP Version: 7.1.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `employee`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetAllEmployees` (IN `per_page` INT(11), IN `skip` INT(11))  NO SQL
BEGIN
   	SELECT *  FROM employees 
   	WHERE deleted_at IS NULL
   	ORDER BY created_at DESC
   	LIMIT per_page
   	OFFSET skip;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `GetEmployeeDetail` (IN `user` TEXT, IN `type` TEXT)  NO SQL
BEGIN
   	SELECT *
	FROM employees
   	WHERE deleted_at IS NULL AND CASE WHEN type = 'int' THEN id = user ELSE fullname = user END;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `GetEmployeesBySearch` (IN `keywords` TEXT, IN `per_page` INT, IN `skip` INT)  NO SQL
BEGIN
   	SELECT *  FROM employees 
   	WHERE deleted_at IS NULL
	AND fullname LIKE CONCAT ('%', keywords, '%')
   	ORDER BY id DESC
   	LIMIT per_page
   	OFFSET skip;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `employees`
--

CREATE TABLE `employees` (
  `id` int(11) NOT NULL,
  `email` varchar(50) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `employees`
--

INSERT INTO `employees` (`id`, `email`, `fullname`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'farid.wicak@gmail.com', 'farid tri wicaksono', '2019-09-18 19:32:13', '2019-09-18 12:32:13', NULL),
(2, 'putrafarid93@gmail.com', 'farid saputra', '2019-09-18 19:32:42', '2019-09-18 12:35:37', '2019-09-18 19:45:36'),
(3, 'putrafarid@gmail.com', 'farid si ganteng', '2019-09-18 19:44:18', '2019-09-18 12:44:18', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `employees`
--
ALTER TABLE `employees`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;


DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetAllEmployees`(IN `per_page` INT(11), IN `skip` INT(11))
    NO SQL
BEGIN
    SELECT *  FROM employees 
    WHERE deleted_at IS NULL
    ORDER BY created_at DESC
    LIMIT per_page
    OFFSET skip;
END$$
DELIMITER ;

DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetEmployeeDetail`(IN `user` TEXT, IN `type` TEXT)
    NO SQL
BEGIN
    SELECT *
  FROM employees
    WHERE deleted_at IS NULL AND CASE WHEN type = 'int' THEN id = user ELSE fullname = user END;
END$$
DELIMITER ;

DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetEmployeesBySearch`(IN `keywords` TEXT, IN `per_page` INT, IN `skip` INT)
    NO SQL
BEGIN
    SELECT *  FROM employees 
    WHERE deleted_at IS NULL
  AND fullname LIKE CONCAT ('%', keywords, '%')
    ORDER BY id DESC
    LIMIT per_page
    OFFSET skip;
END$$
DELIMITER ;

