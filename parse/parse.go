package parse

import (
	"fmt"
	"strings"
	"strconv"
	"log"
)

type P struct {
	op byte
	es []*E
}

type E struct{
	p *P
	v float32
	pure bool
}

func newP(op byte) (p *P) {
	p = &P{}
	p.op = op
	return
}

func (e E) isPure() bool {
	return e.pure
}

func (e *E) simplify() *E {
	if e.isPure() {
		return e
	}
	e.p = e.p.simplify()
	return e
}

func (p *P) simplify() *P {
	if len(p.es) == 1 {
		p = p.es[0].p	// sub p wont be pure by design 
		p = p.simplify()
	}
	for _, e := range p.es {
		e = e.simplify()	// todo: check pointer works properly
	}
	return p
}

func (e E) String() string {
	if  e.isPure() {
		return fmt.Sprint(e.v)
	}
	return e.p.String()

}


func shiftRight(s string) string {
	return strings.Replace(s, "\n", "\n\t", -1)
}

func (p P) String() string {
	ss := make([]string, 1)
	for _,e := range p.es {
		ss = append(ss, e.String())
	}
	return  fmt.Sprintf("{%c:%v\n}", p.op, shiftRight(strings.Join(ss, "\n")))
}

func Parse(exp string) *P {
	ch := make(chan *E)
	es := []*E{}
	go parseAdd(exp, ch)
	for e := range ch {
		es = append(es, e)
	}

	p := newP('+')
	p.es = es
	p = p.simplify()
	return p
}

func parseAdd(exp string, ch chan *E) {
	defer close(ch)
	chsg := make(chan string)
	go parseBySign(exp, '+', chsg)
	for pe := range chsg {
		// fmt.Println("+: parseSign:", pe)
		e := &E{}
		if isPure(pe) {
			e.pure = true
			e.v = str2float32(pe)
			ch <- e
			continue
		}
		sch := make(chan *E)
		p := newP('-')
		go parseSub(pe, sch)
		for se := range sch {
			p.es = append(p.es, se)
		}
		e.pure = false
		e.p = p
		ch <- e
	}
}

func str2float32(s string) (f float32) {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println("type conversion error:", s, err)
		return 0.0
	}
	return float32(value)
}


func parseBySign(s string, sign byte, ch chan string) {
	buf := make([]byte, len(s))
	l := 0
	b := 0

	for i:= 0; i<len(s); i++ {
		if b>0 {
			if s[i] == ')' {
				b--
			}
			buf[l] = s[i]
			l++
			continue
		}
		if s[i] == '(' {
			b++
		}
		if s[i] == sign {
			if l>0 {
				ch <- string(buf[:l])
				// fmt.Println("parseSign:", string(buf[:l]))
			}
			l = 0
			continue
		}
		buf[l] = s[i]
		l++
	}
	if l > 0 {
		ch <- string(buf[:l])
		// fmt.Println("parseSign:", string(buf[:l]))
	}
	close(ch)
}


func parseSub(exp string, ch chan *E) {
	defer close(ch)
	chsg := make(chan string)
	go parseBySign(exp, '-', chsg)
	for pe := range chsg {
		// fmt.Println("-: parseSign:", pe)
		e := &E{}
		if isPure(pe) {
			e.pure = true
			e.v = str2float32(pe)
			ch <- e
			continue
		}
		sch := make(chan *E)
		p := newP('*')
		go parseProd(pe, sch)
		for se := range sch {
			p.es = append(p.es, se)
		}
		e.pure = false
		e.p = p
		ch <- e
	}
}


func parseProd(exp string, ch chan *E) {
	defer close(ch)
	chsg := make(chan string)
	go parseBySign(exp, '*', chsg)
	for pe := range chsg {
		// fmt.Println("*: parseSign:", pe)
		e := &E{}
		if isPure(pe) {
			e.pure = true
			e.v = str2float32(pe)
			ch <- e
			continue
		}
		sch := make(chan *E)
		p := newP('/')
		go parseDiv(pe, sch)
		for se := range sch {
			p.es = append(p.es, se)
		}
		e.pure = false
		e.p = p
		ch <- e
	}
}


func parseDiv(exp string, ch chan *E) {
	defer close(ch)
	chsg := make(chan string)
	go parseBySign(exp, '/', chsg)
	for pe := range chsg {
		// fmt.Println("/: parseSign:", pe)
		e := &E{}
		if isPure(pe) {
			e.pure = true
			e.v = str2float32(pe)
			ch <- e
			continue
		}
		sch := make(chan *E)
		p := newP('+')
		go parseAdd(removeBraces(pe), sch)
		for se := range sch {
			p.es = append(p.es, se)
		}
		e.pure = false
		e.p = p
		ch <- e
	}
}

func removeBraces(exp string) (s string) {
	if exp[0] != '(' {
		log.Fatal("expect (exp)")
	}
	return exp[1:len(exp)-1]
}

func isPure(s string) bool {
	return !strings.ContainsAny(s, "+-*/()")
}
