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