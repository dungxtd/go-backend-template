INSERT INTO users (id, name, email, password, phone_number) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'John Doe', 'john@example.com', '$2a$10$h8moN5PaZ5pPRvzn6JHYkOkrvZKzIwXrD5XH5YU.f/YD9E4R/iEZy', '+1234567890'),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Jane Smith', 'jane@example.com', '$2a$10$h8moN5PaZ5pPRvzn6JHYkOkrvZKzIwXrD5XH5YU.f/YD9E4R/iEZy', '+1987654321');

INSERT INTO tasks (id, title, user_id) VALUES
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Complete project documentation', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'),
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Review pull requests', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'),
('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Design new features', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12'),
('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'Write unit tests', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12');