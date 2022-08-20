create table credit_cards (
    id VARCHAR(255) not null,
    name VARCHAR(255),
    cc_number VARCHAR(255),
    expiration_month INTEGER,
    expiration_year INTEGER,
    cvv integer,
    balance float,
    cc_limit float,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

create table transactions (
    id VARCHAR(255) not null,
    amount float,
    status VARCHAR(255),
    description VARCHAR(255),
    store VARCHAR(255),
    credit_card_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);