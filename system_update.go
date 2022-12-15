// Package advent Day 7
package advent

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	dollarSign = 36
	period     = 46
	slash      = 47
	move       = 99
	directory  = 100
)

type lineKind int

const (
	cd lineKind = iota
	ls
	dir
	file
)

type Line struct {
	kind lineKind
}

type Dir struct {
	children []*Dir
	files    []*File
	name     string
	parent   *Dir
	size     int
}

type File struct {
	name string
	size int
}

func TraverseDirs(data *[]byte) (*Dir, error) {
	var endOfLine, prevIdx int
	var rootDir = &Dir{parent: nil, name: "root"}
	var curDir = rootDir

	for byteIndex, b := range *data {
		if b == newline {
			endOfLine++
			if endOfLine == 1 && prevIdx != byteIndex {
				lineData := (*data)[prevIdx:byteIndex]
				line, err := categorizeLine(&lineData)
				if err != nil {
					return nil, err
				}
				switch {
				case line.kind == cd:
					curDir, err = processMove(&lineData, curDir)
					if err != nil {
						return nil, err
					}
				case line.kind == dir:
					curDir.addChild(&lineData)
				case line.kind == file:
					err = curDir.addFile(&lineData)
					if err != nil {
						return nil, err
					}
				}
			}
			prevIdx = byteIndex + 1
		} else {
			endOfLine = 0
		}
	}

	rootDir.setSize()

	return rootDir, nil
}

func FindSystemDirs(root *Dir) *[]*Dir {
	result := make([]*Dir, 0)
	root.getChildren(&result)

	return &result
}

func FilterDirsBySize(dirs *[]*Dir, size int) *[]*Dir {
	result := make([]*Dir, 0)

	for _, d := range *dirs {
		if d.size >= size {
			result = append(result, d)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].size < result[j].size
	})

	return &result
}

func categorizeLine(data *[]byte) (*Line, error) {
	line := &Line{}
	switch (*data)[0] {
	case dollarSign:
		if (*data)[2] == move {
			line.kind = cd
		} else {
			line.kind = ls
		}
	case directory:
		line.kind = dir
	default:
		if (*data)[0] >= 48 && (*data)[0] <= 57 {
			line.kind = file
		} else {
			return nil, errors.New("could not categorize line")
		}
	}

	return line, nil
}

func processMove(line *[]byte, dir *Dir) (*Dir, error) {
	if (*line)[5] == period && (*line)[6] == period {
		return dir.parent, nil
	} else {
		if (*line)[5] == slash {
			return dir, nil
		} else {
			return dir.findChild(string(*line)[5:])
		}
	}
}

func newDir(parent *Dir, name string) *Dir {
	return &Dir{parent: parent, name: name}
}

func (d *Dir) getChildren(dirs *[]*Dir) {
	for _, c := range d.children {
		c.getChildren(dirs)
	}
	*dirs = append(*dirs, d.children...)
}

func (d *Dir) setSize() {
	for _, f := range d.files {
		d.size += f.size
	}
	for _, c := range d.children {
		c.setSize()
		d.size += c.size
	}
}

func (d *Dir) findChild(name string) (*Dir, error) {
	for _, c := range d.children {
		if c.name == name {
			return c, nil
		}
	}

	return nil, fmt.Errorf("no child directory with name %s found", name)
}

func (d *Dir) addChild(line *[]byte) {
	d.children = append(d.children, newDir(d, string((*line)[4:])))
}

func (d *Dir) addFile(line *[]byte) error {
	fileData := strings.Split(string(*line), " ")
	fileSize, err := strconv.Atoi(fileData[0])
	if err != nil {
		return err
	}
	d.files = append(d.files, &File{name: fileData[1], size: fileSize})

	return nil
}

func (d *Dir) GetSize() int {
	return d.size
}

func (d *Dir) GetParent() string {
	return d.parent.name
}

func (d *Dir) GetName() string {
	return d.name
}
