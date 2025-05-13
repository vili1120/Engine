package devtools

import "os"

func InitMSG(prompt string) {
  println("[INIT] "+prompt)
}

func ErrorMSG(prompt string, err error) {
  println("[ERROR] "+prompt+err.Error())
  os.Exit(0)
}

func SuccessMSG(prompt string) {
  println("[SUCCESS] "+prompt+"\n")
}
