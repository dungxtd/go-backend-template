-- Insert users with different roles (0: customer, 1: owner, 2: admin)
INSERT INTO users (id, name, email, password, phone_number, role) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'John Customer', 'john@example.com', '$2a$10$h8moN5PaZ5pPRvzn6JHYkOkrvZKzIwXrD5XH5YU.f/YD9E4R/iEZy', '+1234567890', 0),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Jane Owner', 'jane@example.com', '$2a$10$h8moN5PaZ5pPRvzn6JHYkOkrvZKzIwXrD5XH5YU.f/YD9E4R/iEZy', '+1987654321', 1),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Admin User', 'admin@example.com', '$2a$10$h8moN5PaZ5pPRvzn6JHYkOkrvZKzIwXrD5XH5YU.f/YD9E4R/iEZy', '+1122334455', 2);

-- Insert venues (sport_type: 0: football, 1: basketball, 2: tennis, 3: badminton)
INSERT INTO venues (id, owner_id, name, address, latitude, longitude, sport_type, price_per_hour) VALUES
(1, 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Downtown Football Arena', '123 Main St, City', 10.762622, 106.660172, 0, 100.00),
(2, 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Central Basketball Court', '456 Park Ave, City', 10.773831, 106.704893, 1, 80.00),
(3, 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Tennis Club', '789 Sports St, City', 10.780461, 106.687912, 2, 120.00);

-- Insert bookings (status: 0: pending, 1: confirmed, 2: canceled)
INSERT INTO bookings (id, user_id, venue_id, booking_date, start_time, end_time, status) VALUES
(1, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 1, '2024-01-20', '10:00', '12:00', 1),
(2, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 2, '2024-01-21', '14:00', '16:00', 0),
(3, 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 3, '2024-01-22', '09:00', '11:00', 2);

-- Insert saved venues
INSERT INTO saved_venues (user_id, venue_id) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 1),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 2),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 3);