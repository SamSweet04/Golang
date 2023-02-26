package main
 
import "fmt"
 
type Stream interface{
    read() string
    write(string)
    close()
}
 
func writeToStream(stream Stream, text string){
    stream.write(text)
}
func closeStream(stream Stream){
    stream.close()
}
 
// ��������� ����
type File struct{
    text string
}
// ��������� �����
type Folder struct{}
 
// ���������� ������� ��� ���� *File
func (f *File) read() string{
    return f.text
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("������ � ���� ������", message)
}
func (f *File) close(){
    fmt.Println("���� ������")
}
// ��������� ������� ��� ���� *Folder
func (f *Folder) close(){
    fmt.Println("����� �������")
}
 
func main() {
     
    myFile := &File{}
    myFolder := &Folder{}
     
    writeToStream(myFile, "hello world")
    closeStream(myFile)
    //closeStream(myFolder)     // ������: ��� *Folder �� ��������� ��������� Stream
    myFolder.close()            // ��� �����
}