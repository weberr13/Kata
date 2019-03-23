(ns clojure-kata.seq)

;; Section 3 Sequential Collections
;; vectors
(assert (= [ 1 2 3 ] (vector 1 2 3)))
(assert (= [ 1 2 3 4 ] (conj [ 1 2 3 ] 4)))
(assert (= (get [ 1 2 3 ] 2) 3))
(assert (= (count [ 1 2 3 4 ]) 4))
(def v [1 2 3])
(assert (= (conj v 4 5 6) [1 2 3 4 5 6]))
(assert (= v [1 2 3]))
;; lists
(def things '("a" 2 :tag))
(assert (= (peek things) "a"))
(def things (conj things 0))
(assert (= (peek things) 0))
(def things (pop things))
(assert (= (peek things) "a"))