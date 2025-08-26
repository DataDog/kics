/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package tag

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/Checkmarx/kics/pkg/logger"
)

const (
	base      = 10
	bitSize64 = 64
)

// Tag contains the tag name reference and its attributes
type Tag struct {
	Name       string
	Attributes map[string]interface{}
}

// Parse tag from following structure
// name1:"expected=private,test=false" name2:"attr=1"
func Parse(ctx context.Context, s string, supportedNames []string) ([]Tag, error) {
	s = strings.TrimLeft(strings.TrimLeft(strings.TrimSpace(s), "/"), " ")
	var tags []Tag
	for _, si := range strings.Split(s, " ") {
		cleanSi := strings.TrimSpace(si)
		if cleanSi == "" {
			continue
		}

		for _, supportedName := range supportedNames {
			if !strings.HasPrefix(cleanSi, supportedName) {
				continue
			}

			tag, err := parseTag(ctx, cleanSi, supportedName)
			if err != nil {
				return nil, err
			}

			tags = append(tags, tag)
		}
	}

	return tags, nil
}

func parseTag(ctx context.Context, s, name string) (Tag, error) {
	logger := logger.FromContext(ctx)
	t := Tag{
		Name:       name,
		Attributes: make(map[string]interface{}),
	}

	attributePart := strings.TrimPrefix(s, name)
	attributePart = strings.TrimPrefix(attributePart, ":")
	attributePart = strings.TrimPrefix(attributePart, "\"")
	attributePart = strings.TrimSuffix(attributePart, "\"")

	if attributePart == "" {
		return t, nil
	}

	sc := &scanner.Scanner{}
	sc.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings
	sc.Init(strings.NewReader(attributePart))

	for {
		tok := sc.Scan()
		switch tok {
		case scanner.EOF:
			return t, nil
		case scanner.Ident:
			ident := sc.TokenText()
			switch sc.Peek() {
			case '=':
				sc.Next()
				value, err := parseValue(ctx, sc)
				if err != nil {
					return Tag{}, err
				}
				t.Attributes[ident] = value
			case '[':
				sc.Next()
				arg, err := parseArgs(ctx, sc)
				if err != nil {
					return Tag{}, err
				}
				t.Attributes[ident] = arg
			case ',':
				sc.Next()
				t.Attributes[ident] = nil
			case scanner.EOF:
				t.Attributes[ident] = nil
			}
		case ',':
			// NOP
		default:
			err := fmt.Errorf("invalid token: %s", sc.TokenText())
			logger.Error().Msg(err.Error())
			return Tag{}, err
		}
	}
}

func parseArray(ctx context.Context, sc *scanner.Scanner) ([]interface{}, error) {
	logger := logger.FromContext(ctx)
	var result []interface{}
	for {
		value, err := parseValue(ctx, sc)
		if err != nil {
			return result, err
		}
		result = append(result, value)
		next := sc.Next()
		if next == ']' {
			return result, nil
		}
		if next == ',' {
			continue
		}
		err = fmt.Errorf(", expected but got %s", string(next))
		logger.Error().Msg(err.Error())
		return result, err
	}
}

func parseValue(ctx context.Context, sc *scanner.Scanner) (interface{}, error) {
	logger := logger.FromContext(ctx)
	switch sc.Peek() {
	case '\'':
		sc.Next()
		return parseString(ctx, sc)
	case '*':
		r := sc.Next()
		return string(r), nil
	case '<', '>':
		r := sc.Next()
		if sc.Peek() == '=' {
			sc.Next()
			return string(r) + "=", nil
		}
		return string(r), nil
	case '!':
		sc.Next()
		if sc.Peek() == '=' {
			sc.Next()
			return "!=", nil
		}
		err := fmt.Errorf("invalid value: %s", sc.TokenText())
		logger.Error().Msg(err.Error())
		return nil, err
	case '[':
		sc.Next()
		return parseArray(ctx, sc)
	default:
		tok := sc.Scan()
		switch tok {
		case scanner.Ident:
			return checkType(sc.TokenText()), nil
		case scanner.String, scanner.Int, scanner.Float:
			if tok == scanner.String {
				str := sc.TokenText()
				return str[1 : len(str)-1], nil
			} else if tok == scanner.Int {
				return strconv.ParseInt(sc.TokenText(), base, bitSize64)
			} else if tok == scanner.Float {
				return strconv.ParseFloat(sc.TokenText(), bitSize64)
			}
		default:
			err := fmt.Errorf("invalid value: %s", sc.TokenText())
			logger.Error().Msg(err.Error())
			return nil, err
		}
	}
	return nil, errors.New("invalid value")
}

func parseArgs(ctx context.Context, sc *scanner.Scanner) (map[string]interface{}, error) {
	logger := logger.FromContext(ctx)
	result := map[string]interface{}{}
	for {
		tok := sc.Scan()
		if tok != scanner.Ident {
			err := fmt.Errorf("invalid attribute name: %s", sc.TokenText())
			logger.Error().Msg(err.Error())
			return result, err
		}
		name := sc.TokenText()
		eq := sc.Next()
		if eq != '=' {
			err := fmt.Errorf("= expected but got %s", string(eq))
			logger.Error().Msg(err.Error())
			return result, err
		}
		value, err := parseValue(ctx, sc)
		if err != nil {
			return result, err
		}
		result[name] = value
		next := sc.Next()
		if next == ']' {
			return result, nil
		}
		if next == ',' {
			continue
		}
		err = fmt.Errorf(") or , expected but got %s", string(next))
		logger.Error().Msg(err.Error())
		return result, err
	}
}

func parseString(ctx context.Context, sc *scanner.Scanner) (string, error) {
	var buf bytes.Buffer
	ch := sc.Next()
	for ch != '\'' {
		if ch == '\n' || ch == '\r' || ch < 0 {
			return "", errors.New("unterminated string")
		}
		if ch == '\\' {
			s, err := parseEscape(ctx, sc)
			if err != nil {
				return "", err
			}
			buf.WriteString(s)
		} else {
			buf.WriteRune(ch)
		}
		ch = sc.Next()
	}
	return buf.String(), nil
}

func parseEscape(ctx context.Context, sc *scanner.Scanner) (string, error) {
	logger := logger.FromContext(ctx)
	ch := sc.Next()
	switch ch {
	case 'a':
		return "\a", nil
	case 'b':
		return "\b", nil
	case 'f':
		return "\f", nil
	case 'n':
		return "\n", nil
	case 'r':
		return "\r", nil
	case 't':
		return "\t", nil
	case 'v':
		return "\v", nil
	case '\\':
		return "\\", nil
	case '"':
		return "\"", nil
	case '\'':
		return "'", nil
	}
	err := fmt.Errorf("invalid escape sequence: %s", string(ch))
	logger.Error().Msg(err.Error())
	return "", err
}

func checkType(s string) interface{} {
	switch s {
	case "true", "TRUE":
		return true
	case "false", "FALSE":
		return false
	default:
		if i, err := strconv.ParseInt(s, base, bitSize64); err == nil {
			return i
		}
		if f, err := strconv.ParseFloat(s, bitSize64); err == nil {
			return f
		}

		return s
	}
}
