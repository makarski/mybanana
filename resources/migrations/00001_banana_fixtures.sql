-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO `banana` (`id`, `name`)
VALUES
	(1, 'Chingan banana'),
	(2, 'Lacatan banana'),
	(3, 'Lady Finger banana '),
	(4, 'Pisang jari buaya'),
	(5, 'Señorita banana'),
	(6, 'Sinwobogi banana');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
TRUNCATE `banana`;
