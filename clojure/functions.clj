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
; ( let [x (rand-int (* 2 Math/PI))]
;        (assert (= (sincos2 x) 1.0))
; )
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


