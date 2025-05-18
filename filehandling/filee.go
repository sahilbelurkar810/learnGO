package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("ex.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	panic(err)
	// }
	// defer f.Close()
	// fmt.Println("File opened successfully")
	// info, err := f.Stat()
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	panic(err)
	// }
	// // fmt.Println("File name:", info.Name())
	// // fmt.Println("File size:", info.Size())
	// // fmt.Println("File mode:", info.Mode())
	// // fmt.Println("File modification time:", info.ModTime())
	// // fmt.Println("Is directory:", info.IsDir())
	// // fmt.Println("System type:", info.Sys())
	// // fmt.Println("File permissions:", info.Mode().Perm())

	// // read file
	// buf := make([]byte, info.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	panic(err)
	// }
	// println("Read bytes:", d)
	// println("Buffer:", string(buf))

	// f, err := os.ReadFile("ex.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	panic(err)
	// }

	// fmt.Println("File opened successfully")
	// fmt.Println("File content:", string(f))

	//read folder
	// dir, err := os.Open(".")
	// if err != nil {
	// 	fmt.Println("Error opening directory:", err)
	// 	panic(err)
	// }

	// defer dir.Close()
	//read all files in directory
	// files, err := dir.Readdir(-1)
	// if err != nil {
	// 	fmt.Println("Error reading directory:", err)
	// 	panic(err)
	// }
	// //print all files in directory
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	// f, err := os.Create("ex2.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	panic(err)
	// }
	// defer f.Close()
	// //write to file
	// _, err = f.WriteString("Hello World")
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	panic(err)
	// }

	// bytes := []byte("Hello World")
	// _, err = f.Write(bytes)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	panic(err)
	// }

	//read and write to another file in streaming fashion

	// sourceFile, err := os.Open("ex.txt")

	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("ex3.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	er := writer.WriteByte(b)
	// 	if er != nil {
	// 		fmt.Println("Error writing to file:", err)
	// 		panic(er)
	// 	}
	// }

	// e := writer.Flush()
	// if e != nil {
	// 	fmt.Println("Error flushing to file:", err)
	// 	panic(e)
	// }
	// fmt.Println("File copied successfully")

	// sourceFile, err := os.Open("ex3.txt")

	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err := os.Remove("ex3.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		panic(err)
	}
	fmt.Println("File deleted successfully")

}
