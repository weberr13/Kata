package warmup

import (
   "fmt"
   "testing"
   "strconv"
)

type EatsThings struct {
   input chan interface{}
   done chan interface{}
   end chan interface{}
   gobbler []interface{}
}
func MakeEater() (*EatsThings) {
   return &EatsThings{input: make(chan interface{}, 1),
                      end: make(chan interface{}, 1),
                      done: make(chan interface {},1 ),
                      gobbler: make([]interface{}, 0, 10)}
}
func (eater *EatsThings) eat() {
   loop:
   for {
      select {
      case thing := <- eater.input:
         eater.gobbler = append(eater.gobbler, thing)
      case <- eater.end:
         inner:
         for {
            select {
            case thing := <- eater.input:
               eater.gobbler = append(eater.gobbler, thing)
            default:
               break inner
            }
         }
         break loop
      }
   }
   eater.done <- struct{}{}
}


func TestAmIWarm(t *testing.T) {
   empty := make(map[string]int)
   if len(empty) != 0 {
      t.Fail()
   }
   for i := 0 ; i < 10 ; i++ {
      empty[fmt.Sprint(i)] = i
   }
   if len(empty) != 10 {
      t.Fail()
   }
   for k, v := range empty {
      parsed, err := strconv.ParseInt(k, 10, 32)
      if err != nil {
         t.Fatalf("can't parse a string")
      }
      if v != int(parsed) {
         t.Fatalf("%v != %v", k, v)
      }
   }
   eater := MakeEater()
   go eater.eat()
   for i:= 0 ; i < 10 ; i++ {
      eater.input <- i
   }
   eater.end <- struct{}{}
   loop:
   for {
      select {
      case <- eater.done:
         break loop
      }
   }
   if len(eater.gobbler) != 10 {
      for _, v := range eater.gobbler {
         fmt.Println("got: ", v)
      }
      t.Fatalf("failed to get expected things: %v", len(eater.gobbler))
   }

   x := 0
   f := func () {x++}
   f()
   f()
   if x != 2 {
      t.Fatalf("closure fail!")
   }

   x = 0
   f2 := func() int {return x+1}
   x = f2()
   x = f2()
   if x != 2 {
      t.Fatalf("impure fail")
   }

   fs := make(map[string]func() int)
   fs["impure"] = f2
   x = 0
   x = fs["impure"]()
   x = fs["impure"]()
   if x != 2 {
      t.Fatalf("impure map fail")
   }
} 


