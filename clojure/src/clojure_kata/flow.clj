(ns clojure-kata.flow)


;; Section 5 flow control
(assert (= (if (even? 2) "even" "odd") "even"))
(assert (= (if (even? 1) "even" "odd") "odd"))
(assert (= (if (true? false) "impossible") nil))
(assert (= (if true :truthy :falsey) :truthy))
(assert (= (if (Object.) :truthy :falsey) :truthy))
(assert (= (if [] :truthy :falsey) :truthy))
(assert (= (if 0 :truthy :falsey) :truthy))
(assert (= (if false :truthy :falsey) :falsey))
(assert (= (if nil :truthy :falsey) :falsey))
(assert (not (if (even? 5)
  (do (println "even")
  true)
  (do (println "odd")
  false)
)))
(assert (when (even? 2) (println "even when") true))
(assert (= (let [x 5]
  (cond
      (< x 2) "less than 2"  
      (< x 10) "less than 10"
  )
) "less than 10" ))
(assert (= (let [x 11]
  (cond
      (< x 2) "less than 2"  
      (< x 10) "less than 10"
      :else "greater than or equal to 10" ; any keyword works, this is convention
  )
) "greater than or equal to 10" ))
;; case
(defn case-example [x]
    (case x
      5 "x is 5"
      10 "x is 10")
)
(assert (= (case-example 5) "x is 5"))
(assert (try 
    (do 
       (case-example 7) 
       false
    )
    (catch Exception e true)
    (finally (println "cleanup"))
))
(defn case-default [x]
  (case x 
      5 "x is 5"
      10 "x is 10"
      "neither")
)
(assert (= (case-default 6) "neither"))
;; dotimes (for range)
(def s [])
(dotimes [i 3]
    (def s (conj s i))
)
(assert (= s [0 1 2]))
;; doseq
(def s2 [])
(doseq [n (range 3)]
    (def s2 (conj s2 n))
)
(assert (= s2 [0 1 2]))
;; nested doseq
(def m {})
(doseq [k [:a :b] v (range 3)]
     (def m (assoc m (str k v) v))
)
(assert (= (m ":a1") 1))
;; for (list comprehension)
(assert (= (for [k [:a :b]
      v (range 3)]
      [k v]) '([:a 0] [:a 1] [:a 2] [:b 0] [:b 1] [:b 2])))
;; loop/recur
(def v [])
(assert (= (loop [i 0]
  (if (< i 10)
    (do 
      (def v (conj v i))
      (recur (inc i)) ;; must be the end of the "do"
    )
    v
  )
)) [0 1 2 3 4 5 6 7 8 9])
;; defn/recur
(def v [])
(defn buildvect [i]
  (if (< i 10)
    (do (def v (conj v i))
      (recur (inc i)) ;; must be the end of the "do"
    )
    v
  )
)
(assert (= (buildvect 3) [3 4 5 6 7 8 9]))
;; throw
(assert (= (try
  (do
     (println "foo")
     (throw (Exception. "something")) 
  )
  (catch Exception e (.getMessage e))
) "something"))
;; clojure exception data
(assert (= (try
  (throw (ex-info "a problem" {:detail 42}))
  (catch Exception e (:detail (ex-data e )))
) 42))
;; file writer
(with-open [f (clojure.java.io/writer "/tmp/foo")]
  (.write f "something")
)