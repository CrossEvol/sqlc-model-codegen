-- Insert records into User table
INSERT INTO User (name, password, email, email_verified, image, role)
VALUES ('Alice', 'password1', 'alice@example.com', '2023-06-01 10:00:00', 'alice.jpg', 'ADMIN'),
       ('Bob', 'password2', 'bob@example.com', '2023-06-02 11:00:00', 'bob.jpg', 'USER'),
       ('Charlie', 'password3', 'charlie@example.com', '2023-06-03 12:00:00', 'charlie.jpg', 'USER'),
       ('David', 'password4', 'david@example.com', '2023-06-04 13:00:00', 'david.jpg', 'GUEST'),
       ('Eve', 'password5', 'eve@example.com', '2023-06-05 14:00:00', 'eve.jpg', 'UNKNOWN'),
       ('Frank', 'password6', 'frank@example.com', '2023-06-06 15:00:00', 'frank.jpg', 'USER'),
       ('Grace', 'password7', 'grace@example.com', '2023-06-07 16:00:00', 'grace.jpg', 'ADMIN'),
       ('Hannah', 'password8', 'hannah@example.com', '2023-06-08 17:00:00', 'hannah.jpg', 'USER'),
       ('Ivy', 'password9', 'ivy@example.com', '2023-06-09 18:00:00', 'ivy.jpg', 'USER'),
       ('Jack', 'password10', 'jack@example.com', '2023-06-10 19:00:00', 'jack.jpg', 'GUEST');

-- Insert records into Post table
INSERT INTO Post (name, created_by_id)
VALUES ('Post 1', '1'),
       ('Post 2', '2'),
       ('Post 3', '3'),
       ('Post 4', '4'),
       ('Post 5', '5'),
       ('Post 6', '6'),
       ('Post 7', '7'),
       ('Post 8', '8'),
       ('Post 9', '9'),
       ('Post 10', '10');

-- Insert records into Tag table
INSERT INTO Tag (name, parent_id)
VALUES ('Tag 1', NULL),
       ('Tag 2', 1),
       ('Tag 3', 1),
       ('Tag 4', NULL),
       ('Tag 5', 4),
       ('Tag 6', NULL),
       ('Tag 7', 6),
       ('Tag 8', NULL),
       ('Tag 9', 8),
       ('Tag 10', 9);

-- Insert records into Category table
INSERT INTO Category (name, desc)
VALUES ('Category 1', 'Description 1'),
       ('Category 2', 'Description 2'),
       ('Category 3', 'Description 3'),
       ('Category 4', 'Description 4'),
       ('Category 5', 'Description 5'),
       ('Category 6', 'Description 6'),
       ('Category 7', 'Description 7'),
       ('Category 8', 'Description 8'),
       ('Category 9', 'Description 9'),
       ('Category 10', 'Description 10');

-- Insert records into Todo table
INSERT INTO Todo (title, score, amount, status, priority, tags, content, created_by, assignee_email, detail_id,
                  category_id)
VALUES ('Todo 1', 10, 100.0, 'PENDING', 'HIGH', '[1]', 'Content 1', 'ADMIN', 'assignee1@gmail.com', 1, 1),
       ('Todo 2', 20, 200.0, 'PAUSED', 'MEDIUM', '[2]', 'Content 2', 'USER', 'assignee2@gmail.com', 2, 2),
       ('Todo 3', 30, 300.0, 'COMPLETED', 'LOW', '[3]', 'Content 3', 'ANONYM', 'assignee3@gmail.com', 3, 3),
       ('Todo 4', 40, 400.0, 'PROCESSING', 'HIGH', '[4]', 'Content 4', 'ADMIN', 'assignee4@gmail.com', 4, 4),
       ('Todo 5', 50, 500.0, 'PENDING', 'MEDIUM', '[5]', 'Content 5', 'USER', 'assignee5@gmail.com', 5, 5),
       ('Todo 6', 60, 600.0, 'PAUSED', 'LOW', '[6]', 'Content 6', 'ANONYM', 'assignee6@gmail.com', 6, 6),
       ('Todo 7', 70, 700.0, 'COMPLETED', 'HIGH', '[7]', 'Content 7', 'ADMIN', 'assignee7@gmail.com', 7, 7),
       ('Todo 8', 80, 800.0, 'PROCESSING', 'MEDIUM', '[8]', 'Content 8', 'USER', 'assignee8@gmail.com', 8, 8),
       ('Todo 9', 90, 900.0, 'PENDING', 'LOW', '[9]', 'Content 9', 'ANONYM', 'assignee9@gmail.com', 9, 9),
       ('Todo 10', 100, 1000.0, 'PAUSED', 'HIGH', '[10]', 'Content 10', 'ADMIN', 'assignee10@gmail.com', 10, 10);

-- Insert records into Detail table
INSERT INTO Detail (desc, img_url, todo_id)
VALUES ('Detail 1', 'detail1.jpg', 1),
       ('Detail 2', 'detail2.jpg', 2),
       ('Detail 3', 'detail3.jpg', 3),
       ('Detail 4', 'detail4.jpg', 4),
       ('Detail 5', 'detail5.jpg', 5),
       ('Detail 6', 'detail6.jpg', 6),
       ('Detail 7', 'detail7.jpg', 7),
       ('Detail 8', 'detail8.jpg', 8),
       ('Detail 9', 'detail9.jpg', 9),
       ('Detail 10', 'detail10.jpg', 10);

-- Insert records into Group table
INSERT INTO [Group] (name, desc)
VALUES ('Group 1', 'Description 1'),
       ('Group 2', 'Description 2'),
       ('Group 3', 'Description 3'),
       ('Group 4', 'Description 4'),
       ('Group 5', 'Description 5'),
       ('Group 6', 'Description 6'),
       ('Group 7', 'Description 7'),
       ('Group 8', 'Description 8'),
       ('Group 9', 'Description 9'),
       ('Group 10', 'Description 10');

-- Insert records into TodosOnGroups table
INSERT INTO TodosOnGroups (todo_id, group_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 5),
       (6, 6),
       (7, 7),
       (8, 8),
       (9, 9),
       (10, 10);
