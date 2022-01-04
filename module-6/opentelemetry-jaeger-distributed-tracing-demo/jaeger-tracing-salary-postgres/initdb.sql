DROP TABLE IF EXISTS salary_grade;
CREATE TABLE IF NOT EXISTS salary_grade (
   id SERIAL,
   grade varchar(100),
   title varchar(100)
);

INSERT INTO salary_grade(grade, title) VALUES ('A1', 'Apprentice Software Engineer');
INSERT INTO salary_grade(grade, title) VALUES ('A2', 'Software Engineer');
INSERT INTO salary_grade(grade, title) VALUES ('A3', 'Senior Software Engineer');
INSERT INTO salary_grade(grade, title) VALUES ('A4', 'Lead Software Engineer');
INSERT INTO salary_grade(grade, title) VALUES ('A5', 'Principal Software Engineer');
