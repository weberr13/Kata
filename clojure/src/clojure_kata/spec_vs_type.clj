(ns clojure-kata.spec-vs-type
  (:require [clojure.spec.alpha :as spec]
            [clojure.string :as string]
            [clojure.tools.logging :as log]
            [clj-time.core :as t]
            [clj-time.format :as f]
            [clj-time.coerce :as c]
            [clojure.spec.gen.alpha :as gen]
            [clojure.spec.test.alpha :as stest]))

(def bday (f/parse (f/formatters :date-time-no-ms) "1979-03-15T03:13:00Z"))

(defn bday?
  [d]
  (let [d (f/parse (f/formatters :date-time-no-ms) d)]
    (and (= (t/month d) (t/month bday))
         (= (t/day d) (t/day bday)))))

(assert (bday? "2019-03-15T00:00:00Z"))
(assert (not (bday? "2019-03-13T00:00:00Z")))

(comment
  (assert (bday? 1)))

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
      (catch Exception _#
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
(assert (bday? "2019-03-15T00:00:00Z"))
(assert (not (bday? "2019-03-13T00:00:00Z")))
(comment
  (assert (throws? (bday? "foo") "not a string")))


(spec/def ::parseable-day
  #(not
     (throws?
       (f/parse (f/formatters :date-time-no-ms) %))))

(defn bday?
  [d]
  (when-not (spec/valid? ::parseable-day d)
    (throw (ex-info "not a day i can parse" {:d d})))
  (let [d (f/parse (f/formatters :date-time-no-ms) d)]
    (and (= (t/month d) (t/month bday))
         (= (t/day d) (t/day bday)))))

(assert (throws? (bday? [:a]) "not a day i can parse"))
(assert (throws? (bday? {:a "b"}) "not a day i can parse"))
(assert (throws? (bday? "foo") "not a day i can parse"))
(assert (bday? "2019-03-15T00:00:00Z"))
(assert (not (bday? "2019-03-14T00:00:00Z")))

(gen/sample (spec/gen ::day))
(comment
  (gen/sample (spec/gen ::parseable-day)))

(spec/def ::instant-str (spec/inst-in #inst "1900" #inst "2100"))
(def generate-day (gen/fmap
                    #(f/unparse (f/formatters :date-time-no-ms)
                                (c/from-long (.toEpochMilli
                                               (.toInstant ^java.util.Date %))))
                    (spec/gen ::instant-str)))

(drop 50 (gen/sample generate-day 55))

;; to illustrate when the generator tries my bday
(comment
  (spec/fdef bday?
             :args (spec/cat :d ::parseable-day)
             :ret boolean?
             :fn #(= (:ret %) false))
  (first (stest/check `bday? {:gen {::parseable-day (fn [] generate-day)}})))

(spec/fdef bday?
        :args (spec/cat :d ::parseable-day)
        :ret boolean?
        :fn #(= (:ret %)
                (string/starts-with?
                  (string/join "-"
                               (-> %
                                   :args
                                   :d
                                   (string/split #"-")
                                   next))
                  "03-15")))

(stest/check `bday? {:gen {::parseable-day (fn [] generate-day)}})


