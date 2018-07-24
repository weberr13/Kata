;; https://clojure.org/guides/learn/functions
(println (+ 1 2))
;; Section 2, functions
;; 1
(defn greet [] (println "hello"))
(greet)
(def greetfn (fn [] (println "hello")) )
(greetfn)
;; 2
(def greetpnd #(println "hello"))
(greetpnd)
;; 3
(defn greeting 
    ([] (greeting "Hello" "World"))
    ([x] (greeting "Hello" x))
    ([x, y] (str x ", " y "!"))
)
(assert (= "Hello, World!" (greeting)))
(assert (= "Hello, Clojure!" (greeting "Clojure")))
(assert (= "Good morning, Clojure!" (greeting "Good morning" "Clojure")))
;; 4
(defn do-nothing [x] x)
(require '[clojure.repl :refer :all])
(source identity)
(assert (= "foo" (do-nothing "foo")))
;; 5
(defn always-thing [& x] :thing)
(assert (= :thing (always-thing "bar")))
;; 6
(defn make-thingy [x] 
    (fn [& y] x)
)
(let [n (rand-int Integer/MAX_VALUE)
      f (make-thingy n)]
  (assert (= n (f)))
  (assert (= n (f :foo)))
  (assert (= n (apply f :foo (range)))))
  (source constantly)
;; 7
(defn triplicate [f] (f) (f) (f))
(triplicate #(println "foo"))
;; 8
(defn opposite [f]
  (fn [& args] (not (apply f args)))
)
(assert not ((opposite #(and %1 %2)) true true))
(assert ((opposite #(or %1 %2)) false false))
;; 9)
(defn triplicate2 [f & args]
  (triplicate #(apply f args))
)
(triplicate2 #(println (greeting %1 %2)) "foo" "bar")
;; 10
(println (Math/cos Math/PI))
(assert (= (Math/cos Math/PI) -1.0))
(defn sincos2 [x] 
    (+ (Math/pow (Math/sin x) 2) (Math/pow (Math/cos x) 2))
)
( let [x (rand-int (* 2 Math/PI))]
       (assert (= (sincos2 x) 1.0))
)
;; 11
(import java.net.URL)
(defn http-get [url]
  (let [
      u (URL. url)
  ]
    (println (slurp (.openStream u)))
  )
)
(http-get "https://google.com")
(defn fast-http-get [url] (slurp url))
(fast-http-get "https://google.com")
;; 12 currying
(defn one-less-arg [f x]
  (fn [& args] (apply f x args))
)
(let [
    f (one-less-arg greeting "robert")] 
    (assert (= (greeting "robert" "weber")) (f "weber"))
    (assert (= (greeting "robert")) (f))
)
(source partial)
;; 13 composition
(defn two-functions [f g]
    #(f (g %1))
)
(assert (= ((two-functions greeting (one-less-arg greeting "robert")) "weber") "Hello, robert, weber!!"))
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


