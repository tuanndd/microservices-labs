package com.jaegartracingjavaservice.repo;

import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Data
@Entity
@Table(name = "salary_grade")
public class SalaryGrade {
    @Id
    private Long id;
    private String grade;
    private String title;
}
