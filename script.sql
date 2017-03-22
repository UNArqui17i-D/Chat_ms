CREATE DATABASE ChatsDB;

USE ChatsDB;

DROP TABLE Messages;

CREATE TABLE Messages
(
	id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	message text,
	userfrom int,
	userto int 
);

INSERT INTO Messages (message, userfrom, userto) VALUES ('Prueba de mensaje', 1, 2);