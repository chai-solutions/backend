DROP TABLE IF EXISTS flights;
DROP TABLE IF EXISTS airlines;
DROP TABLE IF EXISTS airports;

CREATE TABLE IF NOT EXISTS airports (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT UNIQUE NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('John F. Kennedy International Airport', 'JFK', 40.6413, -73.7781);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Los Angeles International Airport', 'LAX', 33.9416, -118.4085);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('O Hare International Airport', 'ORD', 41.9742, -87.9073);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Hartsfield-Jackson Atlanta International Airport', 'ATL', 33.6407, -84.4277);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Dallas/Fort Worth International Airport', 'DFW', 32.8998, -97.0403);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Denver International Airport', 'DEN', 39.8561, -104.6737);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('San Francisco International Airport', 'SFO', 37.7749, -122.4194);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Seattle-Tacoma International Airport', 'SEA', 47.6062, -122.3321);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('McCarran International Airport (Las Vegas)', 'LAS', 36.084, -115.1537);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Miami International Airport', 'MIA', 25.7959, -80.2871);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Charlotte Douglas International Airport', 'CLT', 35.214, -80.9431);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Phoenix Sky Harbor International Airport', 'PHX', 33.4342, -112.0116);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('George Bush Intercontinental Airport (Houston)', 'IAH', 29.9902, -95.3368);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Boston Logan International Airport', 'BOS', 42.3656, -71.0096);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Minneapolis-Saint Paul International Airport', 'MSP', 44.8848, -93.2223);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Detroit Metropolitan Airport', 'DTW', 42.2124, -83.3534);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Philadelphia International Airport', 'PHL', 39.8744, -75.2424);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('San Diego International Airport', 'SAN', 32.7338, -117.1933);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Tampa International Airport', 'TPA', 27.9755, -82.5332);
INSERT INTO airports (name, iata, latitude, longitude) VALUES ('Baltimore/Washington International Thurgood Marshall Airport', 'BWI', 39.1754, -76.6684);

CREATE TABLE IF NOT EXISTS airlines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    iata VARCHAR(5) NOT NULL
);
INSERT INTO airlines (name, iata) VALUES ('American Airlines', 'AA');
INSERT INTO airlines (name, iata) VALUES ('Delta Airlines', 'DL');
INSERT INTO airlines (name, iata) VALUES ('Frontier Airlines', 'F9');
INSERT INTO airlines (name, iata) VALUES ('Spirit Airlines', 'NK');
INSERT INTO airlines (name, iata) VALUES ('United Airlines', 'UA');
INSERT INTO airlines (name, iata) VALUES ('Southwest Airlines', 'WN');

CREATE TABLE IF NOT EXISTS flights (
    id SERIAL PRIMARY KEY,
    flight_number VARCHAR NOT NULL,
    airline INT NOT NULL,
    dep_airport INT NOT NULL,
    arr_airport INT NOT NULL,
    sched_dep_time TIMESTAMP NOT NULL,
    sched_arr_time TIMESTAMP NOT NULL,
    actual_dep_time TIMESTAMP NOT NULL,
    actual_arr_time TIMESTAMP NOT NULL,
    FOREIGN KEY (airline) REFERENCES airlines (id),
    FOREIGN KEY (dep_airport) REFERENCES airports (id),
    FOREIGN KEY (arr_airport) REFERENCES airports (id)
);
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN494', 6, 6, 5, '2024-11-23 14:51:26', '2024-11-23 18:51:26', '2024-11-23 15:03:26', '2024-11-23 19:18:26');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL423', 2, 20, 1, '2024-11-08 01:03:30', '2024-11-08 03:03:30', '2024-11-08 01:32:30', '2024-11-08 03:00:30');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9458', 3, 2, 16, '2024-11-11 11:42:20', '2024-11-11 13:42:20', '2024-11-11 11:34:20', '2024-11-11 13:56:20');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL269', 2, 14, 7, '2024-11-01 03:25:22', '2024-11-01 05:25:22', '2024-11-01 03:30:22', '2024-11-01 05:17:22');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9904', 3, 17, 4, '2024-11-13 13:12:32', '2024-11-13 16:12:32', '2024-11-13 13:17:32', '2024-11-13 16:27:32');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA888', 5, 12, 20, '2024-11-17 12:48:26', '2024-11-17 17:48:26', '2024-11-17 13:17:26', '2024-11-17 17:39:26');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA216', 5, 12, 3, '2024-11-26 17:49:42', '2024-11-26 20:49:42', '2024-11-26 17:48:42', '2024-11-26 20:59:42');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9297', 3, 5, 12, '2024-11-15 06:40:36', '2024-11-15 07:40:36', '2024-11-15 06:54:36', '2024-11-15 07:40:36');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL185', 2, 2, 14, '2024-11-13 05:28:03', '2024-11-13 10:28:03', '2024-11-13 05:55:03', '2024-11-13 10:27:03');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA493', 5, 9, 4, '2024-11-15 05:38:09', '2024-11-15 08:38:09', '2024-11-15 05:26:09', '2024-11-15 08:45:09');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL613', 2, 11, 17, '2024-11-23 06:46:45', '2024-11-23 07:46:45', '2024-11-23 06:34:45', '2024-11-23 07:34:45');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN796', 6, 7, 10, '2024-11-22 15:25:02', '2024-11-22 16:25:02', '2024-11-22 15:29:02', '2024-11-22 16:50:02');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK762', 4, 1, 16, '2024-11-04 15:11:37', '2024-11-04 17:11:37', '2024-11-04 15:10:37', '2024-11-04 17:07:37');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL642', 2, 13, 9, '2024-11-20 01:17:20', '2024-11-20 03:17:20', '2024-11-20 01:42:20', '2024-11-20 03:45:20');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL265', 2, 18, 1, '2024-11-06 11:49:06', '2024-11-06 13:49:06', '2024-11-06 11:42:06', '2024-11-06 13:35:06');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL983', 2, 2, 3, '2024-11-09 07:50:58', '2024-11-09 12:50:58', '2024-11-09 08:07:58', '2024-11-09 13:08:58');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN552', 6, 12, 6, '2024-11-19 23:25:51', '2024-11-20 04:25:51', '2024-11-19 23:54:51', '2024-11-20 04:24:51');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN987', 6, 20, 17, '2024-11-20 23:32:20', '2024-11-21 00:32:20', '2024-11-20 23:40:20', '2024-11-21 00:57:20');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA354', 5, 5, 1, '2024-11-17 19:07:14', '2024-11-17 20:07:14', '2024-11-17 19:15:14', '2024-11-17 20:00:14');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN397', 6, 8, 4, '2024-11-28 08:38:17', '2024-11-28 12:38:17', '2024-11-28 09:04:17', '2024-11-28 12:34:17');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA736', 5, 6, 16, '2024-11-16 00:21:04', '2024-11-16 01:21:04', '2024-11-16 00:40:04', '2024-11-16 01:45:04');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA876', 5, 2, 5, '2024-11-10 00:19:09', '2024-11-10 04:19:09', '2024-11-10 00:46:09', '2024-11-10 04:24:09');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK566', 4, 9, 15, '2024-11-26 23:24:42', '2024-11-27 03:24:42', '2024-11-26 23:12:42', '2024-11-27 03:14:42');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9216', 3, 18, 15, '2024-11-13 15:17:00', '2024-11-13 20:17:00', '2024-11-13 15:38:00', '2024-11-13 20:38:00');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA554', 1, 18, 12, '2024-11-23 19:00:07', '2024-11-23 20:00:07', '2024-11-23 18:59:07', '2024-11-23 19:59:07');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK596', 4, 1, 12, '2024-11-27 02:45:14', '2024-11-27 04:45:14', '2024-11-27 02:54:14', '2024-11-27 05:07:14');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA386', 1, 1, 9, '2024-11-29 05:44:57', '2024-11-29 10:44:57', '2024-11-29 06:10:57', '2024-11-29 10:39:57');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL432', 2, 18, 19, '2024-11-12 16:27:22', '2024-11-12 18:27:22', '2024-11-12 16:48:22', '2024-11-12 18:13:22');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN786', 6, 2, 1, '2024-11-28 15:08:58', '2024-11-28 20:08:58', '2024-11-28 15:10:58', '2024-11-28 20:03:58');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9420', 3, 6, 8, '2024-11-27 06:03:15', '2024-11-27 09:03:15', '2024-11-27 06:27:15', '2024-11-27 08:55:15');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9908', 3, 2, 8, '2024-11-12 05:05:44', '2024-11-12 07:05:44', '2024-11-12 05:28:44', '2024-11-12 07:22:44');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK814', 4, 8, 2, '2024-11-24 18:34:10', '2024-11-24 19:34:10', '2024-11-24 18:21:10', '2024-11-24 19:20:10');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN953', 6, 20, 3, '2024-11-25 22:35:00', '2024-11-26 03:35:00', '2024-11-25 22:50:00', '2024-11-26 03:53:00');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9582', 3, 3, 5, '2024-11-03 11:08:02', '2024-11-03 13:08:02', '2024-11-03 11:09:02', '2024-11-03 12:55:02');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN472', 6, 9, 5, '2024-11-01 13:50:29', '2024-11-01 15:50:29', '2024-11-01 14:11:29', '2024-11-01 16:12:29');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA164', 1, 2, 9, '2024-11-01 19:26:22', '2024-11-01 20:26:22', '2024-11-01 19:40:22', '2024-11-01 20:52:22');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA447', 1, 5, 10, '2024-11-02 07:30:39', '2024-11-02 08:30:39', '2024-11-02 07:40:39', '2024-11-02 08:18:39');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9778', 3, 1, 9, '2024-11-16 13:05:12', '2024-11-16 15:05:12', '2024-11-16 12:57:12', '2024-11-16 15:16:12');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK553', 4, 13, 12, '2024-11-16 16:13:22', '2024-11-16 17:13:22', '2024-11-16 16:18:22', '2024-11-16 17:13:22');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN689', 6, 19, 20, '2024-11-20 16:43:26', '2024-11-20 17:43:26', '2024-11-20 16:57:26', '2024-11-20 18:10:26');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN677', 6, 19, 13, '2024-11-15 06:26:17', '2024-11-15 11:26:17', '2024-11-15 06:12:17', '2024-11-15 11:39:17');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9684', 3, 4, 20, '2024-11-01 14:38:54', '2024-11-01 15:38:54', '2024-11-01 14:32:54', '2024-11-01 15:56:54');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA348', 1, 5, 13, '2024-11-17 16:52:23', '2024-11-17 20:52:23', '2024-11-17 16:58:23', '2024-11-17 21:03:23');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL725', 2, 3, 12, '2024-11-18 13:02:07', '2024-11-18 14:02:07', '2024-11-18 13:32:07', '2024-11-18 14:01:07');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9238', 3, 4, 13, '2024-11-08 12:00:19', '2024-11-08 16:00:19', '2024-11-08 12:18:19', '2024-11-08 15:51:19');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9307', 3, 20, 4, '2024-11-04 11:05:12', '2024-11-04 13:05:12', '2024-11-04 11:02:12', '2024-11-04 13:04:12');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA170', 1, 12, 3, '2024-11-21 09:29:00', '2024-11-21 11:29:00', '2024-11-21 09:25:00', '2024-11-21 11:52:00');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA697', 1, 14, 17, '2024-11-05 09:43:06', '2024-11-05 12:43:06', '2024-11-05 09:50:06', '2024-11-05 13:02:06');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA433', 1, 4, 11, '2024-11-09 14:44:23', '2024-11-09 15:44:23', '2024-11-09 14:54:23', '2024-11-09 15:33:23');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9700', 3, 7, 13, '2024-11-16 17:31:41', '2024-11-16 22:31:41', '2024-11-16 17:55:41', '2024-11-16 22:58:41');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN225', 6, 1, 12, '2024-11-04 17:13:53', '2024-11-04 18:13:53', '2024-11-04 16:59:53', '2024-11-04 18:25:53');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA320', 5, 17, 11, '2024-11-24 03:35:58', '2024-11-24 08:35:58', '2024-11-24 03:21:58', '2024-11-24 08:22:58');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA545', 1, 5, 7, '2024-11-25 11:25:11', '2024-11-25 13:25:11', '2024-11-25 11:51:11', '2024-11-25 13:12:11');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN326', 6, 3, 15, '2024-11-01 15:59:34', '2024-11-01 17:59:34', '2024-11-01 16:23:34', '2024-11-01 18:22:34');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL737', 2, 11, 3, '2024-11-16 16:38:21', '2024-11-16 19:38:21', '2024-11-16 16:37:21', '2024-11-16 19:58:21');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK682', 4, 3, 20, '2024-11-13 08:21:21', '2024-11-13 13:21:21', '2024-11-13 08:31:21', '2024-11-13 13:13:21');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA894', 1, 20, 10, '2024-11-09 21:32:49', '2024-11-10 02:32:49', '2024-11-09 21:46:49', '2024-11-10 02:24:49');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA498', 5, 4, 3, '2024-11-21 13:27:30', '2024-11-21 14:27:30', '2024-11-21 13:53:30', '2024-11-21 14:54:30');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9154', 3, 10, 16, '2024-11-02 18:53:01', '2024-11-02 22:53:01', '2024-11-02 18:38:01', '2024-11-02 23:05:01');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA521', 5, 20, 7, '2024-11-13 05:27:27', '2024-11-13 06:27:27', '2024-11-13 05:42:27', '2024-11-13 06:55:27');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL716', 2, 2, 13, '2024-11-25 23:12:26', '2024-11-26 03:12:26', '2024-11-25 23:21:26', '2024-11-26 03:27:26');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA506', 5, 1, 13, '2024-11-12 07:25:28', '2024-11-12 09:25:28', '2024-11-12 07:39:28', '2024-11-12 09:30:28');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA639', 1, 17, 20, '2024-11-19 12:41:10', '2024-11-19 16:41:10', '2024-11-19 12:49:10', '2024-11-19 16:56:10');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA694', 5, 17, 12, '2024-11-18 20:51:20', '2024-11-18 21:51:20', '2024-11-18 20:45:20', '2024-11-18 21:38:20');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK785', 4, 17, 13, '2024-11-04 02:18:27', '2024-11-04 06:18:27', '2024-11-04 02:04:27', '2024-11-04 06:31:27');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA650', 5, 10, 13, '2024-11-21 12:34:06', '2024-11-21 14:34:06', '2024-11-21 12:25:06', '2024-11-21 14:49:06');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL272', 2, 8, 11, '2024-11-09 01:05:42', '2024-11-09 06:05:42', '2024-11-09 01:13:42', '2024-11-09 06:09:42');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL643', 2, 4, 12, '2024-11-24 07:21:34', '2024-11-24 08:21:34', '2024-11-24 07:33:34', '2024-11-24 08:25:34');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL944', 2, 18, 19, '2024-11-01 04:52:19', '2024-11-01 06:52:19', '2024-11-01 04:44:19', '2024-11-01 06:38:19');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA155', 5, 11, 9, '2024-11-23 17:02:57', '2024-11-23 22:02:57', '2024-11-23 17:21:57', '2024-11-23 22:32:57');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK370', 4, 12, 10, '2024-11-28 19:51:59', '2024-11-29 00:51:59', '2024-11-28 19:37:59', '2024-11-29 01:00:59');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN840', 6, 11, 13, '2024-11-02 12:12:51', '2024-11-02 17:12:51', '2024-11-02 12:18:51', '2024-11-02 17:37:51');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA503', 5, 9, 2, '2024-11-10 08:21:08', '2024-11-10 12:21:08', '2024-11-10 08:13:08', '2024-11-10 12:21:08');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN436', 6, 17, 6, '2024-11-23 00:08:16', '2024-11-23 01:08:16', '2024-11-23 00:34:16', '2024-11-23 01:10:16');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9869', 3, 5, 18, '2024-11-03 00:54:00', '2024-11-03 04:54:00', '2024-11-03 01:14:00', '2024-11-03 05:15:00');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK466', 4, 6, 1, '2024-11-03 08:37:40', '2024-11-03 10:37:40', '2024-11-03 09:07:40', '2024-11-03 10:26:40');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK324', 4, 5, 14, '2024-11-21 07:02:18', '2024-11-21 10:02:18', '2024-11-21 06:51:18', '2024-11-21 10:30:18');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA446', 5, 5, 3, '2024-11-08 09:27:11', '2024-11-08 10:27:11', '2024-11-08 09:27:11', '2024-11-08 10:33:11');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9638', 3, 5, 15, '2024-11-24 21:05:39', '2024-11-25 01:05:39', '2024-11-24 21:17:39', '2024-11-25 00:58:39');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK648', 4, 13, 15, '2024-11-29 10:57:59', '2024-11-29 13:57:59', '2024-11-29 10:47:59', '2024-11-29 14:11:59');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA579', 5, 15, 14, '2024-11-19 09:29:12', '2024-11-19 13:29:12', '2024-11-19 09:55:12', '2024-11-19 13:39:12');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9198', 3, 7, 14, '2024-11-28 10:52:25', '2024-11-28 12:52:25', '2024-11-28 11:21:25', '2024-11-28 12:46:25');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA626', 1, 18, 12, '2024-11-06 07:42:38', '2024-11-06 11:42:38', '2024-11-06 08:04:38', '2024-11-06 11:39:38');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA162', 1, 6, 13, '2024-11-28 08:37:49', '2024-11-28 09:37:49', '2024-11-28 08:36:49', '2024-11-28 09:35:49');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA216', 5, 8, 11, '2024-11-21 22:35:12', '2024-11-21 23:35:12', '2024-11-21 22:28:12', '2024-11-21 23:42:12');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN901', 6, 19, 15, '2024-11-25 20:46:59', '2024-11-25 23:46:59', '2024-11-25 20:50:59', '2024-11-25 23:37:59');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA567', 1, 14, 5, '2024-11-24 06:22:15', '2024-11-24 09:22:15', '2024-11-24 06:34:15', '2024-11-24 09:19:15');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('F9290', 3, 19, 13, '2024-11-09 17:44:39', '2024-11-09 18:44:39', '2024-11-09 18:12:39', '2024-11-09 19:07:39');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN108', 6, 12, 15, '2024-11-04 12:05:52', '2024-11-04 16:05:52', '2024-11-04 12:19:52', '2024-11-04 16:01:52');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA273', 1, 11, 19, '2024-11-25 07:26:22', '2024-11-25 09:26:22', '2024-11-25 07:32:22', '2024-11-25 09:40:22');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('UA915', 5, 9, 1, '2024-11-11 14:48:25', '2024-11-11 19:48:25', '2024-11-11 14:58:25', '2024-11-11 20:08:25');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA969', 1, 3, 2, '2024-11-28 21:59:59', '2024-11-29 01:59:59', '2024-11-28 21:59:59', '2024-11-29 01:47:59');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK644', 4, 14, 5, '2024-11-07 08:15:16', '2024-11-07 09:15:16', '2024-11-07 08:07:16', '2024-11-07 09:09:16');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL867', 2, 16, 8, '2024-11-28 14:46:52', '2024-11-28 16:46:52', '2024-11-28 15:15:52', '2024-11-28 16:53:52');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('DL917', 2, 8, 10, '2024-11-18 22:27:45', '2024-11-19 01:27:45', '2024-11-18 22:22:45', '2024-11-19 01:51:45');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN808', 6, 10, 3, '2024-11-01 23:08:30', '2024-11-02 00:08:30', '2024-11-01 23:18:30', '2024-11-02 00:28:30');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('WN424', 6, 15, 11, '2024-11-23 19:38:19', '2024-11-23 23:38:19', '2024-11-23 19:55:19', '2024-11-24 00:01:19');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK415', 4, 6, 1, '2024-11-02 10:35:37', '2024-11-02 11:35:37', '2024-11-02 11:05:37', '2024-11-02 11:56:37');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('AA109', 1, 18, 5, '2024-11-12 18:10:02', '2024-11-12 20:10:02', '2024-11-12 17:55:02', '2024-11-12 20:32:02');
INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('NK152', 4, 5, 6, '2024-11-21 02:22:40', '2024-11-21 05:22:40', '2024-11-21 02:29:40', '2024-11-21 05:52:40');
