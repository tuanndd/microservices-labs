package com.jaegartracingjavaservice.repo;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface SalaryGradeRepository extends JpaRepository<SalaryGrade, Long> {

    @Query("select distinct t.grade from SalaryGrade t where t.title = ?1")
    Optional<String> findGradeByTitle(String title);
}