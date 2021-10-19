package parser

import (
	"bufio"
	"strings"
)

type Parser struct {
	curQualifiers string
	curTimestamp  string
	curCells      []Cell
}

func New() Parser {
	return Parser{}
}

type Row struct {
	Key     string
	Columns map[string][]Cell
}

type Cell struct {
	Value     string
	Timestamp string
}

func (p *Parser) Parse(str string) ([]Row, error) {
	result := make([]Row, 0)
	rows := p.splitRows(str)[1:]
	for _, r := range rows {
		result = append(result, p.parseRow(r))
	}
	return result, nil
}

var rowDelimiter = strings.Repeat("-", 40)

func (p *Parser) splitRows(str string) []string {
	return strings.Split(str, rowDelimiter+"\n")
}

func (p *Parser) parseRow(str string) Row {
	row := Row{Columns: make(map[string][]Cell, 0)}
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		l := scanner.Text()
		if !strings.HasPrefix(l, " ") {
			// rowkey
			row.Key = l
			continue
		}
		if strings.HasPrefix(l, strings.Repeat(" ", 4)) {
			// value
			c := Cell{
				Value:     l[4:],
				Timestamp: p.curTimestamp,
			}
			p.curCells = append(p.curCells, c)
		} else {
			// column
			q, t := p.parseColumn(l)
			if _, ok := row.Columns[q]; !ok {
				if p.curQualifiers != "" && q != p.curQualifiers {
					row.Columns[p.curQualifiers] = p.curCells
					p.curCells = make([]Cell, 0)
				}
			}
			p.curQualifiers = q
			p.curTimestamp = t
		}
	}
	row.Columns[p.curQualifiers] = p.curCells
	return row
}

func (p *Parser) parseColumn(str string) (string, string) {
	s := strings.TrimSpace(str)
	atIndex := strings.LastIndex(s, "@")
	return strings.TrimSpace(s[:atIndex]), strings.TrimSpace(s[atIndex+1:])
}
