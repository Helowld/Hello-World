;; Official website : https://kotlinlang.org/

;; This excution of Clojure on your terminal will print "Hello, World"
;; Additionally you need a project.clj containing:
;; (defproject hello-world "0.1.0-SNAPSHOT"
;;  :main hello-world.core
;;  :dependencies [[org.clojure/clojure "1.5.1"]])
;; Full intro on https://coderwall.com/p/yb1k-q/hello-world-in-clojure

(ns hello-world.core)

(defn -main [& args]
  (println "Hello, World!"))
