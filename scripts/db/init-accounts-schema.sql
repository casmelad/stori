     CREATE DATABASE accountsdb;
     CREATE SCHEMA stori;
	 
	 CREATE TABLE stori.users (
        id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        name varchar(100), 
        email varchar(50)
    );
   
   
    CREATE TABLE stori.accounts(
        id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        number varchar(100), 
        user_id INT NOT NULL,
		  CONSTRAINT fk_user FOREIGN KEY(user_id)  REFERENCES stori.users(id)
    );
    
    
    CREATE TABLE stori.transactions(
        id INT PRIMARY KEY,
        account_id INT NOT NULL,
        date DATE,
        ammount DECIMAL,
        CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES stori.accounts(id)
    );
    --GO
    INSERT INTO stori.users ("name", "email") VALUES ('casmelad', 'casmelad@gmail.com');
    INSERT INTO stori.accounts ("number", "user_id") VALUES ('100765987777', 1);