package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type NameP struct {
	Name     string `json:"name"`
	Password string `json:"pw"`
}
type Claims struct {
	Userid int `json:"useid"`
	jwt.RegisteredClaims
}
type Cfg struct {
	DatabaseDSN       string `json:"database_dsn"`
	SiliconFlowAPIKey string `json:"siliconflow_api_key"`
	JWTSigningSecret  string `json:"jwt_signing_secret"`
	ServerHost        string `json:"server_host"`
	ServerPort        string `json:"server_port"`
	CookieDomain      string `json:"cookie_domain"`
	CookieSecure      bool   `json:"cookie_secure"`
}

var cfg Cfg

type Embedding struct {
	Em []float64 `json:"em"`
}

func api() *gin.Engine {
	data, err4 := os.ReadFile("config.json")
	if err4 != nil {
		log.Fatal("config read err")
	}

	err5 := json.Unmarshal(data, &cfg)
	if err5 != nil {
		log.Fatal("config read err")
	}
	db, err1 := sql.Open("mysql", cfg.DatabaseDSN)
	if err1 != nil {
		log.Fatal("err sql connection")
	}

	r := gin.Default()

	r.GET("/api/islog", func(c *gin.Context) {
		_, err := getid(c)
		if err != nil {
			c.Status(401)
		} else {
			c.Status(200)
		}

	})

	r.POST("/api/or", func(c *gin.Context) {
		var tzhi NameP
		err := c.ShouldBindJSON(&tzhi)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "error in login",
			})
		}
		tpass := ""
		var id int
		err2 := db.QueryRow(`select password,id from users where 
	 username=?`, tzhi.Name).Scan(&tpass, &id)

		if errors.Is(err2, sql.ErrNoRows) {
			result, err := db.Exec(`insert into users(username,password)
		values(?,?)`, tzhi.Name, tzhi.Password)
			if err != nil {
				c.Status(400)
			}
			newid, err := result.LastInsertId()
			id = int(newid)
		} else if err2 != nil {
			c.JSON(400, gin.H{
				"error": "registe_err",
			})
			return
		} else if tzhi.Password != tpass {
			c.JSON(400, gin.H{
				"error": "err_password",
			})
			return
		}
		claims := Claims{
			Userid: id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokentext, err6 := token.SignedString([]byte(cfg.JWTSigningSecret))
		if err6 != nil {
			log.Fatal("err in token")
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("login", tokentext, 24*60*60, "/", cfg.CookieDomain, cfg.CookieSecure, true)
		c.Status(200)

	})

	r.POST("/api/logout", func(c *gin.Context) {
		c.SetCookie(
			"login",
			"",
			-1,
			"/",
			cfg.CookieDomain,
			cfg.CookieSecure,
			true,
		)

		c.Status(200)
	})

	r.POST("/api/load", func(c *gin.Context) {
		type Ques struct {
			Q string `json:"question"`
		}
		var ques Ques
		err := c.ShouldBindJSON(&ques)

		if err != nil {
			log.Print("err in load note")
			c.Status(400)
			return
		}
		err, quesfloat := rag(ques.Q)
		if err != nil {
			log.Print("err in load note")
			c.Status(400)
			return
		}
		var notes []Notes
		id, err := getid(c)
		order := ` select id,userid,content,title,created_at,updated_at,embedding
		 from notes where userid =?;`
		result, err := db.Query(order, id)

		if err != nil {
			log.Print("err in load note")
			c.Status(400)
			return
		}
		defer result.Close()
		for result.Next() {
			var tnote Notes
			result.Scan(&tnote.ID, &tnote.UserID, &tnote.Content, &tnote.Title,
				&tnote.Createdtime, &tnote.Updatetime, &tnote.Embedding)
			notes = append(notes, tnote)
			index := len(notes) - 1

			type Embedding struct {
				Em []float64 `json:"em"`
			}
			var t1 Embedding
			err = json.Unmarshal([]byte(notes[index].Embedding), &t1)
			confloat := t1.Em
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for index1, _ := range quesfloat {
				sum1 += confloat[index1] * quesfloat[index1]
				sum2 += confloat[index1] * confloat[index1]
				sum3 += quesfloat[index1] * quesfloat[index1]
			}
			tfin := (math.Sqrt(sum2) * math.Sqrt(sum3))
			if tfin == 0 {
				notes[index].Cos = 0
				continue
			}
			fin := sum1 / tfin
			notes[index].Cos = fin
			if strings.TrimSpace(ques.Q) == "" {
				notes[index].Cos = 0
			}
		}
		sort.SliceStable(notes, func(i, j int) bool {
			return notes[i].Cos > notes[j].Cos
		})

		if err = result.Err(); err != nil {
			log.Print("err while reading notes")
			c.Status(400)
			return
		}

		c.JSON(200, gin.H{
			"data": notes,
		})
	})

	r.POST("/api/add", func(c *gin.Context) {
		id, err := getid(c)
		if err != nil {
			c.Status(400)
			return
		}
		type Addnote struct {
			Text    string `json:"text"`
			Content string `json:"content"`
		}
		var addnote Addnote
		err = c.ShouldBindJSON(&addnote)
		if err != nil {
			c.Status(400)
			return
		}
		err, embedding := rag(addnote.Content)
		em1 := Embedding{
			Em: embedding,
		}
		byteembedd, err := json.Marshal(em1)
		order := "insert into notes(userid,content,title,embedding)values(?,?,?,?);"
		_, err = db.Exec(order, id, addnote.Content, addnote.Text, string(byteembedd))
		if err != nil {
			c.Status(400)
			return
		}
		c.Status(200)
	})

	r.PUT("/api/edit", func(c *gin.Context) {
		var note Notes

		err := c.ShouldBindJSON(&note)
		if err != nil {
			c.Status(400)
			return
		}

		userID, err := getid(c)
		if err != nil {
			c.Status(401)
			return
		}

		err, embedding := rag(note.Content)
		if err != nil {
			c.Status(400)
			return
		}

		embeddingJSON, err := json.Marshal(Embedding{
			Em: embedding,
		})
		if err != nil {
			c.Status(500)
			return
		}

		order := `
        update notes
        set title=?, content=?, embedding=?
        where id=? and userid=?
    `

		_, err = db.Exec(
			order,
			note.Title,
			note.Content,
			string(embeddingJSON),
			note.ID,
			userID,
		)
		if err != nil {
			c.Status(400)
			return
		}

		c.Status(200)
	})

	r.PUT("/api/delete/:id1", func(c *gin.Context) {

		tid := c.Param("id1")
		id, err := strconv.Atoi(tid)
		order := `delete from notes where id=?`
		_, err = db.Exec(order, id)
		if err != nil {
			c.Status(400)
			return
		}
		c.Status(200)

	})
	return r

}
