-- Delete data in reverse order of dependencies
DELETE FROM saved_venues;
DELETE FROM bookings;
DELETE FROM venues;
DELETE FROM users;