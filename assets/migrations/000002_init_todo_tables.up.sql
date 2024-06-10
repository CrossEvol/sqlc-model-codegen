-- Create User table
CREATE TABLE User
(
    id            TEXT PRIMARY KEY,
    name          TEXT,
    password      TEXT,
    email         TEXT UNIQUE,
    email_verified DATETIME,
    image         TEXT,
    role          TEXT CHECK ( role in ('ADMIN', 'USER', 'GUEST', 'UNKNOWN') ) NOT NULL DEFAULT 'USER'
);

-- Create Post table
CREATE TABLE Post
(
    id            INTEGER PRIMARY KEY,
    name          TEXT NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by_id TEXT NOT NULL,
    FOREIGN KEY (created_by_id) REFERENCES User (id)
);

CREATE INDEX idx_post_name ON Post (name);

-- Create Tag table
CREATE TABLE Tag
(
    id        INTEGER PRIMARY KEY,
    name      TEXT NOT NULL,
    parent_id INTEGER,
    FOREIGN KEY (parent_id) REFERENCES Tag (id) ON DELETE CASCADE ON UPDATE NO ACTION
);

-- Create Todo table
CREATE TABLE Todo
(
    id             INTEGER PRIMARY KEY,
    title          TEXT                                                                     NOT NULL DEFAULT '',
    score          INTEGER                                                                  NOT NULL DEFAULT 0,
    amount         REAL                                                                     NOT NULL DEFAULT 0.0,
    status         TEXT CHECK ( status in ('PENDING', 'PAUSED', 'COMPLETED', 'PROCESSING')) NOT NULL DEFAULT 'PENDING',
    created_at     DATETIME                                                                          DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME                                                                          DEFAULT CURRENT_TIMESTAMP,
    deadline       DATETIME                                                                          DEFAULT CURRENT_TIMESTAMP,
    priority       TEXT CHECK ( priority in ('HIGH', 'MEDIUM', 'LOW'))                      NOT NULL DEFAULT 'LOW',
    tags           TEXT                                                                     NOT NULL DEFAULT '[]',
    content        TEXT                                                                     NOT NULL DEFAULT '',
    created_by     TEXT CHECK ( created_by in ('ADMIN', 'USER', 'ANONYM'))                  NOT NULL DEFAULT 'ANONYM',
    assignee_email TEXT                                                                     NOT NULL DEFAULT 'assignee@gmail.com',
    detail_id      INTEGER,
    category_id    INTEGER                                                                           DEFAULT 1,
    FOREIGN KEY (category_id) REFERENCES Category (id) ON DELETE NO ACTION ON UPDATE CASCADE
);

-- Create TodoDetail table
CREATE TABLE Detail
(
    id      INTEGER PRIMARY KEY,
    desc    TEXT           NOT NULL DEFAULT '',
    img_url TEXT           NOT NULL DEFAULT '',
    todo_id INTEGER UNIQUE NOT NULL,
    FOREIGN KEY (todo_id) REFERENCES Todo (id) ON DELETE CASCADE ON UPDATE NO ACTION
);

-- Create Category table
CREATE TABLE Category
(
    id         INTEGER PRIMARY KEY,
    name       TEXT NOT NULL DEFAULT '',
    desc       TEXT NOT NULL DEFAULT '',
    created_at DATETIME      DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME      DEFAULT CURRENT_TIMESTAMP
);

-- Create Group table
CREATE TABLE [Group]
(
    id         INTEGER PRIMARY KEY,
    name       TEXT NOT NULL DEFAULT '',
    desc       TEXT NOT NULL DEFAULT '',
    created_at DATETIME      DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME      DEFAULT CURRENT_TIMESTAMP
);

-- Create TodosOnGroups table
CREATE TABLE TodosOnGroups
(
    todo_id     INTEGER NOT NULL,
    group_id    INTEGER NOT NULL,
    assigned_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (todo_id, group_id),
    FOREIGN KEY (todo_id) REFERENCES Todo (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (group_id) REFERENCES [Group] (id) ON DELETE CASCADE ON UPDATE CASCADE
);
