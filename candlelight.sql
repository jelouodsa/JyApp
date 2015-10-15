-- phpMyAdmin SQL Dump
-- version 4.0.10deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Oct 11, 2015 at 07:51 PM
-- Server version: 5.5.44-0ubuntu0.14.04.1
-- PHP Version: 5.5.9-1ubuntu4.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `candlelight`
--

-- --------------------------------------------------------

--
-- Table structure for table `communities`
--

CREATE TABLE IF NOT EXISTS `communities` (
  `name` varchar(40) NOT NULL,
  `admin` varchar(25) NOT NULL,
  `privacy` tinyint(1) NOT NULL,
  `country` varchar(40) NOT NULL,
  `state` varchar(40) NOT NULL,
  `city` varchar(40) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=2 ;

--
-- Dumping data for table `communities`
--

INSERT INTO `communities` (`name`, `admin`, `privacy`, `country`, `state`, `city`, `id`) VALUES
('djbc', 'melvinodsa', 0, 'Belarus', 'Brestskaya (Brest)', 'cjcjf', 1);

-- --------------------------------------------------------

--
-- Table structure for table `communitymember`
--

CREATE TABLE IF NOT EXISTS `communitymember` (
  `id` int(11) NOT NULL,
  `privilage` int(11) NOT NULL,
  `username` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `communitymember`
--

INSERT INTO `communitymember` (`id`, `privilage`, `username`) VALUES
(1, 0, 'melvinodsa');

-- --------------------------------------------------------

--
-- Table structure for table `userlist`
--

CREATE TABLE IF NOT EXISTS `userlist` (
  `username` varchar(25) NOT NULL,
  `email` varchar(40) NOT NULL,
  `password` varchar(25) NOT NULL,
  PRIMARY KEY (`username`,`email`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `userlist`
--

INSERT INTO `userlist` (`username`, `email`, `password`) VALUES
('jjj', 'jj@jj.com', 'jjj'),
('melvinodsa', 'melvinodsa@gmail.com', 'pass');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
