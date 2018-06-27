package common

// import (
//   "fmt"
// )

// type Checker struct {
//   lastError error
//   desc string
// }

// func NewChecker() *Checker { return &Checker{} }

// func (ck *Checker) Check(desc string, fn func() error) {
//   if ck.lastError != nil {
//     return
//   }
//   ck.desc = desc
//   ck.lastError = fn()
// }

// func (ck *Checker) LastError() error { return ck.lastError }

// func (ck *Checker) Error() string {
//   if ck.LastError() == nil {
//     return ""
//   }
//   if ck.desc == "" {
//     return ck.LastError().Error()
//   }
//   return fmt.Sprintf("%s: %v", ck.desc, ck.lastError)
// }

type Checker struct {
  Error bool
}

func NewChecker() *Checker { return &Checker{} }


func (ck *Checker) String(str string, invalid_value string) string {
  if str == invalid_value {
    ck.Error = true
  }
  return str
}

func (ck *Checker) Int(str int, invalid_value int) int {
  if str == invalid_value {
    ck.Error = true
  }
  return str
}

