type CommandFactory struct {
	decoder JsonDecoder // decoder decodes the command
}
// Create decode and validate the command
 func (cf CommandFactory) Create(encoded String) (Command, error) { 
	// decode command
	command, _:= cf.decoder.Decode(data) 
	return command, nil
 }

//  The dependency inversion principle suggests providing an interface I that provides the methods needed by class A, yet class B should implement the interface in order to get used by class A. This way one or many implementations of the interface I may exist, and class A can be used by other classes with different interfaces.

type Command struct{}
type CommandDecoder interface { Decode(data []byte) (Command, error)
type CommandFactory struct {
Decoder CommandDecoder
type JsonDecoder struct{}
func (jcd JsonDecoder) Decode(data string) (Command, error) {
	 // json command decode logic
}
// Initialize CommandFactory with required CommandDecoder 
factory := CommandFactory{Decoder: JsonCommandDecoder{}}
