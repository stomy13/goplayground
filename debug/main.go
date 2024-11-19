package main

import (
	"debug/buildinfo"
	"debug/elf"
	"debug/gosym"
	"debug/macho"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// bi()
	// readElfFile()
	readTableFile()
	// time.Sleep(time.Minute)
}

func bi() {
	info, err := buildinfo.ReadFile("./deb")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", info)
	// ->
	// go      go1.19
	// path    bi
	// mod     bi      (devel)
	// build   -compiler=gc
	// build   CGO_ENABLED=1
	// build   CGO_CFLAGS=
	// build   CGO_CPPFLAGS=
	// build   CGO_CXXFLAGS=
	// build   CGO_LDFLAGS=
	// build   GOARCH=arm64
	// build   GOOS=darwin
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ioReader(file string) io.ReaderAt {
	r, err := os.Open(file)
	check(err)
	return r
}

func readElfFile() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: elftest elf_file")
		os.Exit(1)
	}
	f := ioReader(os.Args[1])
	_elf, err := elf.NewFile(f)
	check(err)

	// Read and decode ELF identifier
	var ident [16]uint8
	f.ReadAt(ident[0:], 0)
	check(err)

	if ident[0] != '\x7f' || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' {
		fmt.Printf("Bad magic number at %d\n", ident[0:4])
		os.Exit(1)
	}

	var arch string
	switch _elf.Class.String() {
	case "ELFCLASS64":
		arch = "64 bits"
	case "ELFCLASS32":
		arch = "32 bits"
	}

	var mach string
	switch _elf.Machine.String() {
	case "EM_AARCH64":
		mach = "ARM64"
	case "EM_386":
		mach = "x86"
	case "EM_X86_64":
		mach = "x86_64"
	}

	fmt.Printf("File Header: ")
	fmt.Println(_elf.FileHeader)
	fmt.Printf("ELF Class: %s\n", arch)
	fmt.Printf("Machine: %s\n", mach)
	fmt.Printf("ELF Type: %s\n", _elf.Type)
	fmt.Printf("ELF Data: %s\n", _elf.Data)
	fmt.Printf("Entry Point: %d\n", _elf.Entry)
	fmt.Printf("Section Addresses: %d\n", _elf.Sections)

}

func readTableFile() {
	table, err := getTable(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	path, _, _ := table.PCToLine(table.LookupFunc("main.main").Entry)
	fmt.Println(path)
	path, _, _ = table.PCToLine(table.LookupFunc("runtime.Version").Entry)
	fmt.Println(path)
}

func getTable(file string) (*gosym.Table, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var textStart uint64
	var symtab, pclntab []byte

	obj, err := elf.NewFile(f)
	if err == nil {
		if sect := obj.Section(".text"); sect == nil {
			return nil, errors.New("empty .text")
		} else {
			textStart = sect.Addr
		}
		if sect := obj.Section(".gosymtab"); sect != nil {
			if symtab, err = sect.Data(); err != nil {
				return nil, err
			}
		}
		if sect := obj.Section(".gopclntab"); sect != nil {
			if pclntab, err = sect.Data(); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("empty .gopclntab")
		}

	} else {
		obj, err := macho.NewFile(f)
		if err != nil {
			return nil, err
		}

		if sect := obj.Section("__text"); sect == nil {
			return nil, errors.New("empty __text")
		} else {
			textStart = sect.Addr
		}
		if sect := obj.Section("__gosymtab"); sect != nil {
			if symtab, err = sect.Data(); err != nil {
				return nil, err
			}
		}
		if sect := obj.Section("__gopclntab"); sect != nil {
			if pclntab, err = sect.Data(); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("empty __gopclntab")
		}
	}

	pcln := gosym.NewLineTable(pclntab, textStart)
	return gosym.NewTable(symtab, pcln)
}
