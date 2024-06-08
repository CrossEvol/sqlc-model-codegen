-- Disable foreign key constraints
PRAGMA foreign_keys = OFF;

-- Down migration to delete all records from TodosOnGroups table
DELETE FROM TodosOnGroups;

-- Down migration to delete all records from Group table
DELETE FROM [Group];

-- Down migration to delete all records from Detail table
DELETE FROM Detail;

-- Down migration to delete all records from Todo table
DELETE FROM Todo;

-- Down migration to delete all records from Category table
DELETE FROM Category;

-- Down migration to delete all records from Tag table
DELETE FROM Tag;

-- Down migration to delete all records from Post table
DELETE FROM Post;

-- Down migration to delete all records from User table
DELETE FROM User;

-- Enable foreign key constraints
PRAGMA foreign_keys = ON;
