DROP TABLE IF EXISTS employee;
CREATE TABLE employee (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(128) NOT NULL,
  dept       VARCHAR(128) NOT NULL,
  age        TINYINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO employee
  (name, dept, age)
VALUES
  ('杉下右京', '人事部', 60),
  ('亀山薫', '人事部', 40),
  ('神戸尊','開発部',38),
  ('甲斐享', '開発部', 35),
  ('冠城亘', '社長室', 31);