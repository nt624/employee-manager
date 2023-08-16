DROP TABLE IF EXISTS employee;
CREATE TABLE employee (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(128) NOT NULL,
  dept       VARCHAR(128) NOT NULL,
  age        TINYINT UNSIGNED NOT NULL,
  email      VARCHAR(128) NOT NULL,
  tel        VARCHAR(128) NULL,
  phone      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO employee
  (name, dept, age, email, tel, phone)
VALUES
  ('杉下右京', '人事部', 60, 'ukyo.sugishita@whi.com', '0312345678', '09012345678'),
  ('亀山薫', '人事部', 48, 'kaoru.kameyama@whi.com', '0451234567', '08087654321'),
  ('神戸尊','営業部',35, 'takeru.kanbe@whi.com', '', '08012345678'),
  ('甲斐享', '開発部', 28, 'toru.kai@whi.com', '', '07012345678'),
  ('冠城亘', '社長室', 42, 'wataru.kaburagi@whi.com', '0441234567', '09087654321');