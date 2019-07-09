// nx/crypto handles cryptology related actions such as encryption & decryption.
package crypto

// EncryptBytes 
func EncryptBytes(key []byte, data interface{}) []byte {
  // TODO: write the code
  return nil
}

// Encrypt is a shortcut that just passed its args to EncryptBytes.
// It returns a string version of EncryptBytes return value.
func Encrypt(key []byte, data interface{}) string {
  return string(EncryptBytes(key, data))
}

func Decrypt(key []byte, data interface{}) {
  // TODO: write the code
}
