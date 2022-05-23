-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 23 Bulan Mei 2022 pada 14.17
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `art-museum`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `artists`
--

CREATE TABLE `artists` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `nationality` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `birth_year` longtext DEFAULT NULL,
  `death_year` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `artists`
--

INSERT INTO `artists` (`id`, `name`, `nationality`, `description`, `birth_year`, `death_year`, `created_at`, `update_at`) VALUES
(9, 'Albert Besnard', 'French', 'Albert Besnard (French, 1849-1934)', '1898', '1934', '2022-05-22 21:58:33.295', '0000-00-00 00:00:00.000');

-- --------------------------------------------------------

--
-- Struktur dari tabel `arts`
--

CREATE TABLE `arts` (
  `id` bigint(20) NOT NULL,
  `name_art` longtext DEFAULT NULL,
  `name_artist` longtext DEFAULT NULL,
  `type` longtext DEFAULT NULL,
  `location` longtext DEFAULT NULL,
  `creation_date` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `arts`
--

INSERT INTO `arts` (`id`, `name_art`, `name_artist`, `type`, `location`, `creation_date`, `created_at`, `update_at`) VALUES
(2, 'Nathaniel Hurd', 'John Singleton Copley', 'Painting', 'The Cleveland Museum of Art', '1765', '2022-05-22 21:53:02.134', '0000-00-00 00:00:00.000');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `artists`
--
ALTER TABLE `artists`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `arts`
--
ALTER TABLE `arts`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `artists`
--
ALTER TABLE `artists`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `arts`
--
ALTER TABLE `arts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
