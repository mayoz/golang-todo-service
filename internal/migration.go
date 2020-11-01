package internal

func Migration() string {
	return `
-- BEGIN
CREATE TABLE IF NOT EXISTS todos (
    id         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    text       VARCHAR(255) NOT NULL,
    completed  TINYINT      NOT NULL DEFAULT 0,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)  ENGINE=INNODB;
-- END
	`
}
