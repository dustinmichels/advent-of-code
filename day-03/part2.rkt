#lang racket

(require 2htdp/batch-io)
(require rackunit)

; recusrively split input list into groups of 3 strings
(define (group-input lst)
  (if (empty? lst)
      lst
      (cons (take lst 3)  (group-input (drop lst 3)))))

; find common items between two rucksacks
(define (find-common-chars r1 r2)
  (filter
   (lambda (char) (string-contains? r2 (~a char)))
   (if (string? r1) (string->list r1) r1)))

; find the single char common in a group of 3 rucksacks
(define (common-in-group group)
  (first
   (find-common-chars
    (find-common-chars (car group) (cadr group))
    (last group))))

; compute the proper char priority
(define (get-priority char)
  (let ([code (char->integer char)])
    (if (< code 91) (- code 38) (- code 96))))


; ----- test case -----

(define test-group
  (list "vJrwpWtwJgWrhcsFMMfFFhFp"
        "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
        "PmmdzqPrVvPwwTWBwg"))

(check-equal? (common-in-group test-group) #\r)


; ----- real deal -----

(define lines (read-lines "input.txt"))

(apply + (map
          (lambda (g)
            (get-priority (common-in-group g)))
          (group-input lines)))
