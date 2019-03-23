(ns clojure-kata.hashed)

;; Section 4, hashed collections
;; sets
(def s1 #{ "a", "b", "c"})
(assert (contains? s1 "a"))
(assert (not (contains? s1 "d")))
(def s2 (conj s1 "d"))
(assert (contains? s2 "d"))
(def s3 (disj s2 "d" "a"))
(assert (not (contains? s3 "a")))
(assert (not (contains? s3 "d")))
(def s4 (conj (sorted-set) "d" "a" "c" "b"))
(assert (= s4 #{"a", "b", "c", "d"}))
(def s5 (conj s4 "0")) ;; sorted attribute retained
(assert (= s5 #{"0", "a", "b", "c", "d"}))
(def v1 ["a" "a" "b"])
(def s6 (into #{"a", "c"} v1))
(assert (= s6 #{"a", "c", "b"}))
;; maps
(def m1 {"a" 1 "b" 2})
(def m2 {"a" 1, "b" 2})
(assert (= m1 m2))
(def m3 (assoc m1 "c" 3))
(assert (= m3 {"a" 1 "b" 2 "c" 3}))
(def m4 (assoc m1 "a" 3))
(assert (= m4 {"a" 3 "b" 2}))
(def m5 (dissoc m3 "c"))
(assert (= m5 m1))
(assert (= (get m1 "a") 1))
(assert (= (m1 "a") 1))
(assert (= (m1 "d") nil))
(assert (= (m1 "not-found" :default) :default))
(assert (= (get m1 "also-not-found" :default) :default))
(assert (contains? m1 "a"))
(assert (= (find m1 "a") ["a" 1]))
(assert (= (keys m1) '("a" "b")))
(assert (= (vals m1) '(1 2)))
;; build with zipmap
(def zip (zipmap s1 (repeat 0)))
(assert (= zip {"a" 0 "b" 0 "c" 0}))
;; with map and into
(def zip2 (into {} (map (fn [si] [si 0]) s1)))
(assert (= zip zip2))
;; with reduce
(def zip3 (reduce (fn [m si]
          (assoc m si 0))
        {} ; initial value
        s1))
(assert (= zip zip3))
;; merge with default conflict resolution
(def m6a {"a" 1 "b" 2 "c" 3})
(def m6b {"c" 4 "d" 5})
(assert (= (merge m6a m6b) {"a" 1 "b" 2 "c" 4 "d" 5}))
(assert (= (merge-with + m6a m6b) {"a" 1 "b" 2 "c" (+ 4 3) "d" 5}))
;; ordered maps
(assert (= (sorted-map
         "Bravo" 204
         "Alfa" 35
         "Sigma" 99
         "Charlie" 100) 
{"Alfa" 35, "Bravo" 204, "Charlie" 100, "Sigma" 99}))
;; Field accessablity
(def person
  {:first-name "Kelly"
   :last-name "Keen"
   :age 32
   :occupation "Programmer"})
(assert (= (person :age) (:age person)))
(assert (= (:notfound person "default") "default"))
;; Nested
(def company
  {:name "WidgetCo"
   :address {:street "123 Main St"
             :city "Springfield"
             :state "IL"}})
(assert (= (get-in company [:address :city]) "Springfield"))
(def updated_company (assoc-in company [:address :street] "303 Broadway"))
(assert (= (get-in updated_company [:address :street]) "303 Broadway"))
;; records
;; Define a record structure
(defrecord Person [first-name last-name age occupation])
;; Positional constructor - generated
(def kelly (->Person "Kelly" "Keen" 32 "Programmer"))
;; Map constructor - generated
(def kelly2 (map->Person
             {:first-name "Kelly"
              :last-name "Keen"
              :age 32
              :occupation "Programmer"}))
(assert (= kelly kelly2))
(assert (= (:age kelly) 32))