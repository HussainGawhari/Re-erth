
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

CREATE TABLE clients_history (
        id SERIAL PRIMARY KEY ,
        client_id INTEGER NOT NULL,
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


 CREATE OR REPLACE FUNCTION clients_history_trigger()
	RETURNS TRIGGER AS $$
	BEGIN
		IF (TG_OP = 'DELETE') THEN
			INSERT INTO clients_history (client_id, first_name, last_name, telephone, email, status, street, postal_code, city, country)
			VALUES (OLD.id, OLD.first_name, OLD.last_name, OLD.telephone, OLD.email, OLD.status, OLD.street, OLD.postal_code, OLD.city, OLD.country);
			RETURN OLD;
		ELSE
			INSERT INTO clients_history (client_id, first_name, last_name, telephone, email, status, street, postal_code, city, country)
			VALUES (NEW.id, NEW.first_name, NEW.last_name, NEW.telephone, NEW.email, NEW.status, NEW.street, NEW.postal_code, NEW.city, NEW.country);
			RETURN NEW;
		END IF;
	END;
	$$ LANGUAGE plpgsql;


    CREATE TRIGGER clients_history_trigger
      AFTER INSERT OR UPDATE OR DELETE ON clients
      FOR EACH ROW EXECUTE FUNCTION clients_history_trigger();
   