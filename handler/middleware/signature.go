package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"crm/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type signState struct {
	mu   sync.Mutex
	data map[string]int64
}

var nonceStore = &signState{data: map[string]int64{}}

func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := viper.GetString("auth.sign.secret")
		if secret == "" {
			httputil.Forbidden(c)
			return
		}
		if !viper.GetBool("auth.sign.enable") {
			httputil.Forbidden(c)
			return
		}

		tsStr := c.GetHeader("X-Timestamp")
		nonce := c.GetHeader("X-Nonce")
		sign := c.GetHeader("X-Signature")
		method := c.GetHeader("X-Sign-Method")
		if method == "" {
			method = "HMAC-SHA256"
		}
		if tsStr == "" || nonce == "" || sign == "" {
			httputil.Forbidden(c)
			return
		}

		ts, err := parseInt64(tsStr)
		if err != nil {
			httputil.Forbidden(c)
			return
		}
		now := time.Now().Unix()
		skew := viper.GetInt64("auth.sign.skew_seconds")
		if skew <= 0 {
			skew = 300
		}
		if ts > now+skew || ts < now-skew {
			httputil.Forbidden(c)
			return
		}

		ttl := viper.GetInt64("auth.sign.nonce_ttl_seconds")
		if ttl <= 0 {
			ttl = 600
		}
		if isReplay(nonce, ts, ttl) {
			httputil.Forbidden(c)
			return
		}

		if strings.ToUpper(method) != "HMAC-SHA256" {
			httputil.Forbidden(c)
			return
		}

		canonicalQuery := canonicalizeQuery(c.Request.URL.Query())
		stringToSign := strings.Join([]string{
			c.Request.Method,
			c.Request.URL.Path,
			canonicalQuery,
			tsStr,
			nonce,
		}, "\n")

		expected := hmacSHA256Hex(secret, stringToSign)
		if !hmacEqual(expected, sign) {
			httputil.Forbidden(c)
			return
		}

		c.Next()
	}
}

func parseInt64(s string) (int64, error) {
	var x int64
	var n int
	var neg bool
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 && ch == '-' {
			neg = true
			continue
		}
		if ch < '0' || ch > '9' {
			return 0, http.ErrNotSupported
		}
		n = int(ch - '0')
		x = x*10 + int64(n)
	}
	if neg {
		x = -x
	}
	return x, nil
}

func canonicalizeQuery(q map[string][]string) string {
	if len(q) == 0 {
		return ""
	}
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	first := true
	for _, k := range keys {
		vals := q[k]
		sort.Strings(vals)
		for _, v := range vals {
			if !first {
				b.WriteByte('&')
			}
			first = false
			b.WriteString(k)
			b.WriteByte('=')
			b.WriteString(v)
		}
	}
	return b.String()
}

func hmacSHA256Hex(secret, data string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func hmacEqual(a, b string) bool {
	ab, bb := []byte(strings.ToLower(strings.TrimSpace(a))), []byte(strings.ToLower(strings.TrimSpace(b)))
	if len(ab) != len(bb) {
		return false
	}
	var res byte
	for i := 0; i < len(ab); i++ {
		res |= ab[i] ^ bb[i]
	}
	return res == 0
}

func isReplay(nonce string, ts int64, ttl int64) bool {
	nonceStore.mu.Lock()
	defer nonceStore.mu.Unlock()
	expireAt := ts + ttl
	now := time.Now().Unix()
	if now > expireAt {
		return true
	}
	if v, ok := nonceStore.data[nonce]; ok {
		if v > now {
			return true
		}
	}
	nonceStore.data[nonce] = expireAt
	if len(nonceStore.data) > 10000 {
		for k, v := range nonceStore.data {
			if v <= now {
				delete(nonceStore.data, k)
			}
		}
	}
	return false
}
