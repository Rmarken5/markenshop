package repository

// language=postgresql
const insertPush = `INSERT INTO push (before_commit, after_commit, pusher_id, sender_id)
VALUES (:before_commit, :after_commit, :pusher_id, :sender_id)
RETURNING id;`

const insertSender = `INSERT INTO sender (login, html_url, avatar_url, url) VALUES (:login, :html_url, :avatar_url, :url ) RETURNING id;`

const insertPusher = `INSERT INTO pusher (name, email) VALUES (:name, :email) returning id;`

const getSenderByName = `SELECT id, login, avatar_url, html_url, url  FROM sender WHERE login = $1;`

const getPusherByName = `select id, name, email from pusher where name = $1;`
