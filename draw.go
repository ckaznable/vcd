package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var tickFlag = true

func draw(s tcell.Screen, d Date) {
	joinStr := ":"
	if tickFlag {
		tickFlag = false
	} else {
		joinStr = " "
		tickFlag = true
	}

	s.Clear()

	if d.IsDefault() {
		drawText(s, "Have good vacation!")
	} else {
		str := fmt.Sprintf("%d%s%0*d%s%0*d", d.Day, joinStr, 2, d.Hours, joinStr, 2, d.Minutes)
		ascii, width, height := getASCII(str)
		drawASCII(s, ascii, width, height)
	}

	s.Show()
}

func drawText(s tcell.Screen, str string) {
	width, height := s.Size()
	y := height / 2
	x := width/2 - len(str)/2

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	for i, char := range str {
		s.SetContent(x+i, y, char, nil, style)
	}
}

func drawASCII(s tcell.Screen, bytes [][]rune, width int, height int) {
	screenWidth, screenHeight := s.Size()
	heightFix := 0
	if screenHeight%2 == 1 {
		heightFix = 1
	}

	y := screenHeight/2 - height/2 - heightFix
	if y < 0 {
		y = 0
	}

	x := screenWidth/2 - width/2
	if x < 0 {
		x = 0
	}

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreen)
	for i, row := range bytes {
		for j, char := range row {
			s.SetContent(x+j, y+i, char, nil, style)
		}
	}
}

func getASCII(str string) ([][]rune, int, int) {
	ascii := [][][]rune{}

	for _, char := range str {
		switch char {
		case '1':
			ascii = append(ascii, get1ASCII())
		case '2':
			ascii = append(ascii, get2ASCII())
		case '3':
			ascii = append(ascii, get3ASCII())
		case '4':
			ascii = append(ascii, get4ASCII())
		case '5':
			ascii = append(ascii, get5ASCII())
		case '6':
			ascii = append(ascii, get6ASCII())
		case '7':
			ascii = append(ascii, get7ASCII())
		case '8':
			ascii = append(ascii, get8ASCII())
		case '9':
			ascii = append(ascii, get9ASCII())
		case ':':
			ascii = append(ascii, getColonASCII())
		case '0':
			ascii = append(ascii, get0ASCII())
		default:
			ascii = append(ascii, getSpaceASCII())
		}

		ascii = append(ascii, getOneWidthSpaceASCII())
	}

	height := 5
	asciiStr := [][]rune{{}, {}, {}, {}, {}}

	for i := 0; i < height; i++ {
		for _, row := range ascii {
			asciiStr[i] = append(asciiStr[i], row[i]...)
		}
	}

	return asciiStr, len(asciiStr[0]), height
}

func get1ASCII() [][]rune {
	return [][]rune{
		{' ', '█', '█'},
		{'█', '█', '█'},
		{' ', '█', '█'},
		{' ', '█', '█'},
		{' ', '█', '█'},
	}
}

func get2ASCII() [][]rune {
	return [][]rune{
		{'█', '█', '█', '█', '█', '█', ' '},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', ' ', ' ', ' '},
		{'█', '█', '█', '█', '█', '█', '█'},
	}
}

func get3ASCII() [][]rune {
	return [][]rune{
		{'█', '█', '█', '█', '█', '█', ' '},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', ' '},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{'█', '█', '█', '█', '█', '█', ' '},
	}
}

func get4ASCII() [][]rune {
	return [][]rune{
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█'},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
	}
}

func get5ASCII() [][]rune {
	return [][]rune{
		{'█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', ' ', ' ', ' ', ' ', ' '},
		{'█', '█', '█', '█', '█', '█', '█'},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█'},
	}
}

func get6ASCII() [][]rune {
	return [][]rune{
		{' ', '█', '█', '█', '█', '█', '█'},
		{'█', '█', ' ', ' ', ' ', ' ', ' '},
		{'█', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', ' '},
	}
}

func get7ASCII() [][]rune {
	return [][]rune{
		{'█', '█', '█', '█', '█', '█', '█'},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{' ', ' ', ' ', ' ', '█', '█', ' '},
		{' ', ' ', ' ', '█', '█', ' ', ' '},
		{' ', ' ', ' ', '█', '█', ' ', ' '},
	}
}

func get8ASCII() [][]rune {
	return [][]rune{
		{' ', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', ' '},
	}
}

func get9ASCII() [][]rune {
	return [][]rune{
		{' ', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', '█'},
		{' ', ' ', ' ', ' ', ' ', '█', '█'},
		{'█', '█', '█', '█', '█', '█', ' '},
	}
}

func get0ASCII() [][]rune {
	return [][]rune{
		{' ', '█', '█', '█', '█', '█', '█', ' '},
		{'█', '█', ' ', ' ', '█', '█', '█', '█'},
		{'█', '█', ' ', '█', '█', ' ', '█', '█'},
		{'█', '█', '█', '█', ' ', ' ', '█', '█'},
		{' ', '█', '█', '█', '█', '█', '█', ' '},
	}
}

func getColonASCII() [][]rune {
	return [][]rune{
		{' ', ' '},
		{'█', '█'},
		{' ', ' '},
		{'█', '█'},
		{' ', ' '},
	}
}

func getSpaceASCII() [][]rune {
	return [][]rune{
		{' ', ' '},
		{' ', ' '},
		{' ', ' '},
		{' ', ' '},
		{' ', ' '},
	}
}

func getOneWidthSpaceASCII() [][]rune {
	return [][]rune{
		{' '},
		{' '},
		{' '},
		{' '},
		{' '},
	}
}
