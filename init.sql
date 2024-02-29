CREATE TABLE IF NOT EXISTS payments (
                                        payment_id SERIAL PRIMARY KEY,
                                        merchant_id INT,
                                        customer_id INT,
                                        amount DECIMAL(10, 2),
                                        currency VARCHAR(3),
                                        status VARCHAR(10),
                                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS refunds (
                                       refund_id SERIAL PRIMARY KEY,
                                       payment_id INT REFERENCES payments(payment_id),
                                       amount DECIMAL(10, 2),
                                       status VARCHAR(10),
                                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS audit_trail (
                                           event_id SERIAL PRIMARY KEY,
                                           event_type VARCHAR(50),
                                           event_timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                           user_id INT,
                                           description TEXT
);

-- Inserts para la tabla payments
INSERT INTO payments (merchant_id, customer_id, amount, currency, status) VALUES
                                                                              (1, 100, 150.00, 'USD', 'completed'),
                                                                              (2, 101, 200.00, 'EUR', 'pending');

-- Inserts para la tabla refunds
INSERT INTO refunds (payment_id, amount, status) VALUES
                                                     (1, 15.00, 'processed'),
                                                     (2, 20.00, 'initiated');

-- Inserts para la tabla audit_trail
INSERT INTO audit_trail (event_type, user_id, description) VALUES
                                                               ('Payment', 100, 'Payment completed for customer 100'),
                                                               ('Refund', 101, 'Refund initiated for customer 101');
