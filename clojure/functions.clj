;; https://clojure.org/guides/learn/functions
(println (+ 1 2))
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