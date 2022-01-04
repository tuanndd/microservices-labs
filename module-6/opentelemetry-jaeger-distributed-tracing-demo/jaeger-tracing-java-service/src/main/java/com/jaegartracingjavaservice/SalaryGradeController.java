package com.jaegartracingjavaservice;

import com.jaegartracingjavaservice.repo.SalaryGradeRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

@CrossOrigin
@RestController
@Slf4j
@RequiredArgsConstructor
public class SalaryGradeController {

    private final SalaryGradeRepository salaryGradeRepository;

    @GetMapping("/salary-grade/{title}")
    public ResponseEntity<Grade> getSalaryGrade(@PathVariable String title) {
        log.info("Receive Request to find grade for title {}", title);

        String grade = salaryGradeRepository.findGradeByTitle(title)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND));

        return new ResponseEntity<>(new Grade(grade), HttpStatus.OK);
    }
}
