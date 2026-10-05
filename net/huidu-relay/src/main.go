package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	deviceIP   = flag.String("device", "", "IP real da placa (ex.: 192.168.11.200)")
	advertise  = flag.String("advertise", "", "IP anunciado as apps (ex.: 10.0.3.27)")
	rewriteFlg = flag.String("rewrite", "", "IP a reescrever nos payloads (default: -device)")
	udpPorts   = flag.String("udp", "9527", "portas UDP de relay, virgula-separadas")
	tcpPorts   = flag.String("tcp", "9527", "portas TCP de proxy, virgula-separadas")
	bindAddr   = flag.String("bind", "0.0.0.0", "endereco de bind dos listeners")
	sessionTTL = flag.Duration("session_ttl", 5*time.Minute, "TTL de sessao UDP inativa")
	singlePC   = flag.Bool("single_pc", true, "permitir TCP de apenas um PC (IP) por vez")
	verbose    = flag.Bool("v", false, "log verboso de pacotes")
)

func vlog(format string, args ...interface{}) {
	if *verbose {
		log.Printf(format, args...)
	}
}

type session struct {
	asker *net.UDPAddr
	up    *net.UDPConn
	last  time.Time
	dead  bool
}

type udpRelay struct {
	port     int
	listen   *net.UDPConn
	device   *net.UDPAddr
	loopIP   net.IP
	rewIP    net.IP
	advIP    net.IP
	ttl      time.Duration
	mu       sync.Mutex
	sessions map[string]*session
}

func (r *udpRelay) get(src *net.UDPAddr) *session {
	key := src.String()
	r.mu.Lock()
	defer r.mu.Unlock()
	if s, ok := r.sessions[key]; ok {
		return s
	}
	up, err := net.ListenUDP("udp4", nil)
	if err != nil {
		log.Printf("udp %d: sessao %s: %v", r.port, key, err)
		return nil
	}
	s := &session{asker: src, up: up, last: time.Now()}
	r.sessions[key] = s
	go r.pump(s)
	vlog("udp %d: nova sessao %s", r.port, key)
	return s
}

func (r *udpRelay) pump(s *session) {
	buf := make([]byte, 65535)
	for {
		n, _, err := s.up.ReadFromUDP(buf)
		if err != nil {
			return
		}
		pkt := make([]byte, n)
		copy(pkt, buf[:n])
		if cnt := rewriteIP(pkt, r.rewIP, r.advIP); cnt > 0 {
			vlog("udp %d: rewrite %d ocorrencia(s) do IP da placa", r.port, cnt)
		}
		if _, err := r.listen.WriteToUDP(pkt, s.asker); err != nil {
			log.Printf("udp %d: resposta para %s: %v", r.port, s.asker, err)
			return
		}
		vlog("udp %d: device -> %s (%d B)", r.port, s.asker, n)
	}
}

func (r *udpRelay) gc() {
	for range time.Tick(30 * time.Second) {
		r.mu.Lock()
		for k, s := range r.sessions {
			if time.Since(s.last) > r.ttl {
				s.dead = true
				s.up.Close()
				delete(r.sessions, k)
				vlog("udp %d: sessao %s expirada", r.port, k)
			}
		}
		r.mu.Unlock()
	}
}

func (r *udpRelay) run() {
	go r.gc()
	buf := make([]byte, 65535)
	for {
		n, src, err := r.listen.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf("udp %d: %v", r.port, err)
		}
		if src.IP.Equal(r.loopIP) {
			continue
		}
		s := r.get(src)
		if s == nil {
			continue
		}
		r.mu.Lock()
		if s.dead {
			r.mu.Unlock()
			continue
		}
		s.last = time.Now()
		r.mu.Unlock()
		if _, err := s.up.WriteToUDP(buf[:n], r.device); err != nil {
			vlog("udp %d: forward %s -> device: %v", r.port, src, err)
			continue
		}
		vlog("udp %d: %s -> device (%d B)", r.port, src, n)
	}
}

func rewriteIP(b []byte, from, to net.IP) int {
	f := from.To4()
	t := to.To4()
	n := 0
	for i := 0; i+4 <= len(b); i++ {
		if bytes.Equal(b[i:i+4], f) {
			copy(b[i:i+4], t)
			n++
			i += 3
		}
	}
	return n
}

type tcpProxy struct {
	port    int
	device  string
	single  bool
	mu      sync.Mutex
	clients map[string]*int32
}

func (p *tcpProxy) otherActive(ip string) string {
	p.mu.Lock()
	defer p.mu.Unlock()
	for k, v := range p.clients {
		if k != ip && *v > 0 {
			return k
		}
	}
	return ""
}

func (p *tcpProxy) acquire(ip string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.clients[ip] == nil {
		c := int32(0)
		p.clients[ip] = &c
	}
	*p.clients[ip]++
}

func (p *tcpProxy) release(ip string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if c := p.clients[ip]; c != nil {
		*c--
		if *c <= 0 {
			delete(p.clients, ip)
		}
	}
}

func (p *tcpProxy) handle(c net.Conn) {
	ip, _, _ := net.SplitHostPort(c.RemoteAddr().String())
	if p.single {
		if other := p.otherActive(ip); other != "" {
			log.Printf("tcp %d: %s recusado (PC %s ja conectado)", p.port, c.RemoteAddr(), other)
			c.Close()
			return
		}
	}
	p.acquire(ip)
	defer p.release(ip)
	d, err := net.DialTimeout("tcp", p.device, 5*time.Second)
	if err != nil {
		log.Printf("tcp %d: conectar a placa: %v", p.port, err)
		c.Close()
		return
	}
	log.Printf("tcp %d: %s <-> placa", p.port, c.RemoteAddr())
	defer d.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { io.Copy(d, c); c.Close(); d.Close(); wg.Done() }()
	go func() { io.Copy(c, d); c.Close(); d.Close(); wg.Done() }()
	wg.Wait()
	log.Printf("tcp %d: %s encerrada", p.port, c.RemoteAddr())
}

func (p *tcpProxy) run(ln net.Listener) {
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalf("tcp %d: %v", p.port, err)
		}
		go p.handle(c)
	}
}

func parsePorts(s string) [][2]int {
	var out [][2]int
	for _, f := range strings.Split(s, ",") {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}
		parts := strings.Split(f, ":")
		listen, device := 0, 0
		var err error
		switch len(parts) {
		case 1:
			listen, err = strconv.Atoi(parts[0])
			device = listen
		case 2:
			listen, err = strconv.Atoi(parts[0])
			if err == nil {
				device, err = strconv.Atoi(parts[1])
			}
		default:
			err = strconv.ErrSyntax
		}
		if err != nil || listen <= 0 || listen > 65535 || device <= 0 || device > 65535 {
			log.Fatalf("porta invalida: %q (use N ou N:M)", f)
		}
		out = append(out, [2]int{listen, device})
	}
	return out
}

func main() {
	flag.Parse()
	if *deviceIP == "" || *advertise == "" {
		log.Fatal("uso: huidu-relay -device <ip placa> -advertise <ip vpn> [-udp 9527] [-tcp 9527]")
	}
	dev := net.ParseIP(*deviceIP)
	adv := net.ParseIP(*advertise)
	if dev == nil || dev.To4() == nil || adv == nil || adv.To4() == nil {
		log.Fatal("-device e -advertise devem ser IPv4")
	}
	rew := dev
	if *rewriteFlg != "" {
		rew = net.ParseIP(*rewriteFlg)
		if rew == nil || rew.To4() == nil {
			log.Fatal("-rewrite deve ser IPv4")
		}
	}
	bind := net.ParseIP(*bindAddr)
	if bind == nil {
		log.Fatal("-bind invalido")
	}

	for _, ports := range parsePorts(*udpPorts) {
		port, devPort := ports[0], ports[1]
		laddr := &net.UDPAddr{IP: bind, Port: port}
		l, err := net.ListenUDP("udp4", laddr)
		if err != nil {
			log.Fatalf("udp %d: %v", port, err)
		}
		r := &udpRelay{
			port:     port,
			listen:   l,
			device:   &net.UDPAddr{IP: dev, Port: devPort},
			loopIP:   dev,
			rewIP:    rew,
			advIP:    adv,
			ttl:      *sessionTTL,
			sessions: map[string]*session{},
		}
		log.Printf("udp relay :%d -> %s (anuncia %s)", port, r.device, adv)
		go r.run()
	}

	for _, ports := range parsePorts(*tcpPorts) {
		port, devPort := ports[0], ports[1]
		ln, err := net.Listen("tcp4", net.JoinHostPort(bind.String(), strconv.Itoa(port)))
		if err != nil {
			log.Fatalf("tcp %d: %v", port, err)
		}
		p := &tcpProxy{port: port, device: net.JoinHostPort(dev.String(), strconv.Itoa(devPort)), single: *singlePC, clients: map[string]*int32{}}
		log.Printf("tcp proxy :%d -> %s", port, p.device)
		go p.run(ln)
	}

	log.Printf("huidu-relay ativo: device=%s advertise=%s", dev, adv)
	select {}
}
