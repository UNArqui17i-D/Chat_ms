CREATE DATABASE ChatsDB;

USE ChatsDB;

DROP TABLE Chats;

CREATE TABLE Chats
(
	id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	message text,
	userfrom int,
	userto int 
);

INSERT INTO Chats (message, userfrom, userto) VALUES ('Prueba de mensaje', 1, 2);