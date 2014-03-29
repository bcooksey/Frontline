// Taken from https://github.com/kjk/fofou/blob/0033b5b909a5e3d551cc46b8d4f3af0afb9155dc/go/log.go
package main

import (
    "fmt"
    "time"
)

type TimestampedMsg struct {
    Time time.Time
    Msg  string
}

type CircularMessagesBuf struct {
    Msgs []TimestampedMsg
    pos  int
    full bool
}

func (m *TimestampedMsg) TimeStr() string {
    return m.Time.Format("2006-01-02 15:04:05")
}

func (m *TimestampedMsg) TimeSinceStr() string {
    return TimeSinceNowAsString(m.Time)
}

func NewCircularMessagesBuf(cap int) *CircularMessagesBuf {
    return &CircularMessagesBuf{
        Msgs: make([]TimestampedMsg, cap, cap),
        pos:  0,
        full: false,
    }
}

func (b *CircularMessagesBuf) Add(s string) {
    var msg = TimestampedMsg{time.Now(), s}
    if b.pos == cap(b.Msgs) {
        b.pos = 0
        b.full = true
    }
    b.Msgs[b.pos] = msg
    b.pos += 1
}

func (b *CircularMessagesBuf) GetOrdered() []*TimestampedMsg {
    size := b.pos
    if b.full {
        size = cap(b.Msgs)
    }
    res := make([]*TimestampedMsg, size, size)
    for i := 0; i < size; i++ {
        p := b.pos - 1 - i
        if p < 0 {
            p = cap(b.Msgs) + p
        }
        res[i] = &b.Msgs[p]
    }
    return res
}

type ServerLogger struct {
    Errors    *CircularMessagesBuf
    Notices   *CircularMessagesBuf
    UseStdout bool
}

func NewServerLogger(errorsMax, noticesMax int, useStdout bool) *ServerLogger {
    l := &ServerLogger{
        Errors:    NewCircularMessagesBuf(errorsMax),
        Notices:   NewCircularMessagesBuf(noticesMax),
        UseStdout: useStdout,
    }
    return l
}

func (l *ServerLogger) Error(s string) {
    l.Errors.Add(s)
    fmt.Printf("Error: %s\n", s)
}

func (l *ServerLogger) Errorf(format string, v ...interface{}) {
    s := fmt.Sprintf(format, v...)
    l.Errors.Add(s)
    fmt.Printf("Error: %s\n", s)
}

func (l *ServerLogger) Notice(s string) {
    l.Notices.Add(s)
    fmt.Printf("%s\n", s)
}

func (l *ServerLogger) Noticef(format string, v ...interface{}) {
    s := fmt.Sprintf(format, v...)
    l.Notices.Add(s)
    fmt.Printf("%s\n", s)
}

func (l *ServerLogger) GetErrors() []*TimestampedMsg {
    return l.Errors.GetOrdered()
}

func (l *ServerLogger) GetNotices() []*TimestampedMsg {
    return l.Notices.GetOrdered()
}

func TimeSinceNowAsString(t time.Time) string {
    d := time.Now().Sub(t)
    minutes := int(d.Minutes()) % 60
    hours := int(d.Hours())
    days := hours / 24
    hours = hours % 24
    if days > 0 {
        return fmt.Sprintf("%dd %dhr", days, hours)
    }
    if hours > 0 {
        return fmt.Sprintf("%dhr %dm", hours, minutes)
    }
    return fmt.Sprintf("%dm", minutes)
}
