#lang racket

(require 2htdp/batch-io)

; divide rucksack into its two compartments
;   string -> '('char . 'char)
(define (div-rucksack r)
  (let ([mid (/ (string-length r) 2)])
    (cons
     (string->list (substring r 0 mid))
     (string->list (substring r mid)))))

; find common item between two compartments
;   '('char . 'char) -> char
(define (find-common-item compartments)
  (car
   (filter
    (lambda (char) (list? (member char (cdr compartments))))
    (car compartments))))

; compute the proper char code
(define (get-priority char)
  (let ([code (char->integer char)])
    (if (< code 91) (- code 38) (- code 96))))


; --- Test cases ---
(get-priority (find-common-item (div-rucksack "vJrwpWtwJgWrhcsFMMfFFhFp")))
(get-priority (find-common-item (div-rucksack "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")))
(get-priority (find-common-item (div-rucksack "PmmdzqPrVvPwwTWBwg")))
(get-priority (find-common-item (div-rucksack "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")))
(get-priority (find-common-item (div-rucksack "ttgJtRGJQctTZtZT")))
(get-priority (find-common-item (div-rucksack "CrZsJsPPZsGzwwsLwLmpwMDw")))


(apply + (map (lambda (i)
                (get-priority (find-common-item (div-rucksack i))))
              (read-lines "input.txt")))
