package core

import (
	err "in-memory-db/1/error_code"
	"math"
	"strconv"
	"strings"
	"time"
)

type ParsedCommand struct {
	Action    string
	key       string
	value     string
	qvalue    []string
	expiry    int32
	timestamp int64
	condition string
	bwait     float64
}

// func (p *ParsedCommand) GetAction {}

var emptyResult = ParsedCommand{}

func CommandParser(command string) (error, ParsedCommand) {
	cmdStrSplit := strings.Split(strings.Trim(command, " "), " ")
	cmd := make([]string, int(math.Max(6, float64(len(cmdStrSplit)))))
	copy(cmd, cmdStrSplit)
	action := cmd[0]
	result := ParsedCommand{action, cmd[1], "", []string{}, -1, (time.Now().UnixMilli() / 1000), "", 0}
	switch action {
	case "QPOP":
		fallthrough
	case "GET":
		if cmd[2] != "" {
			return err.ErrInvCommand, emptyResult
		} else {
			return nil, result
		}
	case "BQPOP":
		{
			wait, er := strconv.ParseFloat(cmd[2], 64)
			if er != nil {
				return err.ErrInvParams, emptyResult
			}
			result.bwait = wait
			return nil, result
		}
	case "SET":
		{
			result.value = cmd[2]
			if result.value == "" {
				return err.ErrInvParams, emptyResult
			}
			opt := cmd[3]
			switch opt {
			case "NX":
				fallthrough
			case "XX":
				{
					result.condition = opt
					return nil, result
				}
			case "EX":
				{
					ex, er := strconv.ParseInt(cmd[4], 10, 32)
					if er == nil {
						result.expiry = int32(ex)
						result.condition = cmd[5]
						return nil, result
					} else {
						return err.ErrInvParams, emptyResult
					}
				}
			case "":
				return nil, result
			default:
				{
					return err.ErrInvParams, emptyResult
				}
			}
		}

	case "QPUSH":
		{
			valueLen := len(cmdStrSplit) - 2
			if valueLen < 1 {
				return err.ErrInvParams, emptyResult
			}
			result.qvalue = cmd[2 : valueLen+2]
			return nil, result
		}
	default:
		{
			return err.ErrInvCommand, emptyResult
		}
	}
}
