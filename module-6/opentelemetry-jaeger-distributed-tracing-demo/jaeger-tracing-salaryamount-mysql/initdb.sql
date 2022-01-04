DROP TABLE IF EXISTS salary_amount;
CREATE TABLE IF NOT EXISTS salary_amount (
   id INT(10) NOT NULL AUTO_INCREMENT,
   grade VARCHAR(100),
   minimum INT(10),
   maximum INT(10),
   PRIMARY KEY (id)
);

INSERT INTO salary_amount(grade, minimum, maximum) VALUES ('A1', 20000, 29999);
INSERT INTO salary_amount(grade, minimum, maximum) VALUES ('A2', 30000, 39999);
INSERT INTO salary_amount(grade, minimum, maximum) VALUES ('A3', 40000, 49999);
INSERT INTO salary_amount(grade, minimum, maximum) VALUES ('A4', 50000, 59999);
INSERT INTO salary_amount(grade, minimum, maximum) VALUES ('A5', 60000, 69999);
