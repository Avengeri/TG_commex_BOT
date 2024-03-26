CREATE TABLE t_users_auth
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(64)  NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    chatId        INTEGER NOT NULL
);


COMMENT
    ON COLUMN t_users_auth.username IS 'Nickname пользователя';
COMMENT
    ON COLUMN t_users_auth.password_hash IS 'Хеш пароля';
COMMENT
    ON COLUMN t_users_auth.chatId IS 'ID чата телеграм';


CREATE TABLE t_users_todo
(
    id   SERIAL PRIMARY KEY,
    age  INTEGER     NOT NULL,
    name VARCHAR(64) NOT NULL
);

COMMENT
    ON COLUMN t_users_todo.id IS 'ID';
COMMENT
    ON COLUMN t_users_todo.age IS 'Возраст';
COMMENT
    ON COLUMN t_users_todo.name IS 'Имя';

