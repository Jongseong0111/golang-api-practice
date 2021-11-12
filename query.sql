-- name: GetUser :one
SELECT * FROM scm_account
WHERE user_id = ? LIMIT 1;

-- name: ListUser :many
SELECT user_name, user_account, user_email FROM scm_account
ORDER BY user_name;

-- name: CreateUser :execresult
INSERT INTO scm_account (
    user_name, user_account, user_email, user_password
) VALUES (
             ?, ?, ?, ?
         );

-- name: UpdateUser :exec
UPDATE scm_account set scm_account.user_name=? where scm_account.user_id=?;

-- name: DeleteUser :exec
DELETE FROM scm_account
WHERE user_id = ?;

-- name: CheckDuplicateAccount :many
SELECT user_account from scm_account where user_account=?;

-- name: CheckDuplicateEmail :many
SELECT user_email from scm_account where user_email=?;