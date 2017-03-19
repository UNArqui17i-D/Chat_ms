CREATE DATABASE ChatsDB;

USE ChatsDB;

CREATE TABLE Chats
(
	id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
	message text,
	idUserFrom int,
	idUserTo int 
);

INSERT INTO Chats (message, idUserFrom, idUserTo) VALUES ('Prueba de mensaje', 1, 2);