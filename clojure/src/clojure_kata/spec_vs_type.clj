(ns clojure-kata.spec-vs-type
  (:require [clojure.spec.alpha :as spec]
            [clojure.tools.logging :as log]
            [clj-time.core :as t]
            [clj-time.format :as f]))

#_(f/show-formatters)
(def bday (f/parse (f/formatters :date-time-no-ms) "1979-03-13T03:13:00Z"))

(defn bday?
  [d]
  (let [d (f/parse (f/formatters :date-time-no-ms) d)]
    (and (= (t/month d) (t/month bday))
         (= (t/day d) (t/day bday)))))

(assert (bday? "2019-03-13T00:00:00Z"))
(assert (not (bday? "2019-03-15T00:00:00Z")))

(spec/def ::day string?)

(defn bday?
  [d]
  (when-not (spec/valid? ::day d)
    (throw (ex-info "not a string" {:d d})))
  (let [d (f/parse (f/formatters :date-time-no-ms) d)]
    (and (= (t/month d) (t/month bday))
         (= (t/day d) (t/day bday)))))

(defmacro throws?
  ([f]
   `(try
      ~f
      false
      (catch Exception _
        true)))
  ([f msg]
   `(try
      ~f
      false
      (catch Exception e#
        (if (= ~msg (ex-message e#))
          true
          (throw e#))))))

(assert (throws? (bday? [:a]) "not a string"))
(assert (throws? (bday? {:a "b"}) "not a string"))
(assert (throws? (bday? "foo") "not a string"))