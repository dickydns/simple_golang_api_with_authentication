-- up
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(255),
    status BOOLEAN DEFAULT true
);
