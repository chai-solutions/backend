import random
from datetime import datetime, timedelta

print(
    """DROP TABLE IF EXISTS flights;
DROP TABLE IF EXISTS airlines;
DROP TABLE IF EXISTS airports;
"""
)

print(
    """CREATE TABLE IF NOT EXISTS airports (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT UNIQUE NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);"""
)

airports = [
    {"id": 1, "iata": "JFK", "name": "John F. Kennedy International Airport",
        "latitude": 40.6413, "longitude": -73.7781},
    {"id": 2, "iata": "LAX", "name": "Los Angeles International Airport",
        "latitude": 33.9416, "longitude": -118.4085},
    {"id": 3, "iata": "ORD", "name": "O Hare International Airport",
        "latitude": 41.9742, "longitude": -87.9073},
    {"id": 4, "iata": "ATL", "name": "Hartsfield-Jackson Atlanta International Airport",
        "latitude": 33.6407, "longitude": -84.4277},
    {"id": 5, "iata": "DFW", "name": "Dallas/Fort Worth International Airport",
        "latitude": 32.8998, "longitude": -97.0403},
    {"id": 6, "iata": "DEN", "name": "Denver International Airport",
        "latitude": 39.8561, "longitude": -104.6737},
    {"id": 7, "iata": "SFO", "name": "San Francisco International Airport",
        "latitude": 37.7749, "longitude": -122.4194},
    {"id": 8, "iata": "SEA", "name": "Seattle-Tacoma International Airport",
        "latitude": 47.6062, "longitude": -122.3321},
    {"id": 9, "iata": "LAS",
        "name": "McCarran International Airport (Las Vegas)", "latitude": 36.084, "longitude": -115.1537},
    {"id": 10, "iata": "MIA", "name": "Miami International Airport",
        "latitude": 25.7959, "longitude": -80.2871},
    {"id": 11, "iata": "CLT", "name": "Charlotte Douglas International Airport",
        "latitude": 35.214, "longitude": -80.9431},
    {"id": 12, "iata": "PHX", "name": "Phoenix Sky Harbor International Airport",
        "latitude": 33.4342, "longitude": -112.0116},
    {"id": 13, "iata": "IAH",
        "name": "George Bush Intercontinental Airport (Houston)", "latitude": 29.9902, "longitude": -95.3368},
    {"id": 14, "iata": "BOS", "name": "Boston Logan International Airport",
        "latitude": 42.3656, "longitude": -71.0096},
    {"id": 15, "iata": "MSP", "name": "Minneapolis-Saint Paul International Airport",
        "latitude": 44.8848, "longitude": -93.2223},
    {"id": 16, "iata": "DTW", "name": "Detroit Metropolitan Airport",
        "latitude": 42.2124, "longitude": -83.3534},
    {"id": 17, "iata": "PHL", "name": "Philadelphia International Airport",
        "latitude": 39.8744, "longitude": -75.2424},
    {"id": 18, "iata": "SAN", "name": "San Diego International Airport",
        "latitude": 32.7338, "longitude": -117.1933},
    {"id": 19, "iata": "TPA", "name": "Tampa International Airport",
        "latitude": 27.9755, "longitude": -82.5332},
    {"id": 20, "iata": "BWI", "name": "Baltimore/Washington International Thurgood Marshall Airport",
        "latitude": 39.1754, "longitude": -76.6684},
]

for airport in airports:
    name = airport["name"]
    iata = airport["iata"]
    latitude = airport["latitude"]
    longitude = airport["longitude"]
    print(f"INSERT INTO airports (name, iata, latitude, longitude) VALUES ('{
          name}', '{iata}', {latitude}, {longitude});")

print()

print(
    """CREATE TABLE IF NOT EXISTS airlines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    iata VARCHAR(5) NOT NULL
);"""
)

airlines = [
    {"id": 1, "iata": "AA", "name": "American Airlines"},
    {"id": 2, "iata": "DL", "name": "Delta Airlines"},
    {"id": 3, "iata": "F9", "name": "Frontier Airlines"},
    {"id": 4, "iata": "NK", "name": "Spirit Airlines"},
    {"id": 5, "iata": "UA", "name": "United Airlines"},
    {"id": 6, "iata": "WN", "name": "Southwest Airlines"},
]

for airline in airlines:
    name = airline["name"]
    iata = airline["iata"]
    print(f"INSERT INTO airlines (name, iata) VALUES ('{name}', '{iata}');")

print()

print(
    """CREATE TABLE IF NOT EXISTS flights (
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
);"""
)


def random_datetime(start, end):
    return start + timedelta(
        seconds=random.randint(0, int((end - start).total_seconds()))
    )


for i in range(1, 101):
    airline = random.choice(airlines)
    departure_airport = random.choice(airports)
    arrival_airport = random.choice(
        [airport for airport in airports if airport["id"] != departure_airport["id"]]
    )

    sched_dep_time = random_datetime(
        datetime(2024, 11, 1, 0, 0), datetime(2024, 11, 30, 0, 0)
    )
    sched_arr_time = sched_dep_time + timedelta(hours=random.randint(1, 5))
    actual_dep_time = sched_dep_time + \
        timedelta(minutes=random.randint(-15, 30))
    actual_arr_time = sched_arr_time + \
        timedelta(minutes=random.randint(-15, 30))
    flight_number = f"{airline['iata']}{random.randint(100, 999)}"
    sched_dep_time_str = sched_dep_time.strftime("%Y-%m-%d %H:%M:%S")
    sched_arr_time_str = sched_arr_time.strftime("%Y-%m-%d %H:%M:%S")
    actual_dep_time_str = actual_dep_time.strftime("%Y-%m-%d %H:%M:%S")
    actual_arr_time_str = actual_arr_time.strftime("%Y-%m-%d %H:%M:%S")

    print(
        f"INSERT INTO flights (flight_number, airline, dep_airport, arr_airport, sched_dep_time, sched_arr_time, actual_dep_time, actual_arr_time) VALUES ('{flight_number}', {
            airline['id']}, {departure_airport['id']}, {arrival_airport['id']}, '{sched_dep_time_str}', '{sched_arr_time_str}', '{actual_dep_time_str}', '{actual_arr_time_str}');"
    )
