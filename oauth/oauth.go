package oauth

import (
    "log"
    "fmt"
    "time"
    "sort"
    "strings"
    "crypto/hmac"
    "crypto/rand"
    "crypto/sha1"
    "encoding/base64"
    "net/http"
    "net/url"
    "io/ioutil"
    "../cabundle"
)

type Credential struct {
    ConsumerKey string
    ConsumerSecret string
    OAuthToken string
    OAuthTokenSecret string
}

func mergeMap(maps ...map[string]string) map[string]string {
    hint := 0
    for _, m := range maps {
        hint += len(m)
    }
    newmap := make(map[string]string, hint)
    for _, m := range maps {
        for k, v := range m {
            newmap[k] = v
        }
    }
    return newmap
}

func forStrings(fn func(string) string, strs []string) []string {
    newstrs := make([]string, len(strs))
    for i, s := range strs {
        newstrs[i] = fn(s)
    }
    return newstrs
}

func encodeRFC3986(s string) string {
    return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}

func encodeMap(
    m map[string]string,
    separator string,
    enclosure string,
) string {
    keys := make([]string, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    sets := make([]string, len(m))
    for i, k := range keys {
        ek := encodeRFC3986(k)
        ev := encodeRFC3986(m[k])
        sets[i] = fmt.Sprintf("%s=%s%s%s", ek, enclosure, ev, enclosure)
    }
    return strings.Join(sets, separator)
}

func generateNonce(n int) string {
    const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    bytes := make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = chars[b % byte(len(chars))]
    }
    return string(bytes)
}

func generateSignature(data string, key string) string {
    hash := hmac.New(sha1.New, []byte(key))
    hash.Write([]byte(data))
    return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (t *Credential) renew(
    uri string,
    additionalParams map[string]string,
    verifier *string,
) *Credential {
    oauthParams := make(map[string]string, 5)
    oauthParams["oauth_consumer_key"] = t.ConsumerKey
    oauthParams["oauth_signature_method"] = "HMAC-SHA1"
    oauthParams["oauth_timestamp"] = fmt.Sprint(uint32(time.Now().Unix()))
    oauthParams["oauth_version"] = "1.0a"
    oauthParams["oauth_nonce"] = generateNonce(32)
    if t.OAuthToken != "" {
        oauthParams["oauth_token"] = t.OAuthToken
    } else if additionalParams["x_auth_mode"] == "" {
        oauthParams["oauth_callback"] = "oob"
    }
    if verifier != nil {
        oauthParams["oauth_verifier"] = *verifier
    }
    base := mergeMap(additionalParams, oauthParams)
    oauthParams["oauth_signature"] = generateSignature(
        strings.Join(forStrings(encodeRFC3986, []string {
            "POST",
            uri,
            encodeMap(base, "&", ""),
        }), "&"),
        strings.Join(forStrings(encodeRFC3986, []string {
            t.ConsumerSecret,
            t.OAuthTokenSecret,
        }), "&"),
    )
    reader := strings.NewReader(encodeMap(additionalParams, "&", ""))
    req, _ := http.NewRequest("POST", uri, reader)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Authorization", "OAuth " + encodeMap(oauthParams, ", ", "\""))
    resp, rerr := cabundle.GetClient().Do(req)
    if rerr != nil {
        log.Fatalln(rerr)
    }
    defer resp.Body.Close()
    bytes, _ := ioutil.ReadAll(resp.Body)
    values, qerr := url.ParseQuery(string(bytes))
    if qerr != nil {
        log.Fatalln(string(bytes))
    }
    ot := values.Get("oauth_token")
    os := values.Get("oauth_token_secret")
    if ot == "" || os == "" {
        log.Fatalln(string(bytes))
    }
    return &Credential{t.ConsumerKey, t.ConsumerSecret, ot, os}
}

func (t *Credential) RenewWithRequestToken() *Credential {
    return t.renew(
        "https://api.twitter.com/oauth/request_token",
        map[string]string {},
        nil,
    )
}

func (t *Credential) RenewWithAccessToken(
    params map[string]string,
    verifier *string,
) *Credential {
    return t.renew(
        "https://api.twitter.com/oauth/access_token",
        params,
        verifier,
    )
}
