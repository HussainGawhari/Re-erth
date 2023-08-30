
     CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY ,
            name VARCHAR(25) NOT NULL,
            email_id VARCHAR(55) NOT NULL,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(20),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    
    CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY ,
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    telephone VARCHAR(20) NOT NULL,
    email VARCHAR(55) NOT NULL,
    status BOOLEAN DEFAULT true,
    street VARCHAR(55) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    city VARCHAR(55) NOT NULL,
    country VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


